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
        mergeSort(arr, low, mid);
        mergeSort(arr, mid + 1, high);

        merge(arr, low, mid, high);
    }
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
        cout << "Time taken by Merge Sort: " << duration.count() << " microseconds" << endl;
    }

    return 0;
}

