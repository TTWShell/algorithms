# -*- encoding: utf-8 -*-


def insert_sort(st_bf):
    for i in xrange(1, len(st_bf)):
        temp = st_bf[i]
        j = i - 1
        while j >= 0 and st_bf[j] > temp:
            st_bf[j + 1] = st_bf[j]
            j -= 1
        st_bf[j + 1] = temp
        print st_bf
    return st_bf


st_bf = [6, 5, 3, 1, 8, 7, 2, 4, 2]
insert_sort(st_bf)
