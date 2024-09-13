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

package lstack

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.stack1 = append(m.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	m.Peek()
	top := len(m.stack2) - 1
	value := m.stack2[len(m.stack2)-1]
	m.stack2 = m.stack2[:top]
	return value
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if len(m.stack2) == 0 {
		for len(m.stack1) > 0 {
			top := len(m.stack1) - 1
			m.stack2 = append(m.stack2, m.stack1[top])
			m.stack1 = m.stack1[:top]
		}
	}
	return m.stack2[len(m.stack2)-1]
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	return len(m.stack1) == 0 && len(m.stack2) == 0
}
