/* https://leetcode.com/problems/repeated-dna-sequences/description/
All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

For example,

    Given s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",

    Return:
    ["AAAAACCCCC", "CCCCCAAAAA"].
*/

package lht

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	maps := map[int]int{}
	// Huffman Coding
	// 8bit --> 2bit: 00 , 01, 10, 11
	// 10 * 8 = 80bit(string) --> 10 * 2 = 20bit < 32bit(int)
	coding := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	code := 0
	for i := 0; i < 10; i++ {
		code <<= 2
		code |= coding[s[i]]
	}
	maps[code]++

	res := []string{}
	for i := 10; i < len(s); i++ {
		// curSubStr is start at index i-9, not i
		code <<= 2
		code |= coding[s[i]]
		// bin(3145728) --> '11' + '0' * 20
		code &= ^3145728

		if cnt, ok := maps[code]; ok && cnt == 1 {
			res = append(res, s[i-9:i+1])
		}
		maps[code]++
	}
	return res
}
