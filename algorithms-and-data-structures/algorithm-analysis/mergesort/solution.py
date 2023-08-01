import time

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

def mergeSort(arr, low, high):
    if low < high:
        mid = (low + high) // 2

        mergeSort(arr, low, mid)
        mergeSort(arr, mid + 1, high)

        merge(arr, low, mid, high)

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
        mergeSort(read_numbers_from_file("../"+filename), 0, len(read_numbers_from_file("../"+filename)) - 1)
        end = time.time()
        duration = (end - start) * 1e6
        print("Time taken by Quick Sort: {:.2f} microseconds".format(duration))

if __name__ == "__main__":
    main()
