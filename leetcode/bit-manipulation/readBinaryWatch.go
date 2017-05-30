/* https://leetcode.com/problems/binary-watch/#/description
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.

https://upload.wikimedia.org/wikipedia/commons/8/8b/Binary_clock_samui_moon.jpg

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

    Input: n = 1
    Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
    The order of output does not matter.
    The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
    The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
*/

package leetcode

import "fmt"

func readBinaryWatch(num int) []string {
	H := [][]int{
		{0},
		{1, 2, 4, 8},
		{3, 5, 9, 6, 10},
		{7, 11},
	}
	S := [][]int{
		{0},
		{1, 2, 4, 8, 16, 32},
		{3, 5, 9, 17, 33, 6, 10, 18, 34, 12, 20, 36, 24, 40, 48},
		{7, 11, 19, 35, 14, 22, 38, 13, 28, 44, 26, 56, 49, 50, 42, 21, 25, 37, 41, 52},
		{15, 23, 39, 30, 46, 29, 57, 58, 27, 53, 45, 43, 51, 54},
		{31, 47, 55, 59},
	}

	r := []string{}

	for i := 0; i <= num; i++ {
		hour, second := 0, 0
		for h := 0; i < 4 && h < len(H[i]); h++ {
			hour = H[i][h]
			for s := 0; num-i < 6 && s < len(S[num-i]); s++ {
				second = S[num-i][s]
				r = append(r, fmt.Sprintf("%d:%02d", hour, second))
			}
		}

	}

	return r
}
