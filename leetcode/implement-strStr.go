/* https://leetcode.com/problems/implement-strstr/#/description
Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
note: 判断一个字符串是否是另一个字符串的子串。
*/

package leetcode

func strStr(haystack string, needle string) int {
	lenh, lenn := len(haystack), len(needle)
	if lenn > lenh {
		return -1
	}
	for i := 0; i <= lenh-lenn; i++ {
		for j := 0; j <= lenn; j++ {
			if j == lenn {
				return i
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
	return -1
}
