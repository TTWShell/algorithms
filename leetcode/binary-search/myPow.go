/* https://leetcode.com/problems/powx-n/description/
Implement pow(x, n).
*/

package lbs

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPowHelper(x, -n)
	}
	return myPowHelper(x, n)
}

func myPowHelper(x float64, n int) float64 {
	// n >= 0
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return myPowHelper(x*x, n/2) * x
	} else {
		return myPowHelper(x*x, n/2)
	}
}
