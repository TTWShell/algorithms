# 归并排序

Merge sort(mergesort)，是创建在归并操作上的一种有效的排序算法，效率为O(n log n)。1945年由约翰·冯·诺伊曼首次提出。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行。

归并操作（merge），也叫归并算法，指的是将两个已经排序的序列合并成一个序列的操作。归并排序算法依赖归并操作。

# 算法原理

![](http://img.blog.csdn.net/20140821222217578?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemhhbmgxMjE4/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

## 迭代法

1. 申请空间，使其大小为两个已经排序串行之和，该空间用来存放合并后的串行；
2. 设定两个指针，最初位置分别为两个已经排序串行的起始位置；
3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
4. 重复步骤3直到某一指针到达串行尾；
5. 将另一串行剩下的所有元素直接复制到合并串行尾。

## 递归法

1. 将序列每相邻两个数字进行归并操作，形成floor(n/2)个序列，排序后每个序列包含两个元素;
2. 将上述序列再次归并，形成floor(n/4)个序列，每个序列包含四个元素;
3. 重复步骤2，直到所有元素排序完毕。
