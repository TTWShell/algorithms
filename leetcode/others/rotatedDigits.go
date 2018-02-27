/* https://leetcode.com/problems/rotated-digits/description/
X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.
A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves;
2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:

    Input: 10
    Output: 4
    Explanation:
    There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
    Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.

Note:

    N  will be in range [1, 10000].
*/

package leetcode

func rotatedDigits(N int) int {
	check := func(num int) bool {
		res := false
		for num > 0 {
			if tmp := num % 10; tmp == 2 || tmp == 5 || tmp == 6 || tmp == 9 {
				res = true
			} else if tmp == 3 || tmp == 4 || tmp == 7 {
				return false
			}
			num /= 10
		}
		return res
	}

	res := 0
	for i := 2; i <= N; i++ {
		if check(i) {
			res++
		}
	}
	return res
}
