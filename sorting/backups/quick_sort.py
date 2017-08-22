# -*- encoding: utf-8 -*-


def quick_sort(st_bf):
    print st_bf
    if len(st_bf) > 1:
        temp = st_bf[0]
        st_now = quick_sort([i for i in st_bf[1:] if i < temp]) + \
            [temp] + \
            quick_sort([j for j in st_bf[1:] if j >= temp])
        return st_now
    else:
        return st_bf


st_bf = [4, 5, 9, 2, 6, 8, 2, 1, 7, 3]
print quick_sort(st_bf)
