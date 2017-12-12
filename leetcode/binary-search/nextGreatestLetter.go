/* https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

Examples:
	Input:
	letters = ["c", "f", "j"]
	target = "a"
	Output: "c"

	Input:
	letters = ["c", "f", "j"]
	target = "c"
	Output: "f"

	Input:
	letters = ["c", "f", "j"]
	target = "d"
	Output: "f"

	Input:
	letters = ["c", "f", "j"]
	target = "g"
	Output: "j"

	Input:
	letters = ["c", "f", "j"]
	target = "j"
	Output: "c"

	Input:
	letters = ["c", "f", "j"]
	target = "k"
	Output: "c"
Note:
    1. letters has a length in range [2, 10000].
    2. letters consists of lowercase letters, and contains at least 2 unique letters.
    3. target is a lowercase letter.
*/

package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	start, end := 0, len(letters)-1
	if target < letters[0] || target >= letters[end] {
		return letters[start]
	}

	var mid int
	for start < end {
		mid = start + (end+1-start)/2
		if letters[mid] < target {
			start = mid
		} else if letters[mid] == target {
			for letters[mid] == target {
				mid++
			}
			break
		} else if mid == end {
			break
		} else {
			end = mid
		}
	}
	return letters[mid]
}
