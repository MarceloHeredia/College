//@author Marcelo Heredia

#include <pthread.h>
#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <unistd.h>
#include <fstream>

using namespace std;

enum State
{
    PENSANDO,
    COMENDO,
    TENTANDO_COMER
};

#define TIME_THINKING 5
#define TIME_EATING 2
#define MAX_WAIT 3
#define FILENAME "filosofos.txt"

int EXEC_TIME = 10;

unsigned int qtd_filosofos;

struct philosopher
{
    string name;
    State actual_state;
    int eating_count;
    int thinking_count;
    int waiting_count;
    pthread_mutex_t fork_left;
    pthread_mutex_t fork_right;
};

void *diner(void *p)
{
    philosopher *ph = (philosopher *)p;
    while (true)
    {
        switch (ph->actual_state)
        {
        case TENTANDO_COMER: //if his last state was thiking or trying to eat, 
        case PENSANDO:              //he'll try to get the forks and eat
            if (pthread_mutex_trylock(&ph->fork_left))
            { //got left fork
                if (pthread_mutex_trylock(&ph->fork_right))
                { //got right fork
                    ph->actual_state = COMENDO; //changes state
                    ph->eating_count++;
                    cout << ph->name << " estado atual: COMENDO" <<endl;

                    sleep(TIME_EATING); //eats for predefined seconds

                    //releases both forks
                    pthread_mutex_unlock(&ph->fork_left);
                    pthread_mutex_unlock(&ph->fork_right);

                }
                else
                {//didnt get right fork
                    pthread_mutex_unlock(&ph->fork_left); //releases left fork
                    ph->actual_state = TENTANDO_COMER; //changes state
                    int wait_time = rand() % MAX_WAIT; //gets random wait time
                    ph->waiting_count++;               //increase waiting counter
                    cout << ph->name << " estado atual: TENTANDO_COMER" << endl;
                    cout << "       tempo de espera: " << wait_time << "s" << endl;
                    sleep(wait_time); //wait for predefined seconds
                }
            }
            else //didnt get left fork
            {
                ph->actual_state = TENTANDO_COMER; //changes state
                int wait_time = rand() % MAX_WAIT; //gets random wait time
                ph->waiting_count++;               //increase waiting counter
                cout << ph->name << " estado atual: TENTANDO_COMER" << endl;
                cout << "       tempo de espera: " << wait_time << "s" << endl;
                sleep(wait_time); //wait for predefined seconds
            }
            break;

        case COMENDO:                    //if his last state was eating, he'll start to think
            ph->actual_state = PENSANDO; //changes state
            ph->thinking_count++;        //increases thinking counter
            cout << ph->name << " estado atual: PENSANDO" << endl;
            sleep(TIME_THINKING); //thinks for predefined seconds
            break;
        }
    }
}

int main(int argc, char *argv[])
{
    if (argc > 1)
    {
        if (argc <= 1)
        {
            cout << "Especifique o tempo de execucao em segundos" << endl;
            cout << "exemplo: ./philosophers_diner 10" << endl;
            exit(0);
        }

        EXEC_TIME = atoi(argv[1]);
    }

    ifstream input;
    input.open(FILENAME, ios::in);

    if (!input)
    {
        cout << "Erro ao abrir " << FILENAME << endl;
        exit(0);
    }

    input >> qtd_filosofos;

    struct philosopher philosophers[qtd_filosofos];
    pthread_t threads[qtd_filosofos];
    pthread_mutex_t semaphores[qtd_filosofos];

    for (int i = 0; i < qtd_filosofos; i++)
    {
        input >> philosophers[i].name;
        philosophers[i].eating_count = 0;
        philosophers[i].thinking_count = 0;
        philosophers[i].waiting_count = 0;
        philosophers[i].fork_left = semaphores[((i - 1) + qtd_filosofos) % qtd_filosofos];
        philosophers[i].fork_right = semaphores[i];
        philosophers[i].actual_state = COMENDO; //pensarao primeiro

    }

    cout << "Configs: " << endl << "Execution time: " << EXEC_TIME << " s" << endl << "Philosophers: " << endl;
    for(int i=0; i<qtd_filosofos; i++)
    {
        cout << philosophers[i].name << " ";
    }
    cout << endl << endl << endl;


    for (int i = 0; i < qtd_filosofos; i++)
    {
        pthread_create(&threads[i], NULL, diner, (struct philosopher *)&philosophers[i]);
    }

    sleep(EXEC_TIME);

    for (int i = 0; i < qtd_filosofos; i++)
    {
        pthread_cancel(threads[i]);
    }

    cout << endl << endl << endl << "Philosophers Report:" << endl;
    for (int i = 0; i < qtd_filosofos; i++)
    {
        cout << "Philosopher " << philosophers[i].name << endl;
        cout << "times he ate: " << philosophers[i].eating_count << endl;
        cout << "times he thought: " << philosophers[i].thinking_count << endl;
        cout << "times he tried to eat: " << philosophers[i].waiting_count << endl
             << endl;
    }

    input.close();
    return 0;
}