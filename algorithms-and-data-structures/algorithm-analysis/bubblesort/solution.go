package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Function to perform Bubble Sort
func bubbleSort(arr []int) {
	n := len(arr)

	// Record the start time
	start := time.Now()

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// If no two elements were swapped in the inner loop, the array is already sorted
		if !swapped {
			break
		}
	}

	// Record the end time
	end := time.Now()

	// Calculate the duration and print the execution time
	duration := end.Sub(start)
	fmt.Println("Bubble Sort Execution Time:", duration)
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
	fmt.Println(arr)
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

	// Apply Bubble Sort to the numbers
	bubbleSort(numbers)

	fmt.Println("Sorted array:")
	printSlice(numbers)
}
