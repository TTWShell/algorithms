# 栈（stack）

栈在计算机科学中，是一种特殊的串列形式的数据结构，它的特殊之处在于只能允许在链接串列或阵列的一端（称为堆叠顶端指标 top）进行加入数据（push）和输出数据（pop）的运算。

![栈](https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Data_stack.svg/400px-Data_stack.svg.png)

由于堆叠数据结构只允许在一端进行操作，因而按照后进先出（LIFO, Last In First Out）的原理运作。

[stack的golang实现](https://github.com/TTWShell/algorithms/blob/master/stack/stack.go)

## 栈的应用

### [平衡符号](https://github.com/TTWShell/algorithms/blob/master/stack/balanceSymbol.go)

应用场景：

检测括号是否能够成对出现，如序列`[()]`是合法的，`[(])`是非法的。

实现思路：

创建一个空栈，读入所有字符，如果字符是一个开放符号，则入栈；
如果是一个封闭符号，则当栈空时报错，否则，将栈弹出；
如果弹出的符号不是对应的开放符号，则报错。
在字符末尾，如果栈非空，则报错。
