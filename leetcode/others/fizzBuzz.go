/* https://leetcode.com/problems/fizz-buzz/#/description
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/

package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	r := make([]string, n, n)
	for i := 1; i <= n; i++ {
		m, n := i%3, i%5
		if m == 0 && n != 0 {
			r[i-1] = "Fizz"
		} else if m != 0 && n == 0 {
			r[i-1] = "Buzz"
		} else if m == 0 && n == 0 {
			r[i-1] = "FizzBuzz"
		} else {
			r[i-1] = strconv.Itoa(i)
		}
	}
	return r
}
