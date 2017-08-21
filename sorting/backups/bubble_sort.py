# -*- encoding: utf-8 -*-


def bubble_sort(st_bf):
    for step in xrange(len(st_bf), 0, -1):
        flag = False
        for i in xrange(1, step):
            if st_bf[i-1] > st_bf[i]:
                st_bf[i-1], st_bf[i] = st_bf[i], st_bf[i-1]
                flag = True
        if not flag:
            break
        print st_bf


st_bf = [6, 5, 3, 1, 8, 7, 2, 4, 2]
bubble_sort(st_bf)
