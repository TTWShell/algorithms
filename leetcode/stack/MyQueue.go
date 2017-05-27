/* https://leetcode.com/problems/implement-queue-using-stacks/#/description
Implement the following operations of a queue using stacks.

    push(x) -- Push element x to the back of queue.
    pop() -- Removes the element from in front of queue.
    peek() -- Get the front element.
    empty() -- Return whether the queue is empty.

Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).
*/

package leetcode

type MyQueue struct {
	q []int
}

/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	// 可以考虑两个 []int 提高push性能，代价是pop／peek 性能降低
	this.q = append([]int{x}, this.q...)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	r := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return r
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.q[len(this.q)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.q) == 0
}
