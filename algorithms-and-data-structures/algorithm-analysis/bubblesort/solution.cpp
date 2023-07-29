#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <chrono>

using namespace std;
using namespace std::chrono;
// Function to perform Bubble Sort
void bubbleSort(vector<int>& arr) {
    int n = arr.size();

    // Record the start time
    auto start = high_resolution_clock::now();

    for (int i = 0; i < n - 1; ++i) {
        bool swapped = false;

        for (int j = 0; j < n - i - 1; ++j) {
            if (arr[j] > arr[j + 1]) {
                swap(arr[j], arr[j + 1]);
                swapped = true;
            }
        }

        // If no two elements were swapped in the inner loop, the array is already sorted
        if (!swapped) {
            break;
        }
    }

    // Record the end time
    auto end = high_resolution_clock::now();

    // Calculate the duration and print the execution time
    auto duration = duration_cast<microseconds>(end - start);
    cout << "Bubble Sort Execution Time: " << duration.count() << " microseconds" << endl;
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
    bubbleSort(numbers);

    cout << "Sorted array: ";
    printArray(numbers);

    return 0;
}
