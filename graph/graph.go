package graph

import (
	"errors"
	"sync"
)

// GraphType values.
const (
	Undirected int = iota
	Directed
)

// Node is a Node(node) in Graph.
type Node struct {
	Value interface{}
}

// Graph .
type Graph struct {
	sync.RWMutex
	// type of graph.
	kind int
	// Bool is used for graph traversal.
	nodes map[*Node]bool
	// Edge connects two Nodes in a graph, float32 is weight of edge.
	edges map[*Node]map[*Node]float32
	// save Directed type to-->from relation.
	reversedEdges map[*Node]map[*Node]float32
}

// New creates and returns an empty graph.
func New(kind int) *Graph {
	return &Graph{
		kind:          kind,
		nodes:         make(map[*Node]bool),
		edges:         make(map[*Node]map[*Node]float32),
		reversedEdges: make(map[*Node]map[*Node]float32),
	}
}

// MakeNode creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode(value interface{}) *Node {
	g.Lock()
	defer g.Unlock()

	newNode := &Node{Value: value}
	if _, ok := g.nodes[newNode]; !ok {
		g.nodes[newNode] = false
		g.edges[newNode] = make(map[*Node]float32)
	}
	return newNode
}

// RemoveNode removes a node from the graph and all edges connected to it.
func (g *Graph) RemoveNode(node *Node) {
	g.Lock()
	defer g.Unlock()

	if node == nil {
		return
	}
	if _, ok := g.nodes[node]; !ok {
		return
	}

	if g.kind == Undirected {
		for to := range g.edges[node] {
			delete(g.edges[to], node)
		}
	} else {
		for to := range g.reversedEdges[node] {
			delete(g.edges[to], node)
		}
		delete(g.reversedEdges, node)
	}
	delete(g.edges, node)

	delete(g.nodes, node)
}

// MakeEdge creates an edge in the graph with a corresponding weight.
// It returns an error if either of the nodes do not belong in the graph.
func (g *Graph) MakeEdge(from, to *Node, weight float32) error {
	g.Lock()
	defer g.Unlock()

	if _, ok := g.nodes[from]; !ok {
		return errors.New("`from` node does not belong to this graph")
	}
	if _, ok := g.nodes[to]; !ok {
		return errors.New("`to` node does not belong to this graph")
	}

	g.edges[from][to] = weight
	if g.kind == Undirected {
		g.edges[to][from] = weight
	} else {
		if _, ok := g.reversedEdges[to]; !ok {
			g.reversedEdges[to] = make(map[*Node]float32)
		}
		g.reversedEdges[to][from] = weight
	}

	return nil
}
