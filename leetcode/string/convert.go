/* https://leetcode.com/problems/zigzag-conversion/#/description
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

	P   A   H   N
	A P L S I I G
	Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

    string convert(string text, int nRows);

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
*/

package leetcode

func convert(s string, numRows int) string {
	// 第一行和最后一行下标间隔都是interval = numRows*2-2 = 8
	// 中间行的间隔是周期性的,第i行的间隔是: interval–2*i,  2*i,  interval–2*i, 2*i, interval–2*i, 2*i, …
	length := len(s)
	if numRows == 1 {
		return s
	}

	interval := numRows<<1 - 2
	result, index := make([]byte, len(s), len(s)), 0
	// 处理第一行
	for i := 0; i < length; i += interval {
		result[index] = s[i]
		index++
	}

	// 处理中间行
	for i := 1; i < numRows-1; i++ {
		inter := i << 1
		for j := i; j < length; j += inter {
			{
				result[index] = s[j]
				index++
				inter = interval - inter
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < length; i += interval {
		result[index] = s[i]
		index++
	}
	return string(result)
}
