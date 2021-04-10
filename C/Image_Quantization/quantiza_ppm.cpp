//Exercicio da disciplina de Computacao Grafica com o professor Marcio Pinho
#include <iostream>
#include <fstream>
#include <cstdlib>
#include <unordered_map>
#include <cmath>
#include <algorithm>
#include <vector>
#include <limits>
#include <bits/stdc++.h>

using namespace std;

class RGB{
public:
    int r,g,b;
    bool operator==(const RGB &other) const
    {
        return (r == other.r &&
                g == other.g &&
                b == other.b);
    }
    double dist(const RGB &other) const
    {
        return sqrt(pow(r - other.r, 2) + pow(g - other.g, 2) + pow(b - other.b, 2));
    }
};

class RGB_hash_func{
public:
    size_t operator()(const RGB& rgb) const
    {
        return (hash<int>()(rgb.r)) ^
               (hash<int>()(rgb.g)) ^
               (hash<int>()(rgb.b));
    }

};

std::pair<RGB, int> find_max_value(
        std::unordered_map<RGB, int, RGB_hash_func> const &x)
{
    return *std::max_element(x.begin(), x.end(),
                             [](const std::pair<RGB, int> &p1,
                                const std::pair<RGB, int> &p2)
                             {
                                 return p1.second < p2.second;
                             });
}

bool safe_incr(unordered_map<RGB, int, RGB_hash_func>& map, RGB &key)
{
    if(map.find(key) == map.end()){
        map[key] = 1;
    }
    else{
        map[key] += 1;
    }
    return true;
}

RGB image[1000][1000];

//global variables
unsigned int width, height;
unordered_map<RGB,int, RGB_hash_func> rgb_count;
vector<RGB> frequent_colors; //will include them in order of frequency

void reduce_color_vector()
{
    while(frequent_colors.size() > 50)//repeat process while size is bigger than needed
    {
        for(int i=0; i<frequent_colors.size()-11; i++)
        {
            double minor_dist = numeric_limits<double>::max();
            int closest_index = -1;
            int n_erases = 0;
            for(int j=i+1; j<frequent_colors.size(); j++)
            {
                auto actual = frequent_colors[j];
                auto actual_dist = frequent_colors[i].dist(actual);
                if(actual_dist < 30)
                {
                    frequent_colors.erase(frequent_colors.begin()+j);
                    n_erases++;
                    if(frequent_colors.size()==50){
                        return;}
                }
                else if(actual_dist < minor_dist)
                {
                    minor_dist = actual_dist;
                    closest_index = j;
                }
            }
            frequent_colors.erase(frequent_colors.begin()+closest_index);
            if(frequent_colors.size() == 50)
                return;
        }
    }
}

void reduce_color_map()
{
    double proximity = 30;
    double factor = 1.1;
    while(frequent_colors.size() + rgb_count.size() > 50)//repeat process while size is bigger than needed
    {
        RGB current = find_max_value(rgb_count).first; //take most frequent color
        for(auto pair = rgb_count.begin(); pair!= rgb_count.end(); pair++)
        {
            RGB i = pair->first;
            if(i == current) {continue;}

            if(current.dist(i) < proximity)
            {
                pair++;
                rgb_count.erase(i);
                if(rgb_count.size() + frequent_colors.size() == 50 || pair==rgb_count.end())
                {
                    break;
                }
            }
        }
        frequent_colors.push_back(current);
        rgb_count.erase(current);

        proximity *= factor;//increments minimum distance to remove
    }
    while(rgb_count.size() > 0)
    {
        auto max = find_max_value(rgb_count).first;
        frequent_colors.push_back(max);
        rgb_count.erase(max);
    }
}

void transform_image()
{
    for(int y= height - 1; y >= 0; y--)
    {
        for(int x=0; x < width; x++)
        {//for each image rgb..
            RGB closest = frequent_colors[0];
            double minor_dist = image[y][x].dist(frequent_colors[0]);
            for(int i=1; i<frequent_colors.size(); i++)
            {
                auto actual = frequent_colors[i];
                auto actual_dist = image[y][x].dist(actual);
                if(actual_dist < minor_dist)
                {
                    closest = actual;
                    minor_dist = actual_dist;
                }
            }
            image[y][x] = closest;
        }
    }
}

//fill frequent_colors vector in order of most frequent colors in rgb_count
void fill_vector()
{
    while(!rgb_count.empty())
    {
        //gets most frequent color from remaining ones in the map
        auto most_frequent = find_max_value(rgb_count);
        //push it to the last position of the vector
        frequent_colors.push_back(most_frequent.first);
        // removes from the map
        rgb_count.erase(most_frequent.first);
    }
}

void salva_ppm(string &nome)
{
    int x,y;

    ofstream arquivo;

    arquivo.open(nome, ios::out);


    arquivo << "P3" << endl;
    arquivo << width << " " << height << endl;
    arquivo << 255 << endl;
    for(y= height - 1; y >= 0; y--)
    {
        for(x=0; x < width; x++)
        {
            arquivo << " " << image[y][x].r << " " << image[y][x].g << " " << image[y][x].b;

        }
        arquivo << endl;
    }
}

void carrega_ppm(string &nome)
{
    int x,y;

    string dummy;

    ifstream arquivo;

    arquivo.open(nome, ios::in);
    if (!arquivo)
    {
        cout << "Arquivo inexistente." << endl;
        exit (1);
    }


    arquivo >> dummy;

    if (dummy != "P3")
    {
        cout << "Formato Invalido" << endl;
        exit (1);
    }
    arquivo >> width;
    if ((width < 0) || (width > 900))
    {
        cout << "Largura Invalida" << endl;
        exit (2);
    }
    arquivo >> height;
    if ((height < 0) || (height > 900))
    {
        cout << "Altura Invalida" << endl;
        exit (3);
    }

    arquivo >> dummy;
    if (dummy != "255")
    {
        cout << "Formato Invalido" << endl;
        exit (4);
    }
    for(y= height - 1; y >= 0; y--)
    {
        for(x=0; x < width; x++)
        {
            RGB px;
            arquivo >> px.r >> px.g >> px.b;
            image[y][x] = px;
            safe_incr(rgb_count, px);
        }
    }
}


int main(int argc, const char * argv[]) {
    time_t start, end;
    int exec_type = 0; //faster execution
    string filename, output;
    if(argc == 2)
    {
        filename = argv[1];
        output = "map_" + (string) argv[1];
    }
    else if(argc == 3)
    {
        filename = argv[1];
        output = argv[1];
        string runmode = argv[2];
        if(runmode == "vector")
        {
            output = "vector_" + output;
            exec_type = 1;
        }
    }
    else{
        filename = "Eagle.ppm";
        output = "map_Eagle.ppm";
    }

    time(&start);
    ios_base::sync_with_stdio(false);
    carrega_ppm(filename);

    cout << rgb_count.size() << endl;

    if(exec_type == 0)
    {
        reduce_color_map();
        transform_image();
    }
    else
    {
        fill_vector();
        cout << frequent_colors.size() << endl;
        reduce_color_vector();
        transform_image();
    }

    salva_ppm(output);
    time(&end);
    double time_taken = double(end - start);
    cout << "Time taken by program is : " << fixed
         << time_taken << setprecision(5);
    cout << " sec " << endl;
    return 0;
}
