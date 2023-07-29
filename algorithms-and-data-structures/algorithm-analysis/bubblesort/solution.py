import time

# Function to perform Bubble Sort
def bubble_sort(arr):
    n = len(arr)

    # Record the start time
    start = time.time()

    for i in range(n - 1):
        swapped = False

        for j in range(n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True

        # If no two elements were swapped in the inner loop, the array is already sorted
        if not swapped:
            break

    # Record the end time
    end = time.time()

    # Calculate the duration and print the execution time
    duration = (end - start) * 1000000
    print("Bubble Sort Execution Time:", duration, "microseconds")

# Function to read numbers from a file and store them in a list
def read_numbers_from_file(filename):
    with open(filename, 'r') as inputFile:
        numbers = [int(line.strip()) for line in inputFile]
    return numbers

# Function to print the elements of a list
def print_list(arr):
    print(*arr)

def main():
    filenames = ["file1.txt", "file2.txt", "file3.txt"]
    numbers = []

    # Read numbers from each file in the loop
    for filename in filenames:
        nums = read_numbers_from_file(filename)
        numbers.extend(nums)

    print("Numbers read from files:")
    print_list(numbers)

    # Apply Bubble Sort to the numbers
    bubble_sort(numbers)

    print("Sorted array:")
    print_list(numbers)

if __name__ == "__main__":
    main()
