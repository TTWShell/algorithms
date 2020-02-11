# 压缩列表（Ziplist）

压缩列表（ziplist）是列表键和哈希键的底层实现之一。

ziplist是 redis 节省内存的典型例子之一，这个数据结构通过特殊的编码方式将数据存储在连续的内存中。在3.2之前是list的基础数据结构之一，在3.2之后被quicklist替代。但是仍然是zset底层实现之一。

## [压缩列表的构成](https://github.com/antirez/redis/blob/3.2.0/src/ziplist.h)

压缩列表是 Redis 为了节约内存而开发的， 由一系列特殊编码的连续内存块组成的顺序型（sequential）数据结构。

一个压缩列表可以包含任意多个节点（entry）， 每个节点可以保存一个字节数组或者一个整数值。

![http://redisbook.com/_images/graphviz-fe42f343a3f32f477efb5e895da547d476a7c97d.png](http://redisbook.com/_images/graphviz-fe42f343a3f32f477efb5e895da547d476a7c97d.png)

压缩列表各个组成部分的详细说明：

| 属性 | 类型 | 长度 | 用途 |
| --- | --- | --- | --- |
| zlbytes | uint32_t | 4字节 | 记录整个压缩列表占用的内存字节数：在对压缩列表进行内存重分配，或者计算 zlend 的位置时使用 |
| zltail | uint32_t | 4字节 | 记录压缩列表表尾节点距离压缩列表的起始地址有多少字节：通过这个偏移量，程序无须遍历整个压缩列表就可以确定表尾节点的地址 |
| zllen | uint16_t | 2字节 | 记录了压缩列表包含的节点数量：当这个属性的值小于 UINT16_MAX（65535）时，这个属性的值就是压缩列表包含节点的数量；当这个值等于 UINT16_MAX 时， 节点的真实数量需要遍历整个压缩列表才能计算得出 |
| entryX | 列表节点 | 不定 | 压缩列表包含的各个节点，节点的长度由节点保存的内容决定 |
| zlend | uint8_t | 1字节 | 特殊值 0xFF（十进制 255 ），用于标记压缩列表的末端 |

## 压缩列表节点的构成

每个压缩列表节点可以保存一个字节数组或者一个整数值，其中，字节数组可以是以下三种长度的其中一种：

1. 长度小于等于 63 （2^{6}-1）字节的字节数组；
2. 长度小于等于 16383 （2^{14}-1） 字节的字节数组；
3. 长度小于等于 4294967295 （2^{32}-1）字节的字节数组；

而整数值则可以是以下六种长度的其中一种：

1. 4 位长，介于 0 至 12 之间的无符号整数；
2. 1 字节长的有符号整数；
3. 3 字节长的有符号整数；
4. int16_t 类型整数；
5. int32_t 类型整数；
6. int64_t 类型整数。

每个压缩列表节点都由 previous_entry_length 、 encoding 、 content 三个部分组成， 如图 7-4 所示。

### previous_entry_length

节点的 previous_entry_length 属性以字节为单位，记录了压缩列表中前一个节点的长度。

previous_entry_length 属性的长度可以是 1 字节或者 5 字节：

* 如果前一节点的长度小于 254 字节， 那么 previous_entry_length 属性的长度为 1 字节： 前一节点的长度就保存在这一个字节里面。
* 如果前一节点的长度大于等于 254 字节， 那么 previous_entry_length 属性的长度为 5 字节： 其中属性的第一字节会被设置为 0xFE （十进制值 254）， 而之后的四个字节则用于保存前一节点的长度。

**压缩列表的从表尾向表头遍历操作就是使用这一原理实现的： 只要我们拥有了一个指向某个节点起始地址的指针， 那么通过这个指针以及这个节点的 previous_entry_length 属性， 程序就可以一直向前一个节点回溯， 最终到达压缩列表的表头节点。**

### encoding

节点的 encoding 属性记录了节点的 content 属性所保存数据的类型以及长度：

* 一字节、两字节或者五字节长， 值的最高位为 00 、 01 或者 10 的是字节数组编码： 这种编码表示节点的 content 属性保存着字节数组，数组的长度由编码除去最高两位之后的其他位记录；
* 一字节长， 值的最高位以 11 开头的是整数编码：这种编码表示节点的 content 属性保存着整数值，整数值的类型和长度由编码除去最高两位之后的其他位记录；

| 编码 | 编码长度 | content 属性保存的值 |
| --- | --- | --- | --- |
| 00bbbbbb | 1 字节	| 长度小于等于63字节的字节数组 |
|01bbbbbb xxxxxxxx | 2 字节 | 长度小于等于16383字节的字节数组 |
|10______ aaaaaaaa bbbbbbbb cccccccc dddddddd | 5 字节 | 长度小于等于4294967295的字节数组 |
| |
| 11000000 | 1 字节 | int16_t 类型的整数 |
| 11010000 | 1 字节 | int32_t 类型的整数 |
| 11100000 | 1 字节 | int64_t 类型的整数 |
| 11110000 | 1 字节 | 24 位有符号整数 |
| 11111110 | 1 字节 | 8 位有符号整数 |
| 1111xxxx | 1 字节 | 使用这一编码的节点没有相应的 content 属性， 因为编码本身的 xxxx 四个位已经保存了一个介于 0 和 12 之间的值， 所以它无须 content 属性 |

### content

节点的 content 属性负责保存节点的值，节点值可以是一个字节数组或者整数，值的类型和长度由节点的 encoding 属性决定。

## 连锁更新

添加新节点（多个连续的、长度介于 250 字节到 253 字节之间的节点 e1 至 eN，前面插入）、删除节点（big、small、多个连续的、长度介于 250 字节到 253 字节之间的节点 e1 至 eN,删除small）都可能会引发连锁更新。

因为连锁更新在最坏情况下需要对压缩列表执行 N 次空间重分配操作，而每次空间重分配的最坏复杂度为 O(N) ，所以连锁更新的最坏复杂度为 O(N^2) 。

要注意的是，尽管连锁更新的复杂度较高，但它真正造成性能问题的几率是很低的：

* 首先，压缩列表里要恰好有多个连续的、长度介于 250 字节至 253 字节之间的节点，连锁更新才有可能被引发，在实际中，这种情况并不多见；
* 其次，即使出现连锁更新，但只要被更新的节点数量不多，就不会对性能造成任何影响：比如说，对三五个节点进行连锁更新是绝对不会影响性能的；

因为以上原因，ziplistPush 等命令的平均复杂度仅为 O(N)，在实际中，我们可以放心地使用这些函数，而不必担心连锁更新会影响压缩列表的性能。

## 重要API

| 函数 | 作用 | 时间复杂度 |
| --- | --- | --- |
| unsigned char *ziplistNew(void) | 创建一个新的压缩列表 | O(1) |
| unsigned char *ziplistPush(unsigned char *zl, unsigned char *s, unsigned int slen, int where) | 创建一个包含给定值的新节点，并将这个新节点添加到压缩列表的表头或者表尾 | 平均 O(N) ，最坏 O(N^2)  |
| unsigned char *ziplistInsert(unsigned char *zl, unsigned char *p, unsigned char *s, unsigned int slen) | 将包含给定值的新节点插入到给定节点之后 | 平均 O(N) ，最坏 O(N^2)  |
| unsigned char *ziplistIndex(unsigned char *zl, int index) | 返回压缩列表给定索引上的节点 | O(N) |
| unsigned char *ziplistFind(unsigned char *p, unsigned char *vstr, unsigned int vlen, unsigned int skip) | 在压缩列表中查找并返回包含了给定值的节点 | 因为节点的值可能是一个字节数组，所以检查节点值和给定值是否相同的复杂度为 O(N)，而查找整个列表的复杂度则为 O(N^2) |
| unsigned char *ziplistNext(unsigned char *zl, unsigned char *p) | 返回给定节点的下一个节点 | O(1) |
| unsigned char *ziplistPrev(unsigned char *zl, unsigned char *p) | 返回给定节点的前一个节点 | O(1) |
| unsigned int ziplistGet(unsigned char *p, unsigned char **sval, unsigned int *slen, long long *lval) | 获取给定节点所保存的值 |	O(1) |
| unsigned char *ziplistDelete(unsigned char *zl, unsigned char **p) | 从压缩列表中删除给定的节点 | 平均 O(N)，最坏 O(N^2)  |
| unsigned char *ziplistDeleteRange(unsigned char *zl, int index, unsigned int num) |	删除压缩列表在给定索引上的连续多个节点 | 平均 O(N)，最坏 O(N^2)  |
| size_t ziplistBlobLen(unsigned char *zl) | 返回压缩列表目前占用的内存字节数 | O(1) |
| unsigned int ziplistLen(unsigned char *zl) | 返回压缩列表目前包含的节点数量 | 节点数量小于 65535 时 O(1) ，大于 65535 时 O(N)  |

因为 ziplistPush 、 ziplistInsert 、 ziplistDelete 和 ziplistDeleteRange 四个函数都有可能会引发连锁更新， 所以它们的最坏复杂度都是 O(N^2) 。

##  总结

1. 压缩列表是一种为节约内存而开发的顺序型数据结构。
2. 压缩列表被用作列表键和哈希键的底层实现之一。
3. 压缩列表可以包含多个节点，每个节点可以保存一个字节数组或者整数值。
4. 添加新节点到压缩列表，或者从压缩列表中删除节点，可能会引发连锁更新操作，但这种操作出现的几率并不高。