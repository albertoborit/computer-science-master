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
        pivotIndex := partition(arr, low, high)
        quickSort(arr, low, pivotIndex-1)
        quickSort(arr, pivotIndex+1, high)
    }
}

// Function to read numbers from a file and store them in a slice
func readNumbersFromFile(filename string) ([]int) {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil
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

    return numbers
}

func main() {
    filenames := []string {"file101.txt", "file1001.txt", "file2001.txt","file3001.txt", "file4001.txt", "file5001.txt","file6001.txt", "file7001.txt", "file8001.txt", "file9001.txt", "file10001.txt", "file20001.txt", "file30001.txt", "file40001.txt", "file50001.txt"}

    // Read numbers from each file in the loop
    for _, filename := range filenames {
        start := time.Now()
        readNumbersFromFile("../"+filename)
        end := time.Now()
        duration := end.Sub(start)
        fmt.Println("Time taken by Quick Sort:", duration)
    }
}
