package bit

func lowbit(x int) int {
	return x & (-x)
}

// BIT is Binary Indexed Tree.
type BIT struct {
	c   []int
	len int
}

// New func is constructor of BIT struct.
func New(nums []int) *BIT {
	len := len(nums) + 1
	c := make([]int, len, len)
	for i := 1; i < len; i++ {
		c[i] = nums[i-1]
		for j := i - 2; j >= i-lowbit(i); j-- {
			c[i] += nums[j]
		}
	}
	return &BIT{c: c, len: len}
}

// Update nums[i], add delta.
func (b *BIT) Update(i int, delta int) {
	for j := i + 1; j < b.len; j += lowbit(j) {
		b.c[j] += delta
		if lowbit(j) == 0 {
			break
		}
	}
}

// Sum 1--i.
func (b *BIT) Sum(i int) int {
	res := 0
	for j := i; j > 0; j -= lowbit(j) {
		res += b.c[j]
	}
	return res
}
