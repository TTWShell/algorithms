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

package leetcode

import "strconv"

type StringIterator struct {
	all []byte
}

func Constructor(compressedString string) StringIterator {
	var all []byte
	for i := 0; i < len(compressedString); {
		if letter := compressedString[i]; letter-'A' > 0 {
			j := i + 1
			for j < len(compressedString) && compressedString[j]-'0' < 10 {
				j++
			}
			repeat, _ := strconv.Atoi(string(compressedString[i+1 : j]))
			for c := 0; c < repeat; c++ {
				all = append(all, letter)
			}
			i = j
		}
	}
	iterator := StringIterator{all: all}
	return iterator
}

func (this *StringIterator) Next() byte {
	if len(this.all) > 0 {
		r := this.all[0]
		this.all = this.all[1:]
		return r
	}
	return ' '
}

func (this *StringIterator) HasNext() bool {
	return len(this.all) > 0
}
