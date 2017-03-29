#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++)
    {
        cout<< st_bf[i] << " ";
    }
    cout<< endl;
}

void select_sort(int *st_bf, size_t size) {
    // 选择排序
    for (size_t i = 0; i < size - 1; i++) {
        int temp = st_bf[i]; // 缓存需要比较的值
        size_t location = i; // 缓存最小值的位置
        for (size_t j = i + 1; j < size; j++) {
            // 找出剩余部分最小值
            if (st_bf[j] < temp) {
                temp = st_bf[j];
                location = j;
            }
        }
        if (i != location) {
            // 交换当前位置数和最小数
            st_bf[location] = st_bf[i];
            st_bf[i] = temp;
        }
        print(st_bf, size);
    }
}

int main() {
    int st_bf[] = {6, 5, 3, 1, 8, 7, 2, 4, 2};
    size_t size = sizeof(st_bf) / sizeof(st_bf[0]);
    select_sort(st_bf, size);
}
