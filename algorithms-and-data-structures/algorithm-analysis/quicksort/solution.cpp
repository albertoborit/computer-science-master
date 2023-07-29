#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <chrono>

using namespace std;
using namespace std::chrono;

// Function to swap two elements
void swap(int& a, int& b) {
    int temp = a;
    a = b;
    b = temp;
}

// Partition function to place the pivot element at the correct position
int partition(vector<int>& arr, int low, int high) {
    int pivot = arr[high];
    int i = low - 1;

    for (int j = low; j < high; j++) {
        if (arr[j] <= pivot) {
            i++;
            swap(arr[i], arr[j]);
        }
    }

    swap(arr[i + 1], arr[high]);
    return i + 1;
}

// Quick Sort function
void quickSort(vector<int>& arr, int low, int high) {
    if (low < high) {
        // Print start time
        auto start = high_resolution_clock::now();

        int pivotIndex = partition(arr, low, high);

        quickSort(arr, low, pivotIndex - 1);
        quickSort(arr, pivotIndex + 1, high);

        // Print end time
        auto end = high_resolution_clock::now();
        auto duration = duration_cast<microseconds>(end - start);
        cout << "Time taken by Quick Sort: " << duration.count() << " microseconds" << endl;
    }
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

    // Apply Quick Sort to the numbers
    quickSort(numbers, 0, numbers.size() - 1);

    cout << "Sorted array: ";
    printArray(numbers);

    return 0;
}
