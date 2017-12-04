package unionfind

import "sync"

type ufNode struct {
	value    interface{}
	ancestor *ufNode
	childs   []*ufNode
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
	var parent *ufNode
	var cIdx int
	if node == node.ancestor {
		parent = nil
	}
	queue := []*ufNode{node.ancestor}
	for len(queue) != 0 {
		nextQueue := []*ufNode{}
		for _, p := range queue {
			for i, c := range p.childs {
				if c == node {
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
	}
}

// assume qNode merge to pNode.
func (uf *UF) union(pNode, qNode *ufNode) {
	pAnc, qAnc := pNode.ancestor, qNode.ancestor
	uf.ufSets[pAnc] += uf.ufSets[qNode.ancestor]

	// transposition qNode tree.
	uf.transposition(qNode, pAnc)

	// conenct pNode and qNode.
	pNode.childs = append(pNode.childs, qNode)
	delete(uf.ufSets, qAnc)
}

// Union : add a connection for p and q.
func (uf *UF) Union(p, q interface{}) {
	if uf.Connected(p, q) {
		return
	}

	uf.Lock()
	defer uf.Unlock()

	var pNode, qNode *ufNode
	if node, ok := uf.nodes[p]; !ok {
		pNode = &ufNode{value: p, childs: make([]*ufNode, 0)}
		pNode.ancestor = pNode
		uf.nodes[p] = pNode
		uf.ufSets[pNode] = 1
	} else {
		pNode = node
	}
	if node, ok := uf.nodes[q]; !ok {
		qNode = &ufNode{value: q, childs: make([]*ufNode, 0)}
		qNode.ancestor = qNode
		uf.nodes[q] = qNode
		uf.ufSets[qNode] = 1
	} else {
		qNode = node
	}

	countP, _ := uf.ufSets[pNode.ancestor]
	countQ, _ := uf.ufSets[qNode.ancestor]
	if countP >= countQ {
		uf.union(pNode, qNode)
	} else {
		uf.union(qNode, pNode)
	}
}
