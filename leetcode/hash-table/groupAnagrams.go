/* https://leetcode.com/problems/group-anagrams/description/
Given an array of strings, group anagrams together.

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:

    [
      ["ate", "eat","tea"],
      ["nat","tan"],
      ["bat"]
    ]
Note: All inputs will be in lower-case.
*/

package leetcode

import (
	"bytes"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	helper := func(word string) string {
		maps := make([]int, 26, 26)
		for _, w := range word {
			maps[w-'a']++
		}
		buffer := bytes.Buffer{}
		for i, c := range maps {
			if c > 0 {
				buffer.WriteRune(rune('a' + i))
				buffer.WriteString(strconv.Itoa(c))
			}
		}
		return buffer.String()
	}

	maps := make(map[string][]string)
	for _, str := range strs {
		key := helper(str)
		if _, ok := maps[key]; ok {
			maps[key] = append(maps[key], str)
		} else {
			maps[key] = []string{str}
		}
	}

	res, index := make([][]string, len(maps), len(maps)), 0
	for _, v := range maps {
		res[index] = v
		index++
	}
	return res
}
