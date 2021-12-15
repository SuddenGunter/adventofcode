package graph

import "errors"

const maxAdj = 4

type Node struct {
	ID       Position
	Adjacent []*Node
	Weight   int
}

func (n *Node) Equals(node *Node) bool {
	return n.ID == node.ID
}

type Graph struct {
	Nodes map[Position]*Node
}

func (g *Graph) Add(id Position, weight int) (*Graph, error) {
	_, found := g.Nodes[id]
	if found {
		return nil, errors.New("node already exists")
	}

	adj := g.getAllAdjacent(id)

	g.Nodes[id] = &Node{
		ID:       id,
		Adjacent: adj,
		Weight:   weight,
	}

	return g, nil
}

func (g *Graph) getAllAdjacent(id Position) []*Node {
	appendIfFound := func(dst []*Node, pos Position, all map[Position]*Node) {
		node, found := all[pos]
		if !found {
			return
		}

		dst = append(dst, node)
	}

	nodes := make([]*Node, 0, maxAdj)

	appendIfFound(nodes, Position{
		X: id.X,
		Y: id.Y + 1,
	}, g.Nodes)
	appendIfFound(nodes, Position{
		X: id.X,
		Y: id.Y - 1,
	}, g.Nodes)
	appendIfFound(nodes, Position{
		X: id.X + 1,
		Y: id.Y,
	}, g.Nodes)
	appendIfFound(nodes, Position{
		X: id.X - 1,
		Y: id.Y,
	}, g.Nodes)

	return nodes
}

// Position is a position of the node in the original input array.
type Position struct {
	X, Y int
}
