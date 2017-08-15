# -*- encoding: utf-8 -*-


def shell_sort(st_bf):
    length = len(st_bf)

    step = length / 2
    while step >= 1:
        print 'steps = {}'.format(step)
        for i in xrange(step, length):
            temp = st_bf[i]
            j = i - step
            while j >= 0 and st_bf[j] > temp:
                st_bf[j+step] = st_bf[j]
                j -= step
            st_bf[j+step] = temp
        print st_bf
        step = step / 2


st_bf = [6, 5, 3, 1, 8, 7, 2, 4, 2]
shell_sort(st_bf)
