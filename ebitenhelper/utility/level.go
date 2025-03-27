package utility

import (
	"errors"
	"io/fs"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	pfValidCache *Smap[Point, bool]

	Name          string
	IsLooping     bool
	AIGridSize    Vector
	AIPathfinding *AStar

	Actors           []any
	Colliders        []Collider
	StaticColliders  []Collider
	MovableColliders []MovableCollider
	InputReceivers   []InputReceiver
	Players          []Player
	BeginPlayers     []BeginPlayer
	EndPlayers       []EndPlayer
	AITickers        []AITicker
	Tickers          []Ticker
	Drawers          [][]Drawer
	Namers           *Smap[string, []Actor]
	DebugDraws       []func(screen *ebiten.Image)
	Trashes          []any
}

func NewLevel(name string, isLooping bool) *Level {
	return &Level{
		pfValidCache: NewSmap[Point, bool](),

		Name:          name,
		IsLooping:     isLooping,
		AIGridSize:    NewVector(32, 32),
		AIPathfinding: NewAStar(),

		Actors:           make([]any, 0, InitialActorCap),
		Colliders:        make([]Collider, 0, InitialStaticColliderCap+InitialMovableColliderCap),
		StaticColliders:  make([]Collider, 0, InitialStaticColliderCap),
		MovableColliders: make([]MovableCollider, 0, InitialMovableColliderCap),
		InputReceivers:   make([]InputReceiver, 0, InitialInputReceiverCap),
		Players:          make([]Player, 0, InitialPlayerCap),
		BeginPlayers:     make([]BeginPlayer, 0, InitialBeginPlayerCap),
		EndPlayers:       make([]EndPlayer, 0, InitialEndPlayerCap),
		AITickers:        make([]AITicker, 0, InitialAITickerCap),
		Tickers:          make([]Ticker, 0, InitialTickerCap),
		Drawers:          make([][]Drawer, 0, ZOrderMax+1),
		Namers:           NewSmap[string, []Actor](),
		DebugDraws:       make([]func(screen *ebiten.Image), 0, DebugInitialDrawsCap),
		Trashes:          make([]any, 0, InitialTrashCap),
	}
}

func (l *Level) Add(actor any) {
	l.Actors = append(l.Actors, actor)

	if a, ok := actor.(Collider); ok {
		l.Colliders = append(l.Colliders, a)
		if m, ok := a.(MovableCollider); ok {
			l.MovableColliders = append(l.MovableColliders, m)
		} else {
			l.StaticColliders = append(l.StaticColliders, a)
		}
	}
	if a, ok := actor.(InputReceiver); ok {
		l.InputReceivers = append(l.InputReceivers, a)
	}
	if a, ok := actor.(Player); ok {
		l.Players = append(l.Players, a)
	}
	if a, ok := actor.(BeginPlayer); ok {
		l.BeginPlayers = append(l.BeginPlayers, a)
	}
	if a, ok := actor.(EndPlayer); ok {
		l.EndPlayers = append(l.EndPlayers, a)
	}
	if a, ok := actor.(AITicker); ok {
		l.AITickers = append(l.AITickers, a)
	}
	if a, ok := actor.(Ticker); ok {
		l.Tickers = append(l.Tickers, a)
	}
	if a, ok := actor.(Drawer); ok {
		z := ZOrderDefault
		if az, ok := a.(ZHolder); ok {
			z = az.ZOrder()
		}

		for range z - len(l.Drawers) + 1 {
			l.Drawers = append(l.Drawers, make([]Drawer, 0, InitialDrawerCap))
		}

		l.Drawers[z] = append(l.Drawers[z], a)
	}
	if a, ok := actor.(Actor); ok {
		n := a.GetName()
		if vs, ok := l.Namers.Load(n); ok {
			l.Namers.Store(n, append(vs, a))
		} else {
			l.Namers.Store(n, []Actor{a})
		}
	}
}

func (l *Level) Remove(actor any) {
	l.Trashes = append(l.Trashes, actor)
}

func (l *Level) EmptyTrashes() {
	for _, actor := range l.Trashes {
		l.Actors = RemoveSliceItem(l.Actors, actor)

		if a, ok := actor.(Collider); ok {
			l.Colliders = RemoveSliceItem(l.Colliders, a)
			if m, ok := a.(MovableCollider); ok {
				l.MovableColliders = RemoveSliceItem(l.MovableColliders, m)
			} else {
				l.StaticColliders = RemoveSliceItem(l.StaticColliders, a)
			}
		}
		if a, ok := actor.(InputReceiver); ok {
			l.InputReceivers = RemoveSliceItem(l.InputReceivers, a)
		}
		if a, ok := actor.(Player); ok {
			l.Players = RemoveSliceItem(l.Players, a)
		}
		if a, ok := actor.(BeginPlayer); ok {
			l.BeginPlayers = RemoveSliceItem(l.BeginPlayers, a)
		}
		if a, ok := actor.(EndPlayer); ok {
			l.EndPlayers = RemoveSliceItem(l.EndPlayers, a)
		}
		if a, ok := actor.(AITicker); ok {
			l.AITickers = RemoveSliceItem(l.AITickers, a)
		}
		if a, ok := actor.(Ticker); ok {
			l.Tickers = RemoveSliceItem(l.Tickers, a)
		}
		if a, ok := actor.(Drawer); ok {
			z := ZOrderDefault
			if az, ok := a.(ZHolder); ok {
				z = az.ZOrder()
			}

			l.Drawers[z] = RemoveSliceItem(l.Drawers[z], a)
		}
		if a, ok := actor.(Actor); ok {
			n := a.GetName()
			if vs, ok := l.Namers.Load(n); ok {
				l.Namers.Store(n, RemoveSliceItem(vs, a))
			}
		}
	}

	l.Trashes = l.Trashes[:0]
}

func GetActors[T any]() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		l := GetLevel()
		for _, a := range l.Actors {
			if ret, ok := a.(T); ok {
				if !yield(ret) {
					return
				}
			}
		}
	}
}

func GetFirstActor[T any]() (actor T, ok bool) {
	for ret := range GetActors[T]() {
		return ret, true
	}

	return *new(T), false
}

func GetActorsByName[T Actor](name string) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		l := GetLevel()
		aSlice, ok := l.Namers.Load(name)
		if !ok {
			return
		}

		for _, a := range aSlice {
			if ret, ok := a.(T); ok {
				if !yield(ret) {
					return
				}
			}
		}
	}
}

func GetFirstActorByName[T Actor](name string) (actor T, ok bool) {
	for ret := range GetActorsByName[T](name) {
		return ret, true
	}

	return *new(T), false
}

func (l *Level) AIMove(self MovableCollider, target Collider) {
	srl := self.GetRealFirstBounds().CenterLocation()
	trl := target.GetRealFirstBounds().CenterLocation()
	spl := l.RealToPFLocation(srl)
	tpl := l.RealToPFLocation(trl)

	if res, ok := l.AIPathfinding.GetResult(spl, tpl); ok {
		switch c := len(res); {
		case c > 2:
			trl1 := l.PFToRealLocation(res[1], true)
			trl2 := l.PFToRealLocation(res[2], true)
			trlave := trl1.Add(trl2.Sub(trl1).DivF(2))
			self.AddInput(trlave.Sub(srl), 1)
		case c == 2:
			trl1 := l.PFToRealLocation(res[1], true)
			self.AddInput(trl1.Sub(srl), 1)
		case c == 1:
			self.AddInput(trl.Sub(srl), 1)
		}

		DrawDebugAIPath(res)
	}
}

func (l *Level) AIIsPFLocationValid(location Point) bool {
	if r, ok := l.pfValidCache.Load(location); ok {
		return r
	}

	loc := l.PFToRealLocation(location, false)
	s := GetScreenSize()
	if loc.X < 0 || loc.Y < 0 || loc.X >= float64(s.X) || loc.Y >= float64(s.Y) {
		return false
	}

	b := NewRectangleF(
		loc.X+AIValidOffset,
		loc.Y+AIValidOffset,
		loc.X+l.AIGridSize.X-AIValidOffset,
		loc.Y+l.AIGridSize.Y-AIValidOffset)
	r, _, _ := Intersect(l.StaticColliders, b, nil)

	l.pfValidCache.Store(location, !r)
	return !r
}

func (l *Level) RealToPFLocation(realLocation Vector) Point {
	return realLocation.Div(l.AIGridSize).Trunc()
}

func (l *Level) PFToRealLocation(pfLocation Point, isCenter bool) Vector {
	r := pfLocation.ToVector().Mul(l.AIGridSize)
	if isCenter {
		r = r.Add(l.AIGridSize.DivF(2))
	}
	return r
}

func (l *Level) GetPFCacheFileName() string {
	return l.Name + ".pfd"
}

func (l *Level) LoadPFCache() error {
	return l.AIPathfinding.LoadCache(l.GetPFCacheFileName())
}

func (l *Level) LoadOrBuildPFCache() error {
	err := l.LoadPFCache()
	if errors.Is(err, fs.ErrNotExist) {
		return l.BuildPFCache()
	}

	return err
}

func (l *Level) BuildPFCache() error {
	pf := l.AIPathfinding
	sz := l.RealToPFLocation(GetScreenSize().ToVector())
	sem := make(chan Empty, runtime.GOMAXPROCS(0)-1)
	wg := sync.WaitGroup{}
	stime := time.Now()

	defer close(sem)
	log.Println("Started building PF cache")
	for sx := range sz.X {
		for sy := range sz.Y {
			for gx := range sz.X {
				for gy := range sz.Y {
					start := NewPoint(sx, sy)
					goal := NewPoint(gx, gy)
					if _, ok := pf.GetCache(start, goal); !ok {
						sem <- Empty{}
						wg.Add(1)
						go func(start Point, goal Point) {
							defer wg.Done()
							defer func() { <-sem }()
							pf.GetResultForce(start, goal)
						}(start, goal)
					}
				}
			}
		}
		log.Printf("Building PF cache: %d%%\n", (sx+1)*100/sz.X)
	}

	wg.Wait()
	log.Printf("Completed building PF cache, %.1fs elapsed.\n", time.Since(stime).Seconds())
	return pf.SaveCache(l.GetPFCacheFileName())
}

func (l *Level) AddDebugDraw(event func(*ebiten.Image)) {
	l.DebugDraws = append(l.DebugDraws, event)
}

func (l *Level) ClearDebugDraw() {
	l.DebugDraws = l.DebugDraws[:0]
}
