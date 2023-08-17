import random

def generate_random_matrix(size):
    matrix = []
    for _ in range(size):
        row = [random.randint(1, size) for _ in range(size)]
        matrix.append(row)
    return matrix

def save_matrix_to_file(filename, matrix):
    with open(filename, 'w') as file:
        for id, row in enumerate(matrix):
            file.write(str(row) + '\n')
def generate_and_save_files(max_size):
    for size in max_size:
        matrix = generate_random_matrix(size)
        filename = f'matrix_{size}x{size}.txt'
        save_matrix_to_file(filename, matrix)
        print(f'{filename} created.')

if __name__ == "__main__":
    max_size = [10, 20, 40, 80, 160, 320, 640, 1280, 2560, 5120]  # Max size for each matrix

    generate_and_save_files(max_size)