# -*- encoding: utf-8 -*-


def merge_sort(list_one, list_two):
    # 归并排序
    list_merge = []
    i, j = 0, 0
    while i < len(list_one) and j < len(list_two):
        if list_one[i] <= list_two[j]:
            list_merge.append(list_one[i])
            i += 1
        else:
            list_merge.append(list_two[j])
            j += 1
        print 'i = {}, j = {}, len(merge) = {}'.format(i, j, len(list_merge))
        print list_merge

    if i < len(list_one):
        list_merge.extend(list_one[i:])
    if j < len(list_two):
        list_merge.extend(list_two[j:])
    print list_merge


list_one = [1, 3, 5, 6, 8, 10, 12, 13]
list_two = [2, 4, 5, 7, 9, 11, 14, 15, 16, 17, 18]
merge_sort(list_one, list_two)
