/* https://leetcode.com/problems/min-stack/#/description
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.getMin();   --> Returns -3.
    minStack.pop();
    minStack.top();      --> Returns 0.
    minStack.getMin();   --> Returns -2.
参考：https://discuss.leetcode.com/topic/11985/my-python-solution
*/
package leetcode

type MinStack struct {
	all [][]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.all) == 0 {
		this.all = append(this.all, []int{x, x})
	} else {
		this.all = append(this.all, []int{x, min(x, this.GetMin())})
	}
}

func (this *MinStack) Pop() {
	this.all = this.all[:len(this.all)-1]
}

func (this *MinStack) Top() int {
	return this.all[len(this.all)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.all[len(this.all)-1][1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
