package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "time"
)

// Function to merge two sorted arrays into one sorted array
func merge(arr []int, low, mid, high int) {
    leftSize := mid - low + 1
    rightSize := high - mid

    left := make([]int, leftSize)
    right := make([]int, rightSize)

    for i := 0; i < leftSize; i++ {
        left[i] = arr[low+i]
    }

    for i := 0; i < rightSize; i++ {
        right[i] = arr[mid+1+i]
    }

    i, j, k := 0, 0, low
    for i < leftSize && j < rightSize {
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < leftSize {
        arr[k] = left[i]
        i++
        k++
    }

    for j < rightSize {
        arr[k] = right[j]
        j++
        k++
    }
}

// Merge Sort function
func mergeSort(arr []int, low, high int) {
    if low < high {
        mid := low + (high-low)/2
        // Record the start time
        start := time.Now()

        mergeSort(arr, low, mid)
        mergeSort(arr, mid+1, high)

        merge(arr, low, mid, high)
        // Record the end time
        end := time.Now()
        // Calculate the duration and print the execution time
        duration := end.Sub(start)
        fmt.Println("Merge Sort Execution Time:", duration)
    }
}

// Function to read numbers from a file and store them in a slice
func readNumbersFromFile(filename string) ([]int, error) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(content), "\n")
    var numbers []int

    for _, line := range lines {
        number, err := strconv.Atoi(line)
        if err != nil {
            continue // Skip if not a valid number
        }
        numbers = append(numbers, number)
    }

    return numbers, nil
}

// Function to print the elements of a slice
func printSlice(arr []int) {
    for _, num := range arr {
        fmt.Print(num, " ")
    }
    fmt.Println()
}

func main() {
    filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
    var numbers []int

    // Read numbers from each file in the loop
    for _, filename := range filenames {
        nums, err := readNumbersFromFile(filename)
        if err != nil {
            fmt.Println("Error reading file:", err)
            continue
        }
        numbers = append(numbers, nums...)
    }

