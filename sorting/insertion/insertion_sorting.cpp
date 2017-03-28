#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout<< st_bf[i] << " ";
    }
    cout<< endl;
}

void insert_sort(int st_bf[], size_t size) {
    // 插入排序代码
    for (size_t i = 1; i < size; ++i) {
        int temp = st_bf[i];
        int j = i - 1;
        while (j >= 0 and st_bf[j] > temp) {
            st_bf[j + 1] = st_bf[j];
            j -= 1;
        }
        st_bf[j + 1] = temp;
        print(st_bf, size);
    }
}

int main(){
    int st_bf[9] = {6, 5, 3, 1, 8, 7, 2, 4, 2};
    size_t size = sizeof(st_bf)/sizeof(st_bf[0]);
    insert_sort(st_bf, size);
}
