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

        mergeSort(arr, low, mid)
        mergeSort(arr, mid+1, high)

        merge(arr, low, mid, high)
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
        fmt.Println("Time taken by Merge Sort:", duration)
    }
}
