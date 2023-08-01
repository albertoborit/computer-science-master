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

        int pivotIndex = partition(arr, low, high);

        quickSort(arr, low, pivotIndex - 1);
        quickSort(arr, pivotIndex + 1, high);
    }
}

// Function to read numbers from a file and store them in a vector
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
        cout << "Time taken by Quick Sort: " << duration.count() << " microseconds" << endl;
    }

    return 0;
}
