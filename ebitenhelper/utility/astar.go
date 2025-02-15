package utility

import (
	"math"
	"slices"
	"sync"
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
		n.Location.AddXY(0, -1),
		n.Location.AddXY(-1, 0),
		n.Location.AddXY(1, 0),
		n.Location.AddXY(0, 1),
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

func (a *AStarInstance) Run(start Point, goal Point) []Point {
	a.currentNode = NewAStarNode(start)
	clear(a.openedNodes)
	clear(a.closedNodes)

	for a.currentNode != nil {
		if a.currentNode.Location == goal {
			return a.GetCurrentPath()
		}

		for _, l := range a.currentNode.GetAroundLocations() {
			if !GetLevel().AIIsPFLocationValid(l) || a.closedNodes[l] != nil {
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

type AStarResultKey struct {
	Start Point
	Goal  Point
}

func NewAStarResultKey(start Point, goal Point) AStarResultKey {
	return AStarResultKey{
		Start: start,
		Goal:  goal,
	}
}

type AStarResultReason int

const (
	AStarResultReasonSucceed AStarResultReason = iota
	AStarResultReasonRequest
	AStarResultReasonFail
)

type AStar struct {
	cache            sync.Map
	runningTaskCount int
}

func NewAStar() *AStar {
	return &AStar{
		cache:            sync.Map{},
		runningTaskCount: 0,
	}
}

func (a *AStar) GetCache(start Point, goal Point) (result []Point, ok bool) {
	if r1, ok := a.cache.Load(NewAStarResultKey(start, goal)); ok {
		if r2, ok := r1.([]Point); ok {
			return r2, true
		}
	}

	return []Point{}, false
}

func (a *AStar) setCache(start Point, goal Point, value []Point) {
	a.cache.Store(NewAStarResultKey(start, goal), value)
}

func (a *AStar) RunForce(start Point, goal Point) []Point {
	res := NewAStarInstance().Run(start, goal)
	reslen := len(res)
	for i := 0; i < reslen; i++ {
		if _, ok := a.GetCache(res[i], res[reslen-1]); ok {
			break
		}
		a.setCache(res[i], res[reslen-1], res[i:])
	}

	return res
}

func (a *AStar) Run(start Point, goal Point) (result []Point, reason AStarResultReason) {
	// Found in cache
	if r, ok := a.GetCache(start, goal); ok {
		return r, AStarResultReasonSucceed
	}

	// Can't create task
	if a.runningTaskCount >= AIMaxTaskCount {
		return []Point{}, AStarResultReasonFail
	}

	// Creating task
	a.runningTaskCount++
	go func() {
		a.RunForce(start, goal)
		a.runningTaskCount--
	}()
	return []Point{}, AStarResultReasonRequest
}
