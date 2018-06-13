/* https://leetcode.com/problems/count-primes/#/description
Description:

Count the number of prime numbers less than a non-negative number, n.

质数（Prime number），又称素数，指在大于1的自然数中，除了1和该数自身外，无法被其他自然数整除的数（也可定义为只有1与该数本身两个因数的数）。

实现原理：埃拉托斯特尼筛法 https://zh.wikipedia.org/wiki/%E5%9F%83%E6%8B%89%E6%89%98%E6%96%AF%E7%89%B9%E5%B0%BC%E7%AD%9B%E6%B3%95
*/

package lmath

func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] == false {
			count++
			for j := 2; i*j < n; j++ {
				isPrimes[i*j] = true
			}
		}
	}
	return count
}
