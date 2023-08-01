import time

# Definition for a binary tree node
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

def insert(root, val):
    if root is None:
        return TreeNode(val)

    if val < root.val:
        root.left = insert(root.left, val)
    else:
        root.right = insert(root.right, val)

    return root

def in_order_traversal(root, sorted_arr):
    if root is None:
        return

    in_order_traversal(root.left, sorted_arr)
    sorted_arr.append(root.val)
    in_order_traversal(root.right, sorted_arr)

def binary_tree_sort(arr):

    root = None
    for num in arr:
        root = insert(root, num)

    arr.clear()
    in_order_traversal(root, arr)
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
        binary_tree_sort(read_numbers_from_file("../"+filename))
        end = time.time()
        duration = (end - start) * 1e6
        print("Time taken by Quick Sort: {:.2f} microseconds".format(duration))

if __name__ == "__main__":
    main()
