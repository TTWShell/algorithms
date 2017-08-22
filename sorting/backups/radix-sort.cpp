#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout<< st_bf[i] << " ";
    }
    cout<< endl;
}

int maxbit(int data[], size_t size) {
    // 辅助函数，求数据的最大位数
    int max_data;
    for (size_t i = 0; i < size; i++) {
        if (max_data < data[i]) {
            max_data = data[i];
        }
    }
    int max_bit = 1;
    while (max_data >= 10) {
        max_data /= 10;
        max_bit++;
    }
    return max_bit;
}

void radix_sort(int data[], size_t size) {
    int max_bit = maxbit(data, size);
    int *tmp = new int[size];
    int *count = new int[10]; // 计数器
    int i, j, k;
    int radix = 1;
    for (i = 1; i <= max_bit; i++) {
        for (j = 0; j < 10; j++) {
            count[j] = 0; // 每次分配前清空计数器
        }
        for (size_t j = 0; j < size; j++) {
            k = (data[j] / radix) % 10; // 统计每个桶中的记录数
            count[k]++;
        }
        for (int j = 1; j < 10; j++) {
            count[j] = count[j - 1] + count[j]; // 将tmp中的位置依次分配给每个桶
        }
        for (j = size - 1; j >= 0; j--) {
            // 将所有桶中记录依次收集到tmp中
            k = (data[j] / radix) % 10;
            tmp[count[k] - 1] = data[j];
            count[k]--;
        }
        for(j = 0; j < int(size); j++) {
            // 将临时数组的内容复制到data中
            data[j] = tmp[j];
        }
        radix = radix * 10;
    }
    delete []tmp;
    delete []count;
}

int main() {
    int st_bf[9] = {126, 35, 3, 10, 856, 72, 21, 4, 2342};
    size_t size = sizeof(st_bf)/sizeof(st_bf[0]);
    radix_sort(st_bf, size);
    print(st_bf, size);
}
