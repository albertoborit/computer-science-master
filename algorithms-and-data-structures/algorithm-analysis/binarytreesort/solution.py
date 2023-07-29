import time

# Definition for a binary tree node
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

# Function to insert a new node into the binary search tree
def insert(root, val):
    if root is None:
        return TreeNode(val)

    if val < root.val:
        root.left = insert(root.left, val)
    else:
        root.right = insert(root.right, val)

    return root

# Function to perform in-order traversal of the binary search tree
def in_order_traversal(root, sorted_arr):
    if root is None:
        return

    in_order_traversal(root.left, sorted_arr)
    sorted_arr.append(root.val)
    in_order_traversal(root.right, sorted_arr)

# Binary Tree Sort function
def binary_tree_sort(arr):
    # Record the start time
    start = time.time()

    root = None
    for num in arr:
        root = insert(root, num)

    arr.clear()
    in_order_traversal(root, arr)

    # Record the end time
    end = time.time()

    # Calculate the duration and print the execution time
    duration = (end - start) * 1000000
    print("Binary Tree Sort Execution Time:", duration, "microseconds")

# Function to read numbers from a file and store them in a list
def read_numbers_from_file(filename):
    with open(filename, 'r') as inputFile:
        numbers = [int(line.strip()) for line in inputFile]
    return numbers

# Function to print the elements of a list
def print_list(arr):
    print(*arr)

def main():
    filenames = ["file1.txt", "file2.txt", "file3.txt"]
    numbers = []

    # Read numbers from each file in the loop
    for filename in filenames:
        nums = read_numbers_from_file(filename)
        numbers.extend(nums)

    print("Numbers read from files:")
    print_list(numbers)

    # Apply Binary Tree Sort to the numbers
    binary_tree_sort(numbers)

    print("Sorted array:")
    print_list(numbers)

if __name__ == "__main__":
    main()
