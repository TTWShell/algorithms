# 树状数组

树状数组或二叉索引树（英语：Binary Indexed Tree），又以其发明者命名为Fenwick树。

其初衷是解决数据压缩里的累积频率（Cumulative Frequency）的计算问题，现多用于高效计算数列的前缀和， 区间和。它可以以 O(log n) 时间得到任意前缀和, 并同时支持在 O(log n) 时间内支持动态单点值的修改。空间复杂度 O(n)。

## 结构起源

按照Peter M. Fenwick的说法，正如所有的整数都可以表示成2的幂和，我们也可以把一串序列表示成一系列子序列的和。

采用这个想法，我们可将一个前缀和划分成多个子序列的和，而划分的方法与数的2的幂和具有极其相似的方式。

一方面，子序列的个数是其二进制表示中1的个数，另一方面，子序列代表的f[i]的个数也是2的幂。

## 存储结构

![https://raw.githubusercontent.com/OIer-Saber/Photos/master/22.jpg](https://raw.githubusercontent.com/OIer-Saber/Photos/master/22.jpg)

### lowbit(x)函数

预备函数，返回参数转为二进制后，最后一个1的位置所代表的数值。

    In [5]: for i in range(1, 9): print(i, i & (-i))
    1 1
    2 2
    3 1
    4 4
    5 1
    6 2
    7 1
    8 8

对于节点i，它的父节点的编号为i+lowbit(i)；i-lowbit(i)则为节点i管辖区间的下个区间的管辖点。

## 基本操作

### 初始化

![https://wikimedia.org/api/rest_v1/media/math/render/svg/1c0b6c1f17c083ac8e2a7fbe8ceef9bb77c5d666](https://wikimedia.org/api/rest_v1/media/math/render/svg/1c0b6c1f17c083ac8e2a7fbe8ceef9bb77c5d666)

### 修改

假设现在要将 A[i] 的值增加delta，那么，需要将 C[i] 覆盖的区间包含 A[i] 的值的节点都加上delta。

### 求和

假设求节点i之前的所有数之和，则为当前管辖范围C[i]累加节点i下个管辖区间，直到所有累加完毕为止。区间则C[m] - C[n]即可。
