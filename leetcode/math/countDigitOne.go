/* https://leetcode.com/problems/number-of-digit-one/
Given an integer n, count the total number of digit 1 appearing in
all non-negative integers less than or equal to n.

Example:

	Input: 13
	Output: 6
Explanation: Digit 1 occurred in the following numbers: 1, 10, 11, 12, 13.
*/

package lmath

/*
1的个数          含1的数字                                                                        数字范围

1                   1                                                                                     [1, 9]

11                 10  11  12  13  14  15  16  17  18  19                              [10, 19]

1                   21                                                                                   [20, 29]

1                   31                                                                                   [30, 39]

1                   41                                                                                   [40, 49]

1                   51                                                                                   [50, 59]

1                   61                                                                                   [60, 69]

1                   71                                                                                   [70, 79]

1                   81                                                                                   [80, 89]

1                   91                                                                                   [90, 99]

11                 100  101  102  103  104  105  106  107  108  109          [100, 109]

21                 110  111  112  113  114  115  116  117  118  119             [110, 119]

11                 120  121  122  123  124  125  126  127  128  129          [120, 129]
*/

func countDigitOne(n int) int {
	count := 0
	for i := 1; i <= n; i *= 10 {
		count += (n/i + 8) / 10 * i
		if n/i%10 == 1 {
			count += n%i + 1
		}
	}
	return count
}
