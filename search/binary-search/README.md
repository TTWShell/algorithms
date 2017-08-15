# 二分搜索算法

在计算机科学中，二分搜索（英语：binary search），也称折半搜索（英语：half-interval search）、对数搜索（英语：logarithmic search），是一种在有序数组中查找某一特定元素的搜索演算法。

搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。

![](https://upload.wikimedia.org/wikipedia/commons/f/f7/Binary_search_into_array.png)

# 步骤

給予一個包含n個帶值元素的陣列A或是記錄A0 ... An−1，使得A0 ≤ ... ≤ An−1，以及目標值T，還有下列用來搜尋T在A中位置的子程式。

1. 令L為0，R為n− 1。
2. 如果L > R，則搜尋以失敗告終。
3. 令m（中間值元素）為「(L + R) / 2」加上下高斯符號。
4. 如果Am < T，令L為m + 1並回到步驟二。
5. 如果Am > T，令R為m - 1並回到步驟二。
6. 當Am = T，搜尋結束；回傳值m。

[https://zh.wikipedia.org/wiki/二分搜索算法](https://zh.wikipedia.org/wiki/二分搜索算法)
