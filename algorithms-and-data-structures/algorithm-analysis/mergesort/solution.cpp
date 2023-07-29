#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <chrono>

using namespace std;
using namespace std::chrono;

// Function to merge two sorted arrays into one sorted array
void merge(vector<int>& arr, int low, int mid, int high) {
    int leftSize = mid - low + 1;
    int rightSize = high - mid;

    vector<int> left(leftSize);
    vector<int> right(rightSize);

    for (int i = 0; i < leftSize; i++) {
        left[i] = arr[low + i];
    }

    for (int i = 0; i < rightSize; i++) {
        right[i] = arr[mid + 1 + i];
    }

    int i = 0, j = 0, k = low;
    while (i < leftSize && j < rightSize) {
        if (left[i] <= right[j]) {
            arr[k] = left[i];
            i++;
        } else {
            arr[k] = right[j];
            j++;
        }
        k++;
    }

    while (i < leftSize) {
        arr[k] = left[i];
        i++;
        k++;
    }

    while (j < rightSize) {
        arr[k] = right[j];
        j++;
        k++;
    }
}

// Merge Sort function
void mergeSort(vector<int>& arr, int low, int high) {
    if (low < high) {
        int mid = low + (high - low) / 2;
        // Record the start time
        auto start = high_resolution_clock::now();

        mergeSort(arr, low, mid);
        mergeSort(arr, mid + 1, high);

        merge(arr, low, mid, high);
        // Record the end time
        auto end = high_resolution_clock::now();
        // Calculate the duration and print the execution time
        auto duration = duration_cast<microseconds>(end - start);
        cout << "Merge Sort Execution Time: " << duration.count() << " microseconds" << endl;
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

    // Apply Merge Sort to the numbers
    mergeSort(numbers, 0, numbers.size() - 1);

    cout << "Sorted array: ";
    printArray(numbers);

    return 0;
}
