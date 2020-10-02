#include <pthread.h>
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <unistd.h>

using namespace std;

// Max Cores 6 (main + 5)
const int MAX_THREADS = 5;
//matrix size
int SIZE = 10;

//STRUCT A SER MANDADA PARA A THREAD
typedef struct
{
    int init; //inicio da multiplicacao
    int end;  //fim da multiplicacao
} Indexes;

// PROTOTIPOS
void TimeInit(void);
double TimeStart(void);
double TimeStop(double);

// VALOR DO OVERHEAD DA MEDICAO DE TEMPO
static double TimeOverhead = 0.0;

// ESTRUTURA DE DADOS COMPARTILHADA
//int m1[SIZE][SIZE], m2[SIZE][SIZE], mres[SIZE][SIZE];
int **m1, **m2, **mres;
int l1, c1, l2, c2, lres, cres;

//NUMERO DE THREADS A SER UTILIZADO (sera modificado no main) default 1
int n_threads = 1;
int len_thread = SIZE; //tamanho que cada thread executara

// FUNCAO QUE CALCULA O OVERHEAD DA MEDICAO DE TEMPO
void TimeInit()
{
    double t;

    TimeOverhead = 0.0;
    t = TimeStart();
    TimeOverhead = TimeStop(t);
}

// FUNCAO QUE CAPTURA O TEMPO INICIAL DO TRECHO A SER MEDIDO
double TimeStart()
{
    struct timeval tv;
    struct timezone tz;

    if (gettimeofday(&tv, &tz) != 0)
        exit(1);
    return tv.tv_sec + tv.tv_usec / 1000000.0;
}

// FUNCAO QUE CALCULA O TEMPO GASTO NO FINAL DO TRECHO A SER MEDIDO
double TimeStop(double TimeInitial)
{
    struct timeval tv;
    struct timezone tz;
    double Time;

    if (gettimeofday(&tv, &tz) != 0)
        exit(1);
    Time = tv.tv_sec + tv.tv_usec / 1000000.0;
    return Time - TimeInitial - TimeOverhead;
}

void fill_matrix()
{
    int k = 1;
    for (int i = 0; i < SIZE; i++)
    {
        for (int j = 0; j < SIZE; j++)
        {
            if (k % 2 == 0)
                m1[i][j] = -k;
            else
                m1[i][j] = k;
        }
        k++;
    }
    k = 1;
    for (int j = 0; j < SIZE; j++)
    {
        for (int i = 0; i < SIZE; i++)
        {
            if (k % 2 == 0)
                m2[i][j] = -k;
            else
                m2[i][j] = k;
        }
        k++;
    }
}

void allocate_matrix(){
    m1 = new int*[SIZE];
    m2 = new int*[SIZE];
    mres = new int*[SIZE];
    for(int i=0; i<SIZE; i++){
        m1[i] = new int[SIZE];
        m2[i] = new int[SIZE];
        mres[i] = new int[SIZE];
    }
}

int test_matrix()
{
    // VERIFICA SE O RESULTADO DA MULTIPLICACAO ESTA CORRETO
    for (int i = 0; i < SIZE; i++)
    {
        int k = SIZE * (i + 1);
        for (int j = 0; j < SIZE; j++)
        {
            int k_col = k * (j + 1);
            if (i % 2 == 0)
            {
                if (j % 2 == 0)
                {
                    if (mres[i][j] != k_col)
                        return 1;
                }
                else
                {
                    if (mres[i][j] != -k_col)
                        return 1;
                }
            }
            else
            {
                if (j % 2 == 0)
                {
                    if (mres[i][j] != -k_col)
                        return 1;
                }
                else
                {
                    if (mres[i][j] != k_col)
                        return 1;
                }
            }
        }
    }
    return 0;
}

void define_nthreads()
{ //FUNCIONA APENAS PARA TAMANHOS MULTIPLOS DO NUMERO DE THREADS ESCOLHIDOS
    for (int i = MAX_THREADS; i > 1; i--)
    {
        if (SIZE % i == 0)
        {
            n_threads = i;
            len_thread = SIZE / i;
            break;
        }
    }
    cout << "used threads: " << n_threads << endl;
    cout << "threads size: " << len_thread << endl;
}

void *multiply(void *p)
{
    Indexes *idxs = (Indexes *)p;
    // REALIZA A MULTIPLICACAO
    for (int i = idxs->init; i < idxs->end; i++)
    {
        for (int j = 0; j < cres; j++)
        {
            mres[i][j] = 0;
            for (int k = 0; k < c1; k++)
            {
                mres[i][j] += m1[i][k] * m2[k][j];
            }
        }
    }
    return NULL;
}

//Metodo sem threads para testar tempos de exec
void st_multiply()
{
    // REALIZA A MULTIPLICACAO
    for (int i = 0; i < lres; i++)
    {
        for (int j = 0; j < cres; j++)
        {
            mres[i][j] = 0;
            for (int k = 0; k < c1; k++)
            {
                mres[i][j] += m1[i][k] * m2[k][j];
            }
        }
    }
}

int main(int argc, char** argv)
{
    if(argc <= 1){
        cout << "Especifique a dimensao das matrizes" << endl;
        cout << "exemplo: ./multmatx 10" <<endl;
        exit(1);
    }
    SIZE = atoi(argv[1]);
    len_thread = 10;

    int i, j, k;
    double inicio, total;

    // INICIALIZA OS ARRAYS A SEREM MULTIPLICADOS
    l1 = c1 = SIZE;
    l2 = c2 = SIZE;
    if (c1 != l2)
    {
        fprintf(stderr, "Impossivel multiplicar matrizes: parametros invalidos.\n");
        return 1;
    }

    lres = l1;
    cres = c2;
    define_nthreads();

    pthread_t thrd_m[n_threads];
    Indexes idxs[n_threads];
    allocate_matrix();
    fill_matrix();

    for (i = 0; i < n_threads; i++)
    {
        idxs[i].init = i * len_thread;
        idxs[i].end = (i + 1) * len_thread;
        pthread_create(&thrd_m[i], NULL, multiply, &idxs[i]);
    }
    // PREPARA PARA MEDIR TEMPO
    TimeInit();
    inicio = TimeStart();

    //st_multiply();
    for (int i = 0; i < n_threads; i++)
    {
        pthread_join(thrd_m[i], NULL);
        printf("thread end %d \n", i);
    }

    // OBTEM O TEMPO
    total = TimeStop(inicio);

    int res = test_matrix();
    if (res != 0)
    {
        cout << "multiply failed" << endl;
    }

    // MOSTRA O TEMPO DE EXECUCAO
    printf("time %lf \n", total);
    return 0;
}