package ltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorder(t *testing.T) {
	assert := assert.New(t)

	// [1,null,3,2,4,null,5,6]
	root :=
		&Node{
			Val: 1,
			Children: []*Node{&Node{
				Val: 3,
				Children: []*Node{&Node{
					Val:      5,
					Children: []*Node{},
				}, &Node{
					Val:      6,
					Children: []*Node{},
				}},
			}, &Node{
				Val:      2,
				Children: []*Node{},
			}, &Node{
				Val:      4,
				Children: []*Node{},
			}},
		}
	assert.Equal([]int{5, 6, 3, 2, 4, 1}, postorder(root))
}
