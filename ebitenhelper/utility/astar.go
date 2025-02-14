package utility

import (
	"context"
	"log"
	"math"
	"os"
	"os/signal"
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

type AStarRequest struct {
	Request    AStarResultKey
	ResponseCh chan []Point
}

func NewAStarRequest(start Point, goal Point, responseCh chan []Point) AStarRequest {
	return AStarRequest{
		Request:    NewAStarResultKey(start, goal),
		ResponseCh: responseCh,
	}
}

type AStar struct {
	Finish context.CancelFunc

	requestCh chan AStarRequest
}

func StartAStar() *AStar {
	a := &AStar{
		requestCh: make(chan AStarRequest, AIRequestCap),
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)
	ctx, cancel := context.WithCancel(context.Background())
	a.Finish = cancel

	// Run goroutine
	go a.loop(signalCh, ctx)

	return a
}

// Loop uses goroutine
func (a *AStar) loop(signalCh <-chan os.Signal, ctx context.Context) {
	results := make(map[AStarResultKey][]Point, AIInitialResultCap)

	for {
		select {
		case sig := <-signalCh:
			log.Printf("Received signal: %s\n", sig)
			a.Finish()
		case <-ctx.Done():
			log.Println("Finishing AStar Pathfinding")
			return
		case req := <-a.requestCh:
			res, ok := results[req.Request]
			if !ok {
				res = NewAStarInstance().Run(req.Request.Start, req.Request.Goal)
				for i := 0; i < len(res); i++ {
					results[NewAStarResultKey(res[i], req.Request.Goal)] = res[i:]
				}
			}

			req.ResponseCh <- res
			close(req.ResponseCh)
		}
	}
}

func (a *AStar) Run(start Point, goal Point) []Point {
	rch := make(chan []Point, 1)

	// Send request
	select {
	case a.requestCh <- NewAStarRequest(start, goal, rch):
	default:
		return []Point{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), AIResponseTimeout)
	defer cancel()

	// Receive response
	for {
		select {
		case r := <-rch:
			return r
		case <-ctx.Done():
			return []Point{}
		}
	}
}
