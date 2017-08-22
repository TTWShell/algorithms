#插入排序

![](http://img.blog.csdn.net/20140818135209035?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemhhbmgxMjE4/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：（来自维基百科）

1. 从第一个元素开始，该元素可以认为已经被排序；
2. 取出下一个元素，在已经排序的元素序列中从后向前扫描；
3. 如果该元素（已排序）大于新元素，将该元素移到下一位置；
4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
5. 将新元素插入到该位置后；
6. 重复步骤2~5。
如果比较操作的代价比交换操作大的话，可以采用二分查找法来减少比较操作的数目。该算法可以认为是插入排序的一个变种，称为二分查找排序。

更直观的图片如下：【图片来源于互联网】

![](http://img.blog.csdn.net/20140818135952480?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemhhbmgxMjE4/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

简言之：逐步构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
