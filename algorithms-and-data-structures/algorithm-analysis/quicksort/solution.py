import time

# Function to swap two elements
def swap(arr, a, b):
    arr[a], arr[b] = arr[b], arr[a]

# Partition function to place the pivot element at the correct position
def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1

    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            swap(arr, i, j)

    swap(arr, i + 1, high)
    return i + 1

# Quick Sort function
def quick_sort(arr, low, high):
    if low < high:
        # Print start time
        start = time.time()

        pivot_index = partition(arr, low, high)

        quick_sort(arr, low, pivot_index - 1)
        quick_sort(arr, pivot_index + 1, high)

        # Print end time
        end = time.time()
        duration = (end - start) * 1e6
        print("Time taken by Quick Sort: {:.2f} microseconds".format(duration))

# Function to read numbers from a file and store them in a list
def read_numbers_from_file(filename):
    numbers = []
    with open(filename, 'r') as file:
        for line in file:
            numbers.append(int(line.strip()))
    return numbers

# Function to print the elements of a list
def print_list(arr):
    print(*arr)

def main():
    filenames = ["file1.txt", "file2.txt", "file3.txt"]
    numbers = []

    # Read numbers from each file in the loop
    for filename in filenames:
        numbers.extend(read_numbers_from_file(filename))

    print("Numbers read from files:")
    print_list(numbers)

    # Apply Quick Sort to the numbers
    quick_sort(numbers, 0, len(numbers) - 1)

    print("Sorted array:")
    print_list(numbers)

if __name__ == "__main__":
    main()
