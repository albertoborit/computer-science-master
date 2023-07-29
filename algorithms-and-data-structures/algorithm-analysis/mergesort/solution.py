import time

# Function to merge two sorted arrays into one sorted array
def merge(arr, low, mid, high):
    left = arr[low:mid+1]
    right = arr[mid+1:high+1]

    i = j = 0
    k = low
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            arr[k] = left[i]
            i += 1
        else:
            arr[k] = right[j]
            j += 1
        k += 1

    while i < len(left):
        arr[k] = left[i]
        i += 1
        k += 1

    while j < len(right):
        arr[k] = right[j]
        j += 1
        k += 1

# Merge Sort function
def mergeSort(arr, low, high):
    if low < high:
        mid = (low + high) // 2
        # Record the start time
        start = time.time()

        mergeSort(arr, low, mid)
        mergeSort(arr, mid + 1, high)

        merge(arr, low, mid, high)
        # Record the end time
        end = time.time()
        # Calculate the duration and print the execution time
        duration = (end - start) * 1000000  # Microseconds
        print(f"Merge Sort Execution Time: {duration:.2f} microseconds")

# Function to read numbers from a file and store them in a list
def readNumbersFromFile(filename):
    numbers = []
    with open(filename, "r") as inputFile:
        for line in inputFile:
            number = int(line.strip())
            numbers.append(number)
    return numbers

# Function to print the elements of a list
def printList(arr):
    print(*arr)

if __name__ == "__main__":
    filenames = ["file1.txt", "file2.txt", "file3.txt"]
    numbers = []

    # Read numbers from each file in the loop
    for filename in filenames:
        numbers.extend(readNumbersFromFile(filename))

    print("Numbers read from files:")
    printList(numbers)

    # Apply Merge Sort to the numbers
    mergeSort(numbers, 0, len(numbers) - 1)

    print("Sorted array:")
    printList(numbers)
