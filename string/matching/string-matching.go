package stringmatching

/*
字符串匹配算法主要有:
朴素算法（Naive Algorithm）
Rabin-Karp 算法
有限自动机算法（Finite Automation）
Knuth-Morris-Pratt 算法（即 KMP Algorithm）
Boyer-Moore 算法
Simon 算法
Colussi 算法
Galil-Giancarlo 算法
Apostolico-Crochemore 算法
Horspool 算法
Sunday 算法
*/

/*
朴素的字符串匹配算法（Naive String Matching Algorithm）

朴素的字符串匹配算法又称为暴力匹配算法（Brute Force Algorithm），它的主要特点是：
    1. 没有预处理阶段；
    2. 滑动窗口总是后移 1 位；
    3. 对模式中的字符的比较顺序不限定，可以从前到后，也可以从后到前；
    4. 匹配阶段需要 O((n - m + 1)m) 的时间复杂度；
    5. 需要 2n 次的字符比较；
*/
func Naive(s, txt string) bool {
	return true
}
