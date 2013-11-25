// History: Nov 22 13 tcolar Creation

package algo

import ()

// A graph (nodes & edges)
type Graph interface {
	Nodes() []Node
	Edges() []Edge
}

// A graph node
type Node interface {
	NdId() int
}

// Edge between two nodes
type Edge interface {
	NdFrom() int
	NdTo() int
}
