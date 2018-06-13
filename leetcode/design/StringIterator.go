/* https://leetcode.com/problems/design-compressed-string-iterator/#/description
Design and implement a data structure for a compressed string iterator. It should support the following operations: next and hasNext.

The given compressed string will be in the form of each letter followed by a positive integer representing the number of this letter existing in the original uncompressed string.

next() - if the original string still has uncompressed characters, return the next letter; Otherwise return a white space.
hasNext() - Judge whether there is any letter needs to be uncompressed.

Note:
    Please remember to RESET your class variables declared in StringIterator, as static/class variables are persisted across multiple test cases. Please see here for more details.

Example:

    StringIterator iterator = new StringIterator("L1e2t1C1o1d1e1");

    iterator.next(); // return 'L'
    iterator.next(); // return 'e'
    iterator.next(); // return 'e'
    iterator.next(); // return 't'
    iterator.next(); // return 'C'
    iterator.next(); // return 'o'
    iterator.next(); // return 'd'
    iterator.hasNext(); // return true
    iterator.next(); // return 'e'
    iterator.hasNext(); // return false
    iterator.next(); // return ' '
*/

package ldesign

import "strconv"

type StringIterator struct {
	String    string
	nextIndex int
	curCount  int
	curLetter byte
	curTotal  int
	hasNext   bool
}

func Constructor(compressedString string) StringIterator {
	iterator := StringIterator{String: compressedString}
	if len(compressedString) > 0 {
		iterator.hasNext = true
	}
	return iterator
}

func (this *StringIterator) Next() byte {
	if this.hasNext == false {
		return ' '
	}

	this.curCount++
	r := this.curLetter

	if this.curCount >= this.curTotal &&
		this.nextIndex >= len(this.String)-1 {
		this.hasNext = false
	}

	// 第一次取或者需要取下一个字母
	if this.curCount > this.curTotal {
		for i := this.nextIndex; i < len(this.String); {
			letter, j := this.String[i], i+1
			for j < len(this.String) && this.String[j]-'0' < 10 {
				j++
			}
			this.curTotal, _ = strconv.Atoi(string(this.String[i+1 : j]))
			this.curCount, this.curLetter, this.nextIndex, this.hasNext = 1, letter, j, true
			return letter
		}
	}
	return r
}

func (this *StringIterator) HasNext() bool {
	return this.hasNext
}
