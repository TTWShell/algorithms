# 简单动态字符串(Simple Dynamic String)


## Struct定义

3.2版本之前结构体定义为：

见：https://github.com/antirez/redis/blob/3.0.7/src/sds.h

```
typedef char *sds;

struct sdshdr {
    unsigned int len; // 记录buf数组中已使用字节的数量，即sds所保存字符串的长度
    unsigned int free; // 记录buf数组中未使用字节的数量
    char buf[]; // 字节数组，用于保存字符串
};
```

3.2版本及之后，针对不同的长度范围定义了不同的结构，见：https://github.com/antirez/redis/blob/3.2.0/src/sds.h

```
typedef char *sds;

/* Note: sdshdr5 is never used, we just access the flags byte directly.
 * However is here to document the layout of type 5 SDS strings. */
struct __attribute__ ((__packed__)) sdshdr5 {
    unsigned char flags; /* 3 lsb of type, and 5 msb of string length */
    char buf[];
};
struct __attribute__ ((__packed__)) sdshdr8 {
    uint8_t len; /* used 字符串的长度 */
    uint8_t alloc; /* excluding the header and null terminator 已经分配的总长度 */
    unsigned char flags; /* 3 lsb of type, 5 unused bits lag用3bit来标明类型，其余5bit目前没有使用 */
    char buf[]; // 数组，以'\0'结尾
};
struct __attribute__ ((__packed__)) sdshdr16 {
    uint16_t len; /* used */
    uint16_t alloc; /* excluding the header and null terminator */
    unsigned char flags; /* 3 lsb of type, 5 unused bits */
    char buf[];
};
struct __attribute__ ((__packed__)) sdshdr32 {
    uint32_t len; /* used */
    uint32_t alloc; /* excluding the header and null terminator */
    unsigned char flags; /* 3 lsb of type, 5 unused bits */
    char buf[];
};
struct __attribute__ ((__packed__)) sdshdr64 {
    uint64_t len; /* used */
    uint64_t alloc; /* excluding the header and null terminator */
    unsigned char flags; /* 3 lsb of type, 5 unused bits */
    char buf[];
};
```

## SDS与C字符串的区别

C字符串: 使用N+1长度的字符数组表示字符串，且末尾一定是 `\0` 字符。

区别：

1. SDS记录了长度，优化获取长度操作时间复杂度从O(1) --> O(N);
2. 杜绝拼接字符串等操作的缓冲区溢出（buffer overflow）；
3. 减少修改字符串带来的内存分配次数；

    SDS中len、free属性带来的好处：

    1. append时，根据free是否够用（**空间预分配**），api自动决定是否需要内存重新分配，避免缓冲区溢出；
    2. trim操作时，**惰性空间释放**可以避免内存泄漏；
    3. **空间预分配**：当 `len< 1MB` 时，每次分配的 `free == len` ，当 `len >= 1MB` ，每次分配 固定为 `free == 1MB`;
    4. **惰性空间释放**: 每次操作更新free，并有释放内存的api，保证可以主动释放，不会浪费。

4. 由于记录了len，`\0` 不会影响数据的读取，二进制安全；
5. 由于同样使用 `\0` 结尾，兼容部分C字符串函数。


## 主要操作API

| 函数 | 作用 | 时间复杂度 |
| --- | --- | --- |
| sds sdsnew(const char *init); | 创建一个包含给定C字符串的SDS | O(N)，init的长度决定 |
| sds sdsempty(void); | 创建一个不包含给定C字符串的SDS | O(1) |
| void sdsfree(sds s); | 释放给定的SDS | O(N) |
| size_t sdslen(const sds s); | 返回SDS的已使用空间字节数 | O(1)，读取len属性 |
| size_t sdsavail(const sds s); | 返回SDS的未使用空间字节数 | O(1)，读取free属性 |
| sds sdsdup(const sds s); | 创建一个给定SDS的副本（copy） | O(N) |
| void sdsclear(sds s); | 清空SDS保存的字符串内容 | O(1)，惰性释放 |
| sds sdscat(sds s, const char *t); | 将给定C字符串拼接到SDS字符串的末尾 | O(N) |
| sds sdscatsds(sds s, const sds t); | 将给定SDS字符串拼接到另一个SDS字符串的末尾 | O(N) |
| sds sdscpy(sds s, const char *t); | 将给定C字符串复制到SDS里面，覆盖SDS原有的字符串 | O(N) |
| sds sdsgrowzero(sds s, size_t len); | 用空字符串将SDS扩展至给定的长度 | O(N) |
| void sdsrange(sds s, int start, int end); | 保留SDS给定区间内的数据，不在区间内的数据会被覆盖或清除 | O(N) |
| sds sdstrim(sds s, const char *cset); | 接受一个SDS和一个C字符串作为参数，从SD中移除所有在C字符串中出现过的字符 | O(N^2) |
| int sdscmp(const sds s1, const sds s2); | 对比两个SDS字符串是否相同 | O(N) |