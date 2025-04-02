package utility

import (
	"compress/gzip"
	"encoding/gob"
	"math"
	"os"
	"slices"

	"maps"

	"github.com/jun10000/Ichimoudajin/assets"
)

type AStarNode struct {
	Location  Point
	IsAllInit bool
	GDistance int
	HDistance int
	Parent    *AStarNode
}

func NewAStarNode(location Point) *AStarNode {
	return &AStarNode{Location: location}
}

func (n *AStarNode) GetAroundLocations() func(yield func(Point) bool) {
	return func(yield func(Point) bool) {
		if !yield(n.Location.AddXY(0, -1)) {
			return
		}
		if !yield(n.Location.AddXY(-1, 0)) {
			return
		}
		if !yield(n.Location.AddXY(1, 0)) {
			return
		}
		if !yield(n.Location.AddXY(0, 1)) {
			return
		}
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
	isValid := GetLevel().AIIsPFLocationValid
	if !isValid(start) || !isValid(goal) {
		return []Point{}
	}

	a.currentNode = NewAStarNode(start)
	clear(a.openedNodes)
	clear(a.closedNodes)

	for a.currentNode != nil {
		if a.currentNode.Location == goal {
			return a.GetCurrentPath()
		}

		for l := range a.currentNode.GetAroundLocations() {
			if !isValid(l) || a.closedNodes[l] != nil {
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
	gd := a.currentNode.GDistance + a.currentNode.Location.Distance2(node.Location)
	if !node.IsAllInit {
		node.GDistance = gd
		node.HDistance = node.Location.Distance2(goal)
		node.Parent = a.currentNode
		node.IsAllInit = true
	} else if gd < node.GDistance {
		node.GDistance = gd
		node.Parent = a.currentNode
	}
}

func (a *AStarInstance) GetNextOpenNode() *AStarNode {
	min := math.MaxInt
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
	r := make([]Point, 0, InitialPFResultCap)
	r = append(r, a.currentNode.Location)
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

type AStar struct {
	cache     *Smap[AStarResultKey, []Point]
	taskCount int
}

func NewAStar() *AStar {
	return &AStar{
		cache:     NewSmap[AStarResultKey, []Point](),
		taskCount: 0,
	}
}

func (a *AStar) GetCache(start Point, goal Point) (result []Point, ok bool) {
	return a.cache.Load(NewAStarResultKey(start, goal))
}

func (a *AStar) setCache(start Point, goal Point, value []Point) {
	a.cache.Store(NewAStarResultKey(start, goal), value)
}

func (a *AStar) LoadCache(filename string) error {
	f, err := assets.Assets.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	g, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	defer g.Close()
	d := gob.NewDecoder(g)
	m := map[AStarResultKey][]Point{}
	err = d.Decode(&m)
	if err != nil {
		return err
	}

	a.cache.Clear()
	for k, v := range m {
		a.cache.Store(k, v)
	}

	return nil
}

func (a *AStar) SaveCache(filename string) error {
	f, err := os.Create("assets/" + filename)
	if err != nil {
		return err
	}

	defer f.Close()
	g := gzip.NewWriter(f)

	defer g.Close()
	e := gob.NewEncoder(g)
	m := maps.Collect(a.cache.Range())
	err = e.Encode(m)
	if err != nil {
		return err
	}

	return nil
}

func (a *AStar) GetResultForce(start Point, goal Point) []Point {
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

func (a *AStar) GetResult(start Point, goal Point) (result []Point, ok bool) {
	// Found in cache
	if r, ok := a.GetCache(start, goal); ok {
		return r, true
	}

	// Can't create task
	if a.taskCount >= AIMaxTaskCount {
		return []Point{}, false
	}

	// Creating task
	a.taskCount++
	go func() {
		a.GetResultForce(start, goal)
		a.taskCount--
	}()
	return []Point{}, false
}
