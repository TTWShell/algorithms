# -*-coding:utf-8-*-


def heap_sort(lst):
    def sift_down(start, end):
        """最大堆调整"""
        root = start
        while True:
            child = 2 * root + 1
            if child > end:
                break
            if child + 1 <= end and lst[child] < lst[child + 1]:
                child += 1
            if lst[root] < lst[child]:
                lst[root], lst[child] = lst[child], lst[root]
                root = child
            else:
                break
        print lst[start:end]

    # 创建最大堆
    for start in xrange((len(lst) - 2) // 2, -1, -1):
        sift_down(start, len(lst) - 1)
    print lst

    # 堆排序
    for end in xrange(len(lst) - 1, 0, -1):
        lst[0], lst[end] = lst[end], lst[0]
        sift_down(0, end - 1)
        print lst
    return lst


l = [9, 2, 1, 7, 6, 8, 5, 3, 4]
heap_sort(l)
