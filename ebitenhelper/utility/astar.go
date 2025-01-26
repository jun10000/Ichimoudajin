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

type AStarInstance struct {
	currentNode *AStarNode
	openedNodes map[Point]*AStarNode
	closedNodes map[Point]*AStarNode
}

func NewAStarInstance() *AStarInstance {
	return &AStarInstance{
		openedNodes: map[Point]*AStarNode{},
		closedNodes: map[Point]*AStarNode{},
	}
}

func (a *AStarInstance) Run(start Point, goal Point, isLocationValid func(location Point) bool) []Point {
	a.currentNode = NewAStarNode(start)
	clear(a.openedNodes)
	clear(a.closedNodes)

	for a.currentNode != nil {
		if a.currentNode.Location == goal {
			return a.GetCurrentPath()
		}

		for _, l := range a.currentNode.GetAroundLocations() {
			if !isLocationValid(l) || a.closedNodes[l] != nil {
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

	return []Point{}
}

func (a *AStarInstance) UpdateNode(node *AStarNode, goal Point) {
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

func (a *AStarInstance) GetNextOpenNode() *AStarNode {
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

func (a *AStarInstance) GetCurrentPath() []Point {
	r := []Point{a.currentNode.Location}
	n := a.currentNode.Parent

	for n != nil {
		r = append(r, n.Location)
		n = n.Parent
	}

	slices.Reverse(r)
	return r
}

type AStar struct {
	isRunning bool
	savedPath []Point
}

func NewAStar() *AStar {
	return &AStar{
		savedPath: []Point{},
	}
}

func (a *AStar) Run(start Point, goal Point, isLocationValid func(location Point) bool) []Point {
	if a.isRunning {
		return a.GetPath(start)
	}
	a.isRunning = true

	go func(fs Point, fg Point, ff func(location Point) bool) {
		a.savedPath = NewAStarInstance().Run(fs, fg, ff)
		a.isRunning = false
	}(start, goal, isLocationValid)

	return a.GetPath(start)
}

func (a *AStar) GetPath(start Point) []Point {
	s := slices.Index(a.savedPath, start)
	if s == -1 {
		return a.savedPath
	}

	return a.savedPath[s:]
}
