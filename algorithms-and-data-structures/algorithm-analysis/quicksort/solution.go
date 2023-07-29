package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "time"
)

// Function to swap two elements
func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

// Partition function to place the pivot element at the correct position
func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1

    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            swap(arr, i, j)
        }
    }

    swap(arr, i+1, high)
    return i + 1
}

// Quick Sort function
func quickSort(arr []int, low, high int) {
    if low < high {
        // Print start time
        start := time.Now()

        pivotIndex := partition(arr, low, high)

        quickSort(arr, low, pivotIndex-1)
        quickSort(arr, pivotIndex+1, high)

        // Print end time
        end := time.Now()
        duration := end.Sub(start)
        fmt.Println("Time taken by Quick Sort:", duration)
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

    fmt.Println("Numbers read from files:")
    printSlice(numbers)

    // Apply Quick Sort to the numbers
    quickSort(numbers, 0, len(numbers)-1)

    fmt.Println("Sorted array:")
    printSlice(numbers)
}
