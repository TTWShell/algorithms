# 字典（Hash Tables）

## Struct定义

```
typedef struct dictEntry {
    void *key; // 键
    union {
        void *val;
        uint64_t u64;
        int64_t s64;
        double d;
    } v; // 值
    struct dictEntry *next; // 拉链法解决冲突，下一个节点
} dictEntry;

// 当存在相同hash值的节点时，采用拉链法解决冲突。
// 字典节点的val为一个union，可以是指针、uint64_t、int64_t或者是double。

typedef struct dictType {  // 各种字典操作
    unsigned int (*hashFunction)(const void *key);   // 计算hash值的函数
    void *(*keyDup)(void *privdata, const void *key);  // 键复制
    void *(*valDup)(void *privdata, const void *obj);  // 值复制
    int (*keyCompare)(void *privdata, const void *key1, const void *key2);  // 键比较
    void (*keyDestructor)(void *privdata, void *key);  // 键销毁
    void (*valDestructor)(void *privdata, void *obj);  // 值销毁
} dictType;

/* This is our hash table structure. Every dictionary has two of this as we
 * implement incremental rehashing, for the old to the new table. */
typedef struct dictht {  // hash表
    dictEntry **table;  // 节点数组
    unsigned long size;  // hash表大小
    unsigned long sizemask;  // hash表掩码，等于size-1，用于计算hash值
    unsigned long used;  // 已有节点数量
} dictht;

typedef struct dict {  // 字典
    dictType *type;  // 各种字典操作方法
    void *privdata;  // 私有数据，用于传给操作函数
    dictht ht[2];  // 两个hash表，一个用来存储当前使用的，一个用来rehash
    long rehashidx; /* rehashing not in progress if rehashidx == -1 // rehash标志位，用于判断是否在rehash和记录rehash进度 */
    int iterators; /* number of iterators currently running // 迭代器的运行数量 */
} dict;

/* If safe is set to 1 this is a safe iterator, that means, you can call
 * dictAdd, dictFind, and other functions against the dictionary even while
 * iterating. Otherwise it is a non safe iterator, and only dictNext()
 * should be called while iterating. */
typedef struct dictIterator {
    dict *d;
    long index;
    int table, safe;
    dictEntry *entry, *nextEntry;
    /* unsafe iterator fingerprint for misuse detection. */
    long long fingerprint;
} dictIterator;
```

## rehash步骤

扩展和收缩哈希表的工作可以通过执行 rehash （重新散列）操作来完成， Redis 对字典的哈希表执行 rehash 的步骤如下：

1. 为字典的 ht[1] 哈希表分配空间， 这个哈希表的空间大小取决于要执行的操作， 以及 ht[0] 当前包含的键值对数量 （也即是 ht[0].used 属性的值）：

    * 如果执行的是扩展操作， 那么 ht[1] 的大小为第一个大于等于 ht[0].used * 2 的 2^n （2 的 n 次方幂）；
    * 如果执行的是收缩操作， 那么 ht[1] 的大小为第一个大于等于 ht[0].used 的 2^n 。
2. 将保存在 ht[0] 中的所有键值对 rehash 到 ht[1] 上面： rehash 指的是重新计算键的哈希值和索引值， 然后将键值对放置到 ht[1] 哈希表的指定位置上。
3. 当 ht[0] 包含的所有键值对都迁移到了 ht[1] 之后 （ht[0] 变为空表）， 释放 ht[0] ， 将 ht[1] 设置为 ht[0] ， 并在 ht[1] 新创建一个空白哈希表， 为下一次 rehash 做准备。

当以下条件中的任意一个被满足时， 程序会自动开始对哈希表执行扩展操作：

1. 服务器目前没有在执行 BGSAVE 命令或者 BGREWRITEAOF 命令， 并且哈希表的负载因子大于等于 1 ；
2. 服务器目前正在执行 BGSAVE 命令或者 BGREWRITEAOF 命令， 并且哈希表的负载因子大于等于 5 ；

其中哈希表的负载因子可以通过公式：

```
# 负载因子 = 哈希表已保存节点数量 / 哈希表大小
load_factor = ht[0].used / ht[0].size
```

根据 BGSAVE 命令或 BGREWRITEAOF 命令是否正在执行， 服务器执行扩展操作所需的负载因子并不相同， 这是因为在执行 BGSAVE 命令或 BGREWRITEAOF 命令的过程中， Redis 需要创建当前服务器进程的子进程， 而大多数操作系统都采用写时复制（copy-on-write）技术来优化子进程的使用效率， 所以在子进程存在期间， 服务器会提高执行扩展操作所需的负载因子， 从而尽可能地避免在子进程存在期间进行哈希表扩展操作， 这可以避免不必要的内存写入操作， 最大限度地节约内存。

另一方面， 当哈希表的负载因子小于 0.1 时， 程序自动开始对哈希表执行收缩操作。

## 渐进式 rehash

为了避免 rehash 对服务器性能造成影响（一次性rehash所有数据会导致一段时间无法提供服务）， 服务器不是一次性将 ht[0] 里面的所有键值对全部 rehash 到 ht[1] ， 而是分多次、渐进式地将 ht[0] 里面的键值对慢慢地 rehash 到 ht[1] 。

以下是哈希表渐进式 rehash 的详细步骤：

1. 为 ht[1] 分配空间， 让字典同时持有 ht[0] 和 ht[1] 两个哈希表。
2. 在字典中维持一个索引计数器变量 rehashidx ， 并将它的值设置为 0 ， 表示 rehash 工作正式开始。
3. 在 rehash 进行期间， 每次对字典执行添加、删除、查找或者更新操作时， 程序除了执行指定的操作以外， 还会顺带将 ht[0] 哈希表在 rehashidx 索引上的所有键值对 rehash 到 ht[1] ， 当 rehash 工作完成之后， 程序将 rehashidx 属性的值增一。
4. 随着字典操作的不断执行， 最终在某个时间点上， ht[0] 的所有键值对都会被 rehash 至 ht[1] ， 这时程序将 rehashidx 属性的值设为 -1 ， 表示 rehash 操作已完成。

渐进式 rehash 的好处在于它采取分而治之的方式， 将 rehash 键值对所需的计算工作均滩到对字典的每个添加、删除、查找和更新操作上， 从而避免了集中式 rehash 而带来的庞大计算量。

### 渐进式 rehash 执行期间的哈希表操作

因为在进行渐进式 rehash 的过程中， 字典会同时使用 ht[0] 和 ht[1] 两个哈希表， 所以在渐进式 rehash 进行期间， 字典的删除（delete）、查找（find）、更新（update）等操作会在两个哈希表上进行： 比如说， 要在字典里面查找一个键的话， 程序会先在 ht[0] 里面进行查找， 如果没找到的话， 就会继续到 ht[1] 里面进行查找， 诸如此类。

另外， 在渐进式 rehash 执行期间，**新添加到字典的键值对一律会被保存到 ht[1] 里面， 而 ht[0] 则不再进行任何添加操作： 这一措施保证了 ht[0] 包含的键值对数量会只减不增， 并随着 rehash 操作的执行而最终变成空表**。

## 主要设计点：

1. redis 的字典是由两个hash表组成，第一个hash表是正常保存数据，第二个hash表仅用来rehash；
2. hash表使用链地址法（separate chaining）来解决冲突。因为 dictEntry 节点组成的链表没有指向链表表尾的指针，所以为了速度考虑，**程序总是将新节点添加到链表的表头位置（复杂度为 O(1)）， 排在其他已有节点的前面**；
3. 当没有子进程在运行时，hash表的负载大于1时就会进行rehash扩展hash表大小；有子进程时负载的阈值提高到了5；
4. 当hash表的负载小于0.1时，hash表会进行rehash收缩hash表大小；
5. redis 的rehash渐进式的，每次增删改查操作都会执行一步rehash。

## 主要API

| 函数 | 作用 | 时间复杂度 |
| --- | --- | --- |
| dict *dictCreate(dictType *type, void *privDataPtr) |	创建一个新的字典 | O(1) |
| int dictAdd(dict *d, void *key, void *val) | 将给定的键值对添加到字典里面 | O(1) |
| int dictReplace(dict *d, void *key, void *val) | 将给定的键值对添加到字典里面， 如果键已经存在于字典，那么用新值取代原有的值 |	O(1) |
| void *dictFetchValue(dict *d, const void *key) |	返回给定键的值 | O(1) |
| dictEntry *dictGetRandomKey(dict *d) | 从字典中随机返回一个键值对 | O(1) |
| int dictDelete(dict *d, const void *key) | 从字典中删除给定键所对应的键值对 | O(1) |
| void dictRelease(dict *d) | 释放给定字典，以及字典中包含的所有键值对 | O(N) ，N 为字典包含的键值对数量 |
