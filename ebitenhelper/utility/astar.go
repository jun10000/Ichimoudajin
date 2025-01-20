package utility

import (
	"math"
	"slices"
)

type AStarNode struct {
	Location  Point
	IsAllInit bool
	GDistance float64
	HDistance float64
	Parent    *AStarNode
}

func NewAStarNode(location Point) *AStarNode {
	return &AStarNode{Location: location}
}

func (n *AStarNode) GetAroundLocations() []Point {
	return []Point{
		n.Location.AddXY(-1, -1),
		n.Location.AddXY(0, -1),
		n.Location.AddXY(1, -1),
		n.Location.AddXY(-1, 0),
		n.Location.AddXY(1, 0),
		n.Location.AddXY(-1, 1),
		n.Location.AddXY(0, 1),
		n.Location.AddXY(1, 1),
	}
}

type AStar struct {
	IsLocationValidCB func(location Point) bool

	currentNode *AStarNode
	openedNodes map[Point]*AStarNode
	closedNodes map[Point]*AStarNode
}

func NewAStar(islocvalidcb func(location Point) bool) *AStar {
	return &AStar{
		IsLocationValidCB: islocvalidcb,
		openedNodes:       map[Point]*AStarNode{},
		closedNodes:       map[Point]*AStarNode{},
	}
}

func (a *AStar) Run(start Point, goal Point) (result []Point, ok bool) {
	a.currentNode = NewAStarNode(start)

	for a.currentNode != nil {
		if a.currentNode.Location == goal {
			return a.GetCurrentPath(), true
		}

		for _, l := range a.currentNode.GetAroundLocations() {
			if !a.IsLocationValidCB(l) || a.closedNodes[l] != nil {
				continue
			}

			if a.openedNodes[l] == nil {
				a.openedNodes[l] = NewAStarNode(l)
			}
			a.UpdateNode(a.openedNodes[l], goal)
		}

		a.closedNodes[a.currentNode.Location] = a.currentNode
		delete(a.openedNodes, a.currentNode.Location)
		a.currentNode = a.GetNextOpenNode()
	}

	return nil, false
}

func (a *AStar) UpdateNode(node *AStarNode, goal Point) {
	gd := a.currentNode.GDistance + a.currentNode.Location.Distance(node.Location)
	if !node.IsAllInit {
		node.GDistance = gd
		node.HDistance = node.Location.Distance(goal)
		node.Parent = a.currentNode
		node.IsAllInit = true
	} else {
		if gd < node.GDistance {
			node.GDistance = gd
			node.Parent = a.currentNode
		} /* else {
		}*/
	}
}

func (a *AStar) GetNextOpenNode() *AStarNode {
	min := math.MaxFloat64
	var r *AStarNode

	for _, n := range a.openedNodes {
		if !n.IsAllInit {
			continue
		}

		gh := n.GDistance + n.HDistance
		if gh < min {
			min = gh
			r = n
		}
	}

	return r
}

func (a *AStar) GetCurrentPath() []Point {
	r := []Point{a.currentNode.Location}
	n := a.currentNode.Parent

	for n != nil {
		r = append(r, n.Location)
		n = n.Parent
	}

	slices.Reverse(r)
	return r
}
