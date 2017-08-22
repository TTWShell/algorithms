#include <iostream>
using namespace std;

void swap(int* a, int* b) {
    int temp = *b;
    *b = *a;
    *a = temp;
}

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout<< st_bf[i] << " ";
    }
    cout<< endl;
}

void max_heapify(int arr[], size_t start, size_t end) {
    size_t dad = start;
    size_t son = dad * 2 + 1;
    while (son <= end) {
        // 先比较两个子节点的大小，选择最大的
        if (son + 1 <= end and arr[son] < arr[son + 1]) {
            son++;
        }
        // 如果父节点大于子节点代表调整完成，直接跳出函数
        if (arr[dad] > arr[son]) {
            return;
        } else { // 否则交换父子内容再继续子节点和孙子节点比较
            swap(arr[dad], arr[son]);
            dad = son;
            son = dad * 2 + 1;
        }
    }
    print(arr, end);
}

void heap_sort(int arr[], size_t len) {
    int i;
    // 初始化，i从最后一个父节点开始调整
    for (i = len / 2 - 1; i >= 0; i--) {
        max_heapify(arr, i, len - 1);
        print(arr, len);
    }
    // 将第一个元素和已排好元素的前一个做交换，再从新调整，直到排序完成
    for (i = len - 1; i > 0; i--) {
        swap(&arr[0], &arr[i]);
        max_heapify(arr, 0, i - 1);
        print(arr, len);
    }
}

int main() {
    int st_bf[9] = {6, 5, 3, 1, 8, 7, 2, 4, 2};
    size_t size = sizeof(st_bf) / sizeof(st_bf[0]);
    heap_sort(st_bf, size);
}
