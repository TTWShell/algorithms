#include <iostream>
using namespace std;


void print(int st_bf[], size_t size) {
    for (size_t i = 0; i < size; i++)
    {
        cout << st_bf[i] << " ";
    }
    cout << endl;
}


void shellsort(int *st_bf, size_t size) {
    for (size_t k = size/2; k >= 1; k = k/2) {
        cout << "k = " << k << endl;
        for (int i = k; i < size; ++i) {
            int temp = st_bf[i];
            int j = i - k;
            while (j >= 0 and st_bf[j] > temp) {
                st_bf[j+k] = st_bf[j];
                j -= k;
            }
            st_bf[j+k] = temp;
        }
        print(st_bf, size);
    }
}


int main(){
    int st_bf[] = {6, 5, 3, 1, 8, 7, 2, 4, 2};
    size_t size = sizeof(st_bf) / sizeof(st_bf[0]);
    shellsort(st_bf, size);
}
