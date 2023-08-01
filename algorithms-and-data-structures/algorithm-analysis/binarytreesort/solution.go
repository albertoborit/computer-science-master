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
