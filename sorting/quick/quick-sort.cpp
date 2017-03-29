#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout << st_bf[i] << " ";
    }
    cout << endl;
}

void quick_sort(int *st_bf, size_t size) {
    size_t j = 0; // 标记哨兵位置
    if (size > 1) {
        int tag = st_bf[0], temp; // 设置哨兵

        for (size_t i = 1; i < size; ++i) {
            if (st_bf[i] < tag) {
                temp = st_bf[i];
                st_bf[i] = st_bf[j+1]; // 互换哨兵后一个位置和当前小于哨兵值的数
                st_bf[j+1] = tag; //  新的哨兵位置为原哨兵位置+1
                st_bf[j++] = temp; // 小于哨兵的数移到原哨兵位置
            } // 大于tag不用处理
        }
    } // 空数组，无需处理
    print(st_bf, size);
    // 递归处理左右子区间
    if (j > 0) quick_sort(st_bf, j);
    size_t len_right = size-j-1;
    if (len_right > 0) quick_sort(st_bf+j+1, len_right);
}

int main(){
    int st_bf[] = {4, 5, 9, 2, 6, 8, 2, 1, 7, 3};
    size_t size = sizeof(st_bf) / sizeof(st_bf[0]);
    quick_sort(st_bf, size);
    print(st_bf, size);
}
