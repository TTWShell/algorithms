/* https://leetcode.com/problems/restore-ip-addresses/description/
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

For example:
Given "25525511135",

return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)
*/

package lbacktracking

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var helper func(s string, part int) [][]string
	helper = func(s string, part int) [][]string {
		res := [][]string{}
		if length := len(s); length < part || length > part*3 {
			return res
		}

		validate := func(s string) bool {
			if len(s) > 1 && s[0] == '0' {
				return false
			}
			if num, _ := strconv.Atoi(s); num > 255 {
				return false
			}
			return true
		}

		if part == 1 {
			if validate(s) {
				res = append(res, []string{s})
			}
		} else {
			for i := 1; i <= 3 && i <= len(s); i++ {
				strNum := s[:i]
				if !validate(strNum) {
					continue
				}
				for _, sub := range helper(s[i:], part-1) {
					res = append(res, append([]string{strNum}, sub...))
				}
			}
		}
		return res
	}

	ips := helper(s, 4)
	res := make([]string, len(ips))
	for i, ip := range ips {
		res[i] = strings.Join(ip, ".")
	}
	return res
}
