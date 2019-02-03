/* https://leetcode.com/problems/to-lower-case/
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

Example 1:

	Input: "Hello"
	Output: "hello"

Example 2:

	Input: "here"
	Output: "here"

Example 3:

	Input: "LOVELY"
	Output: "lovely"
*/

package lstring

func toLowerCase(str string) string {
	res := []byte(str)
	for i := 0; i < len(res); i++ {
		if 'A' <= res[i] && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
	}
	return string(res)
}
