/* https://leetcode.com/problems/repeated-dna-sequences/description/
All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

For example,

    Given s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",

    Return:
    ["AAAAACCCCC", "CCCCCAAAAA"].
*/

package leetcode

func findRepeatedDnaSequences(s string) []string {
	maps := map[string]int{}

	for i := 0; i <= len(s)-10; i++ {
		maps[s[i:i+10]]++
	}

	res := []string{}
	for sequence, count := range maps {
		if count > 1 {
			res = append(res, sequence)
		}
	}
	return res
}
