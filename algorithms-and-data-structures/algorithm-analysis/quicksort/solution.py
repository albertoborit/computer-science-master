import time

def swap(arr, a, b):
    arr[a], arr[b] = arr[b], arr[a]

def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1

    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            swap(arr, i, j)

    swap(arr, i + 1, high)
    return i + 1

def quick_sort(arr, low, high):
    if low < high:
        pivot_index = partition(arr, low, high)
        quick_sort(arr, low, pivot_index - 1)
        quick_sort(arr, pivot_index + 1, high)

def read_numbers_from_file(filename):
    numbers = []
    with open(filename, 'r') as file:
        for line in file:
            numbers.append(int(line.strip()))
    return numbers
    
def main():
    filenames = ["file101.txt", "file1001.txt", "file2001.txt","file3001.txt", "file4001.txt", "file5001.txt","file6001.txt", "file7001.txt", "file8001.txt", "file9001.txt", "file10001.txt", "file20001.txt", "file30001.txt", "file40001.txt", "file50001.txt"]
    for filename in filenames:
        start = time.time()
        quick_sort(read_numbers_from_file("../"+filename), 0, len(read_numbers_from_file("../"+filename)) - 1)
        end = time.time()
        duration = (end - start) * 1e6
        print("Time taken by Quick Sort: {:.2f} microseconds".format(duration))

if __name__ == "__main__":
    main()
