import random
import numpy as np
from numpy.random import default_rng

loads = [27, 7, 6, 5, 4, 6, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 27, 7, 6, 5, 4, 6, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

chromosome = 11

population = []
fit_ind = []
partial_pop = []


def init_population():
    global population, fit_ind, partial_pop
    population = np.random.randint(2, size=(chromosome, len(loads)))
    fit_ind = np.zeros([chromosome])
    partial_pop = np.full([chromosome, len(loads)], -1, dtype=int)


def print_population():
    j = 0
    for i in range(0, chromosome):
        print('C', i, '-> [ ', end='', sep='')
        for j in range(0, len(loads)):
            print(population[i, j], end=' ')
        print('] F:', fit_ind[i])


def print_cromo(chromo):
    print('[ ', end='')
    for i in range(0, len(population[chromo])):
        print(population[chromo][i], end=' ')
    print('] F:', fit_ind[chromo])


def fitness_func():
    for i, ind in enumerate(population):
        fit_ind[i] = fitness_individual(ind)


def fitness_individual(individual):
    i0 = 0
    i1 = 0
    for i, v in enumerate(individual):
        if v == 0:
            i0 += loads[i]
        else:
            i1 += loads[i]
    return abs(i0 - i1)


def elitism():
    best_one = np.argmin(fit_ind)
    partial_pop[0] = population[best_one].copy()


def tournament():
    rng = default_rng()
    indexes = rng.choice(chromosome, size=2, replace=False)
    if fit_ind[indexes[0]] < fit_ind[indexes[1]]:
        return indexes[0]
    else:
        return indexes[1]


def crossover():
    for i in range(1, chromosome, 2):
        id1 = tournament()
        id2 = tournament()
        partial_pop[i] = np.concatenate([population[id1, :int(len(loads) / 2)], population[id2, int(len(loads) / 2):]])
        partial_pop[i + 1] = np.concatenate(
            [population[id2, :int(len(loads) / 2)], population[id1, int(len(loads) / 2):]])


def found_solution():
    if np.amin(fit_ind) == 0:
        return True
    return False


def mutation():
    global partial_pop
    rng = default_rng()
    n = rng.integers(3) + 1
    for i in range(n):
        individual = rng.integers(chromosome - 1) + 1
        position = rng.integers(len(loads))
        print('Chromosome', individual, 'mutated in load n', position)
        partial_pop[individual][position] ^= 1  # xor to flip 1 to 0 and 0 to 1


def selection():
    global partial_pop, population
    partial_pop = np.full([chromosome, len(loads)], -1, dtype=int)  # reset partial population
    elitism()
    crossover()
    mutation()


def main():
    global population, partial_pop
    print("Generation 0:")
    init_population()
    fitness_func()
    print_population()
    if found_solution():
        print('Solution found!!!')
        print_cromo(np.argmin(fit_ind))
        return

    for i in range(1, 1000):
        selection()
        population = partial_pop
        fitness_func()
        print('\n', '#' * 50, '\n', sep='')
        print("Generation", i, ':')
        print_population()
        if found_solution():
            print('Solution found!!!')
            print_cromo(np.argmin(fit_ind))
            break


if __name__ == '__main__':
    main()
