import random
import time

# Parámetros del algoritmo genético
population_size = 50
mutation_rate = 0.2
generations = 100

def main():
    # Algoritmo genético principal
    

    filenames = ["matrix_10x10.txt", "matrix_20x20.txt","matrix_40x40.txt","matrix_80x80.txt", "matrix_160x160.txt", "matrix_320x320.txt", "matrix_640x640.txt", "matrix_1280x1280.txt", "matrix_2560x2560.txt", "matrix_5120x5120.txt"]
    for filename in filenames:
        start = time.time()
        matrix = read_matrix_from_file(filename)
        population = [create_individual(matrix) for _ in range(population_size)]
        for _ in range(generations):
            new_population = []
            for _ in range(population_size // 2):
                parent1 = select_parents(population, matrix)
                parent2 = select_parents(population, matrix)
                child1 = crossover(parent1, parent2)
                child2 = crossover(parent2, parent1)
                mutate(child1)
                mutate(child2)
                new_population.extend([child1, child2])
            population = new_population
        best_individual = min(population, key=lambda ind: fitness(ind, matrix))
        best_distance = fitness(best_individual, matrix)
        print("Mejor ruta:", best_individual)
        print("Distancia:", best_distance)
        end = time.time()
        duration = (end - start) * 1e6
        print("Time taken by Quick Sort: {:.2f} microseconds".format(duration))

# Crear una población inicial aleatoria
def create_individual(matrix):
    return random.sample(range(len(matrix)), len(matrix))

def read_matrix_from_file(filename):
    matrix = []
    with open(filename, 'r') as file:
        for line in file:
            row_str = line.strip()[1:-1]  # Remove leading '[' and trailing ']\n'
            row = [int(num) for num in row_str.split(',')]
            matrix.append(row)
    return matrix

# Función de aptitud (menor distancia es mejor)
def fitness(individual, matrix):
    total_distance = 0
    for i in range(len(individual) - 1):
        total_distance += matrix[individual[i]][individual[i + 1]]
    total_distance += matrix[individual[-1]][individual[0]]  # Volver al inicio
    return total_distance

# Selección de padres (torneo)
def select_parents(population, matrix):
    parent1 = random.choice(population)
    parent2 = random.choice(population)
    return parent1 if fitness(parent1, matrix) < fitness(parent2, matrix) else parent2

# Cruza de dos padres para crear un hijo
def crossover(parent1, parent2):
    crossover_point = random.randint(1, len(parent1) - 1)
    child = parent1[:crossover_point]
    for gene in parent2:
        if gene not in child:
            child.append(gene)
    return child

# Mutación de un individuo (swap)
def mutate(individual):
    if random.random() < mutation_rate:
        idx1, idx2 = random.sample(range(len(individual)), 2)
        individual[idx1], individual[idx2] = individual[idx2], individual[idx1]


if __name__ == "__main__":
    main()
