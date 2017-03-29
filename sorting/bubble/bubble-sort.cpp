#include <iostream>
using namespace std;

void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        cout << st_bf[i] << " ";
    }
    cout << endl;
}

void bubble_sort(int *st_bf, size_t size) {
    int temp;
    bool flag;
    for(size_t step = size; step > 0; --step) {
        flag = false;
        for (size_t i = 1; i< step; i++) {
            if (st_bf[i-1] > st_bf[i]) {
                temp = st_bf[i];
                st_bf[i] = st_bf[i-1];
                st_bf[i-1] = temp;
                flag = true;
            }
        }
        if (!flag) {
            break;
        }
        print(st_bf, size);
    }
}

int main(){
    int st_bf[] = {6, 5, 3, 1, 8, 7, 2, 4, 2};
    size_t size = sizeof(st_bf) / sizeof(st_bf[0]);
    bubble_sort(st_bf, size);
}
