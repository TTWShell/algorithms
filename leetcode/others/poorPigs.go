/* https://leetcode.com/problems/poor-pigs/description/
There are 1000 buckets, one and only one of them contains poison, the rest are filled with water.
They all look the same. If a pig drinks that poison it will die within 15 minutes.
What is the minimum amount of pigs you need to figure out which bucket contains the poison within one hour.

Answer this question, and write an algorithm for the follow-up general case.

Follow-up:

If there are n buckets and a pig drinking poison will die within m minutes,
how many pigs (x) you need to figure out the "poison" bucket within p minutes? There is exact one bucket with poison.
*/
package lothers

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 1 {
		return 0
	}

	pow := func(a, b int) int {
		res := 1
		for ; b > 0; b-- {
			res = res * a
		}
		return res
	}

	res, a := 1, minutesToTest/minutesToDie+1
	for pow(a, res) < buckets {
		res++
	}
	return res
}
