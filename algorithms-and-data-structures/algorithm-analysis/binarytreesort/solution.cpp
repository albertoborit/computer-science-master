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

    TreeNode* root = nullptr;
    for (int num : arr) {
        root = insert(root, num);
    }

    arr.clear();
    inOrderTraversal(root, arr);
}

void readNumbersFromFile(const string& filename) {
    ifstream inputFile(filename);
    vector<int> numbers;

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

int main() {
    vector<string> filenames = {"file101.txt", "file1001.txt", "file2001.txt","file3001.txt", "file4001.txt", "file5001.txt","file6001.txt", "file7001.txt", "file8001.txt", "file9001.txt", "file10001.txt", "file20001.txt", "file30001.txt", "file40001.txt", "file50001.txt"};

    for (const string& filename : filenames) {
        auto start = high_resolution_clock::now();
        readNumbersFromFile("../"+filename);
        auto end = high_resolution_clock::now();
        auto duration = duration_cast<microseconds>(end - start);
        cout << "Time taken by Binary Tree Sort: " << duration.count() << " microseconds" << endl;
    }

    return 0;
}

