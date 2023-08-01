import random

def generate_random_numbers(size):
    print(size)
    return [random.randint(1, size) for _ in range(size)]

def save_numbers_to_file(filename, numbers):
    with open(filename, 'w') as file:
        for num in numbers:
            file.write(str(num) + '\n')

def generate_and_save_files(max_numbers):
    for i in max_numbers:
        numbers = generate_random_numbers(i)
        print(numbers)
        filename = f'file{i + 1}.txt'
        save_numbers_to_file(filename, numbers)
        print(f'{filename} created.')

if __name__ == "__main__":
    max_numbers = [100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000]  # Max numbers for each file

    generate_and_save_files(max_numbers)
