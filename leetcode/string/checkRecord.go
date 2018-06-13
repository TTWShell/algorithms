/* https://leetcode.com/problems/student-attendance-record-i/#/description
You are given a string representing an attendance record for a student. The record only contains the following three characters:

    'A' : Absent.
    'L' : Late.
    'P' : Present.

A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

You need to return whether the student could be rewarded according to his attendance record.

Example 1:
    Input: "PPALLP"
    Output: True

Example 2:
    Input: "PPALLL"
    Output: False
*/

package lstring

func checkRecord(s string) bool {
	countA, countL := 0, 0
	for _, c := range s {
		switch c {
		case 'A':
			countA++
			if countA > 1 {
				return false
			}
			countL = 0
		case 'L':
			countL++
			if countL > 2 {
				return false
			}
		case 'P':
			countL = 0
		}
	}
	return true
}
