# 双向链表(A generic doubly linked list)

## [Strcut定义](https://github.com/antirez/redis/blob/3.2.0/src/adlist.h)

```
/* Node, List, and Iterator are the only data structures used currently. */

typedef struct listNode {
    struct listNode *prev;
    struct listNode *next;
    void *value;
} listNode;

typedef struct listIter {
    listNode *next;
    int direction;
} listIter;

typedef struct list {
    listNode *head;
    listNode *tail;
    void *(*dup)(void *ptr);  // 节点值复制函数
    void (*free)(void *ptr);  // 节点值释放函数
    int (*match)(void *ptr, void *key);  // 节点值对比函数
    unsigned long len;
} list;
```

## 实现特性

1. 双端：链表节点带有 `prev` 和 `next` 指针， 获取某个节点的前置节点和后置节点的复杂度都是 O(1) 。
2. 无环：表头节点的 `prev` 指针和表尾节点的 `next` 指针都指向 `NULL` ， 对链表的访问以 `NULL` 为终点。
3. 带表头指针和表尾指针：通过 `list` 结构的 `head` 指针和 `tail` 指针， 程序获取链表的表头节点和表尾节点的复杂度为 O(1) 。
4. 带链表长度计数器：程序使用 `list` 结构的 `len` 属性来对 `list` 持有的链表节点进行计数，程序获取链表中节点数量的复杂度为 O(1) 。
5. 多态：链表节点使用 `void*` 指针来保存节点值，并且可以通过 `list` 结构的 `dup` 、 `free` 、 `match` 三个属性为节点值设置类型特定函数，所以链表可以用于保存各种不同类型的值。

##  链表和链表节点 API

| 函数 | 作用 | 时间复杂度 |
| --- | --- | --- |
| #define listSetDupMethod(l,m) ((l)->dup = (m)) | 将给定的函数设置为链表的节点值复制函数 | O(1) |
| #define listGetDupMethod(l) ((l)->dup) | 返回链表当前正在使用的节点值复制函数 | 复制函数可以通过链表的 dup 属性直接获得， O(1) |
| #define listSetFreeMethod(l,m) ((l)->free = (m)) | 将给定的函数设置为链表的节点值释放函数 | O(1) |
| #define listGetFree(l) ((l)->free) | 返回链表当前正在使用的节点值释放函数 | 释放函数可以通过链表的 free 属性直接获得， O(1) |
| #define listSetMatchMethod(l,m) ((l)->match = (m)) | 将给定的函数设置为链表的节点值对比函数 | O(1) |
| #define listGetMatchMethod(l) ((l)->match) | 返回链表当前正在使用的节点值对比函数 | 对比函数可以通过链表的 match 属性直接获得， O(1) |
| #define listLength(l) ((l)->len) | 返回链表的长度（包含了多少个节点） | 链表长度可以通过链表的 len 属性直接获得， O(1)  |
| #define listFirst(l) ((l)->head) | 返回链表的表头节点 | 表头节点可以通过链表的 head 属性直接获得， O(1)  |
| #define listLast(l) ((l)->tail) |	返回链表的表尾节点 | 表尾节点可以通过链表的 tail 属性直接获得， O(1)  |
| #define listPrevNode(n) ((n)->prev) | 返回给定节点的前置节点 | 前置节点可以通过节点的 prev 属性直接获得， O(1)  |
| #define listNextNode(n) ((n)->next) |	返回给定节点的后置节点 | 后置节点可以通过节点的 next 属性直接获得， O(1)  |
| #define listNodeValue(n) ((n)->value) | 返回给定节点目前正在保存的值 | 节点值可以通过节点的 value 属性直接获得， O(1)  |
| list *listCreate(void) | 创建一个不包含任何节点的新链表 | O(1) |
| list *listAddNodeHead(list *list, void *value) | 将一个包含给定值的新节点添加到给定链表的表头 | O(1) |
| list *listAddNodeTail(list *list, void *value) | 将一个包含给定值的新节点添加到给定链表的表尾 | O(1) |
| list *listInsertNode(list *list, listNode *old_node, void *value, int after) |	将一个包含给定值的新节点添加到给定节点的之前或者之后 | O(1) |
| listNode *listSearchKey(list *list, void *key) | 查找并返回链表中包含给定值的节点 | O(N) ， N 为链表长度 |
| listNode *listIndex(list *list, long index) | 返回链表在给定索引上的节点 | O(N) ， N 为链表长度 |
| void listDelNode(list *list, listNode *node) | 从链表中删除给定节点 |	O(1) |
| void listRotate(list *list) | 将链表的表尾节点弹出，然后将被弹出的节点插入到链表的表头，成为新的表头节点 | O(1) |
| list *listDup(list *orig) | 复制一个给定链表的副本 | O(N) ， N 为链表长度 |
| void listRelease(list *list) | 释放给定链表，以及链表中的所有节点 | O(N) ， N 为链表长度 |

## 重点

1. 链表被广泛用于实现 Redis 的各种功能， 比如列表键， 发布与订阅， 慢查询， 监视器， 等等。
2. 每个链表节点由一个 listNode 结构来表示， 每个节点都有一个指向前置节点和后置节点的指针， 所以 Redis 的链表实现是双端链表。
3. 个链表使用一个 list 结构来表示， 这个结构带有表头节点指针、表尾节点指针、以及链表长度等信息。
4. 因为链表表头节点的前置节点和表尾节点的后置节点都指向 NULL ， 所以 Redis 的链表实现是无环链表。
5. 通过为链表设置不同的类型特定函数， Redis 的链表可以用于保存各种不同类型的值。