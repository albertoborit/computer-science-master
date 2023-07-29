package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to insert a new node into the binary search tree
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

// Function to perform in-order traversal of the binary search tree
func inOrderTraversal(root *TreeNode, sortedArr *[]int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, sortedArr)
	*sortedArr = append(*sortedArr, root.Val)
	inOrderTraversal(root.Right, sortedArr)
}

// Binary Tree Sort function
func binaryTreeSort(arr *[]int) {
	// Record the start time
	start := time.Now()

	var root *TreeNode
	for _, num := range *arr {
		root = insert(root, num)
	}

	*arr = []int{}
	inOrderTraversal(root, arr)

	// Record the end time
	end := time.Now()

	// Calculate the duration and print the execution time
	duration := end.Sub(start)
	fmt.Println("Binary Tree Sort Execution Time:", duration)
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

	// Apply Binary Tree Sort to the numbers
	binaryTreeSort(&numbers)

	fmt.Println("Sorted array:")
	printSlice(numbers)
}
