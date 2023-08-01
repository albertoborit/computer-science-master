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
        fmt.Println("Time taken by Bubble Sort:", duration)
    }
}
