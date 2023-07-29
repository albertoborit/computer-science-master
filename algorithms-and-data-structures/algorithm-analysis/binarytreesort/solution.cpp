#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <chrono>

using namespace std;
using namespace std::chrono;

// Definition for a binary tree node
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

// Function to insert a new node into the binary search tree
TreeNode* insert(TreeNode* root, int val) {
    if (root == nullptr) {
        return new TreeNode(val);
    }

    if (val < root->val) {
        root->left = insert(root->left, val);
    } else {
        root->right = insert(root->right, val);
    }

    return root;
}

// Function to perform in-order traversal of the binary search tree
void inOrderTraversal(TreeNode* root, vector<int>& sortedArr) {
    if (root == nullptr) {
        return;
    }

    inOrderTraversal(root->left, sortedArr);
    sortedArr.push_back(root->val);
    inOrderTraversal(root->right, sortedArr);
}

// Binary Tree Sort function
void binaryTreeSort(vector<int>& arr) {
    // Record the start time
    auto start = high_resolution_clock::now();

    TreeNode* root = nullptr;
    for (int num : arr) {
        root = insert(root, num);
    }

    arr.clear();
    inOrderTraversal(root, arr);

    // Record the end time
    auto end = high_resolution_clock::now();

    // Calculate the duration and print the execution time
    auto duration = duration_cast<microseconds>(end - start);
    cout << "Binary Tree Sort Execution Time: " << duration.count() << " microseconds" << endl;
}

// Function to read numbers from a file and store them in a vector
void readNumbersFromFile(const string& filename, vector<int>& numbers) {
    ifstream inputFile(filename);

    if (!inputFile.is_open()) {
        cout << "Error opening file: " << filename << endl;
        return;
    }

    int number;
    while (inputFile >> number) {
        numbers.push_back(number);
    }

    inputFile.close();
}

// Function to print the elements of an array
void printArray(const vector<int>& arr) {
    for (int num : arr) {
        cout << num << " ";
    }
    cout << endl;
}

int main() {
    vector<string> filenames = {"file1.txt", "file2.txt", "file3.txt"};
    vector<int> numbers;

    // Read numbers from each file in the loop
    for (const string& filename : filenames) {
        readNumbersFromFile(filename, numbers);
    }

    cout << "Numbers read from files:" << endl;
    printArray(numbers);

    // Apply Binary Tree Sort to the numbers
    binaryTreeSort(numbers);

    cout << "Sorted array: ";
    printArray(numbers);

    return 0;
}
