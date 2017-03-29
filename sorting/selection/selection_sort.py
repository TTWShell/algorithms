# -*- encoding: utf-8 -*-


def select_sort(st_bf):
    length = len(st_bf)
    for i in xrange(length - 1):
        temp, location = st_bf[i], i

        # 找出剩余部分最小值
        for j in xrange(i+1, length):
            if st_bf[j] < temp:
                temp, location = st_bf[j], j

        if location != i:
            st_bf[location], st_bf[i] = st_bf[i], st_bf[location]
        print st_bf


st_bf = [6, 5, 3, 1, 8, 7, 2, 4, 2]
select_sort(st_bf)
