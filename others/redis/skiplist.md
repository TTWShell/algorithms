# 跳跃表（Skiplist）

跳跃表（skiplist）是一种有序数据结构， 它通过在每个节点中维持多个指向其他节点的指针， 从而达到快速访问节点的目的。

跳跃表支持平均 O(log N) 最坏 O(N) 复杂度的节点查找， 还可以通过顺序性操作来批量处理节点。

在大部分情况下， 跳跃表的效率可以和平衡树相媲美， 并且因为跳跃表的实现比平衡树要来得更为简单， 所以有不少程序都使用跳跃表来代替平衡树。

和链表、字典等数据结构被广泛地应用在 Redis 内部不同， Redis 只在两个地方用到了跳跃表， 一个是实现有序集合键， 另一个是在集群节点中用作内部数据结构， 除此之外， 跳跃表在 Redis 里面没有其他用途。

## [Struct定义](https://github.com/antirez/redis/blob/3.2.0/src/server.h#L634)

```
/* ZSETs use a specialized version of Skiplists */
typedef struct zskiplistNode {  // 跳跃表节点
    robj *obj;  // redis成员对象
    double score;  // 分值
    struct zskiplistNode *backward;  // 后退指针
    struct zskiplistLevel {  // 层
        struct zskiplistNode *forward;  // 前进指针
        unsigned int span;  // 跨度
    } level[];
} zskiplistNode;

typedef struct zskiplist {
    struct zskiplistNode *header, *tail;  // 表头节点和表尾节点
    unsigned long length;  // 表中节点的数量
    int level;  // 表中层数最大的节点的层数
} zskiplist;
```

redis 的跳跃表是一个双向的链表，并且在zskiplist结构体中保存了跳跃表的长度和头尾节点，方便从头查找或从尾部遍历。

### 层

跳跃表节点的 level 数组可以包含多个元素，每个元素都包含一个指向其他节点的指针，程序可以通过这些层来加快访问其他节点的速度，一般来说，层的数量越多，访问其他节点的速度就越快。

每次创建一个新跳跃表节点的时候，程序都根据幂次定律（power law，越大的数出现的概率越小）随机生成一个介于 1 和 32 之间的值作为 level 数组的大小， 这个大小就是层的“高度”。

### 前进指针

每个层都有一个指向表尾方向的前进指针（level[i].forward 属性）， 用于从表头向表尾方向访问节点。

### 跨度
层的跨度（level[i].span 属性）用于记录两个节点之间的距离：

* 两个节点之间的跨度越大， 它们相距得就越远。
* 指向 NULL 的所有前进指针的跨度都为 0 ， 因为它们没有连向任何节点。

初看上去， 很容易以为跨度和遍历操作有关， 但实际上并不是这样 —— 遍历操作只使用前进指针就可以完成了， 跨度实际上是用来计算排位（rank）的： 在查找某个节点的过程中， 将沿途访问过的所有层的跨度累计起来， 得到的结果就是目标节点在跳跃表中的排位。

### 后退指针

节点的后退指针（backward 属性）用于从表尾向表头方向访问节点： 跟可以一次跳过多个节点的前进指针不同， 因为每个节点只有一个后退指针， 所以每次只能后退至前一个节点。

### 分值和成员

节点的分值（score 属性）是一个 double 类型的浮点数， 跳跃表中的所有节点都按分值从小到大来排序。

节点的成员对象（obj 属性）是一个指针， 它指向一个字符串对象， 而字符串对象则保存着一个 SDS 值。

在同一个跳跃表中， 各个节点保存的成员对象必须是唯一的， 但是多个节点保存的分值却可以是相同的： 分值相同的节点将按照成员对象在字典序中的大小来进行排序，成员对象较小的节点会排在前面（靠近表头的方向），而成员对象较大的节点则会排在后面（靠近表尾的方向）。


### zskiplist

header 和 tail 指针分别指向跳跃表的表头和表尾节点， 通过这两个指针， 程序定位表头节点和表尾节点的复杂度为 O(1) 。

通过使用 length 属性来记录节点的数量， 程序可以在 O(1) 复杂度内返回跳跃表的长度。

level 属性则用于在 O(1) 复杂度内获取跳跃表中层高最大的那个节点的层数量， 注意表头节点的层高并不计算在内。

## 主要API

| 函数 | 作用 | 时间复杂度 |
| --- | --- | --- |
| zskiplist *zslCreate(void) | 创建一个新的跳跃表 | O(1) |
| void zslFree(zskiplist *zsl) | 释放给定跳跃表，以及表中包含的所有节点 | O(N) ， N 为跳跃表的长度 |
| zskiplistNode *zslInsert(zskiplist *zsl, double score, robj *obj) | 将包含给定成员和分值的新节点添加到跳跃表中 | 平均 O(log N) ，最坏 O(N) ， N 为跳跃表长度 |
| int zslDelete(zskiplist *zsl, double score, robj *obj) | 删除跳跃表中包含给定成员和分值的节点 | 平均 O(log N) ，最坏 O(N) ， N 为跳跃表长度 |
| unsigned long zslGetRank(zskiplist *zsl, double score, robj *o) |	返回包含给定成员和分值的节点在跳跃表中的排位 | 平均 O(log N) ，最坏 O(N) ， N 为跳跃表长度 |
| zskiplistNode* zslGetElementByRank(zskiplist *zsl, unsigned long rank) |返回跳跃表在给定排位上的节点 | 平均 O(log N) ，最坏 O(N) ， N 为跳跃表长度 |
| int zslIsInRange(zskiplist *zsl, zrangespec *range) | 给定一个分值范围（range），比如0到15，20到28，诸如此类，如果给定的分值范围包含在跳跃表的分值范围之内， 那么返回1 ，否则返回0 | 通过跳跃表的表头节点和表尾节点，这个检测可以用 O(1) 复杂度完成 |
| zskiplistNode *zslFirstInRange(zskiplist *zsl, zrangespec *range) | 给定一个分值范围，返回跳跃表中第一个符合这个范围的节点 | 平均 O(log N) ，最坏 O(N) 。 N 为跳跃表长度 |
| zskiplistNode *zslLastInRange(zskiplist *zsl, zrangespec *range) | 给定一个分值范围，返回跳跃表中最后一个符合这个范围的节点 | 平均 O(log N) ，最坏 O(N) 。 N 为跳跃表长度 |
| unsigned long zslDeleteRangeByScore(zskiplist *zsl, zrangespec *range, dict *dict) | 给定一个分值范围，删除跳跃表中所有在这个范围之内的节点 | O(N) ， N 为被删除节点数量 |
| unsigned long zslDeleteRangeByRank(zskiplist *zsl, unsigned int start, unsigned int end, dict *dict) |	给定一个排位范围，删除跳跃表中所有在这个范围之内的节点 | O(N) ， N 为被删除节点数量 |


## 总结

1. 跳跃表是一种有序的链表，持有多层索引，利用空间换取时间，平均查找效率为O(logN);
2. redis 的跳跃表底层为双向链表，并持有尾指针方便从尾遍历;
3. reids 的跳跃表最大索引层数为32层，用于支持2^32个元素的索引建立;
4. 出于节省内存的目的，redis 的跳跃表每个元素到上一层索引的概率为0.25；
5. 在同一个跳跃表中， 多个节点可以包含相同的分值， 但每个节点的成员对象必须是唯一的；
6. 跳跃表中的节点按照分值大小进行排序，当分值相同时，节点按照成员对象的大小进行排序。
