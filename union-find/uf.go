package unionfind

import "sync"

type ufNode struct {
	value            interface{}
	ancestor, parent *ufNode
	childs           []*ufNode
}

// UF : Union-Find.
type UF struct {
	sync.RWMutex
	// map all value to node and, for find method
	nodes map[interface{}]*ufNode
	// save trees, for count method, int value is count of node in set
	ufSets map[*ufNode]int
}

// New : create a empty Union-Find.
func New() *UF {
	return &UF{nodes: make(map[interface{}]*ufNode, 0), ufSets: make(map[*ufNode]int, 0)}
}

// Find : return ancestor's value.
func (uf *UF) Find(value interface{}) interface{} {
	uf.RLock()
	defer uf.RUnlock()

	ufN, ok := uf.nodes[value]
	if ok {
		return ufN.ancestor.value
	}
	return nil
}

// Connected : return true if p ∈ setA && q ∈ setA else false.
func (uf *UF) Connected(p, q interface{}) bool {
	uf.RLock()
	defer uf.RUnlock()

	pAnc, qAnc := uf.Find(p), uf.Find(q)
	if pAnc != nil && qAnc != nil && pAnc == qAnc {
		return true
	}
	return false
}

// Count : count of disjoint sets.
func (uf *UF) Count() int {
	uf.RLock()
	defer uf.RUnlock()

	return len(uf.ufSets)
}

// transposition: node --> root & set ancestor.
func (uf *UF) transposition(node, newAnc *ufNode) {
	// first find parent node.
	var anc, parent *ufNode
	var cIdx int
	anc = node.ancestor
	queue := []*ufNode{anc}
	for len(queue) != 0 {
		nextQueue := []*ufNode{}
		for _, p := range queue {
			for i, c := range p.childs {
				if parent == nil && c == node {
					parent = p
					cIdx = i
				}
			}
			p.ancestor = newAnc
			nextQueue = append(nextQueue, p.childs...)
		}
		queue = nextQueue
	}

	if parent != nil {
		parent.childs = append(append([]*ufNode{}, parent.childs[0:cIdx]...), parent.childs[cIdx+1:]...)
		node.childs = append(node.childs, parent)

		ago, cur, parent := node, parent, parent.parent
		cur.parent = node
		for cur != anc {
			idx := -1
			for i, n := range parent.childs {
				if n == cur {
					idx = i
					break
				}
			}
			parent.childs = append(append([]*ufNode{}, parent.childs[0:idx]...), parent.childs[idx+1:]...)
			cur.childs = append(cur.childs, parent)

			ago, cur, parent = cur, parent, parent.parent
			cur.parent = ago
		}
		cur.parent = ago
	}
}

// assume qNode merge to pNode.
func (uf *UF) union(pNode, qNode *ufNode) {
	pAnc, qAnc := pNode.ancestor, qNode.ancestor
	uf.ufSets[pAnc] += uf.ufSets[qAnc]
	delete(uf.ufSets, qAnc)

	// transposition qNode tree.
	uf.transposition(qNode, pAnc)

	// conenct pNode and qNode.
	pNode.childs = append(pNode.childs, qNode)
	qNode.parent = pNode
}

func (uf *UF) getOrCreateNode(value interface{}) *ufNode {
	node, ok := uf.nodes[value]
	if !ok {
		node = &ufNode{value: value, childs: make([]*ufNode, 0)}
		node.ancestor = node
		node.parent = node
		uf.nodes[value] = node
		uf.ufSets[node] = 1
	}
	return node
}

// Union : add a connection for p and q.
func (uf *UF) Union(p, q interface{}) {
	if uf.Connected(p, q) {
		return
	}

	uf.Lock()
	defer uf.Unlock()

	if p == q {
		uf.getOrCreateNode(p)
		return
	}

	pNode, qNode := uf.getOrCreateNode(p), uf.getOrCreateNode(q)
	countP, _ := uf.ufSets[pNode.ancestor]
	countQ, _ := uf.ufSets[qNode.ancestor]
	if countP >= countQ {
		uf.union(pNode, qNode)
	} else {
		uf.union(qNode, pNode)
	}
}
