#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout << st_bf[i] << " ";
    }
    cout << endl;
}

void merge_sort(int *list_one, size_t len_list_one, int *list_two, size_t len_list_two) {
    // 申请空间
    int list_merge[len_list_one + len_list_two];
    size_t i = 0, j = 0, k = 0; // k保存合并后的长度
    for (; (i < len_list_one) and (j < len_list_two); ++k) {
        if (list_one[i] <= list_two[j]) {
            list_merge[k] = list_one[i++];
        } else {
            list_merge[k] = list_two[j++];
        }
        cout << "i = " << i << ", j = " << j << ", k = " << k << endl;
        print(list_merge, k+1);
    }

    // 以下复制尾部大元素序列
    while (i < len_list_one) list_merge[k++] = list_one[i++];
    while (j < len_list_two) list_merge[k++] = list_two[j++];

    cout << "i = " << i << ", j = " << j << ", k = " << k << endl;
    print(list_merge, k);

}

int main(){
    int list_one[] = {1, 3, 5, 6, 8, 10, 12, 13};
    int list_two[] = {2, 4, 5, 7, 9, 11, 14, 15, 16, 17, 18};
    size_t len_list_one = sizeof(list_one) / sizeof(list_one[0]);
    size_t len_list_two = sizeof(list_two) / sizeof(list_two[0]);
    merge_sort(list_one, len_list_one, list_two, len_list_two);
}
