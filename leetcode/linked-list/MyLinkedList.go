/* https://leetcode.com/problems/design-linked-list/
Design your implementation of the linked list.
You can choose to use the singly linked list or the doubly linked list.
A node in a singly linked list should have two attributes: val and next.
val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list.
Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

	get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
	addAtHead(val) : Add a node of value val before the first element of the linked list.
		After the insertion, the new node will be the first node of the linked list.
	addAtTail(val) : Append a node of value val to the last element of the linked list.
	addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list.
		If index equals to the length of linked list, the node will be appended to the end of linked list.
		If index is greater than the length, the node will not be inserted.
	deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.

Example:

	MyLinkedList linkedList = new MyLinkedList();
	linkedList.addAtHead(1);
	linkedList.addAtTail(3);
	linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
	linkedList.get(1);            // returns 2
	linkedList.deleteAtIndex(1);  // now the linked list is 1->3
	linkedList.get(1);            // returns 3
Note:

	All values will be in the range of [1, 1000].
	The number of operations will be in the range of [1, 1000].
	Please do not use the built-in LinkedList library.
*/

package lll

type MyLinkedList struct {
	size       int
	head, tail *node
}

type node struct {
	val  int
	next *node
}

// Constructor Initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	res := this.head
	for ; index > 0; index-- {
		res = res.next
	}
	return res.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &node{val: val}
	if this.size == 0 {
		this.head, this.tail = head, head
	} else {
		head.next = this.head
		this.head = head
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.size == 0 {
		this.AddAtHead(val)
	} else {
		tail := &node{val: val}
		this.tail.next = tail
		this.tail = tail
		this.size++
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	switch index {
	case 0:
		this.AddAtHead(val)
	case this.size:
		this.AddAtTail(val)
	default:
		cur := this.head
		for ; index > 1; index-- { // find parent
			cur = cur.next
		}
		cur.next = &node{val: val, next: cur.next}
		this.size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	if index == 0 {
		if this.size == 1 {
			this.head, this.tail = nil, nil
		} else {
			this.head = this.head.next
		}
	} else {
		cur := this.head
		for ; index > 1; index-- { // find parent
			cur = cur.next
		}
		if cur.next == this.tail {
			cur.next = nil
			this.tail = cur
		} else {
			cur.next = cur.next.next
		}
	}
	this.size--
}
