package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type EnemySpawner struct {
	tickIndex int

	SpawnSeconds          float32
	SpawnRetryCount       int
	InvalidPlayerDistance float64
}

func (g ActorGeneratorStruct) NewEnemySpawner() *EnemySpawner {
	return &EnemySpawner{
		SpawnSeconds:          3,
		SpawnRetryCount:       10,
		InvalidPlayerDistance: 300,
	}
}

func (a *EnemySpawner) Spawn() {
	p := ActorGenerator.NewAIPawn2(utility.ZeroVector(), 0, utility.NewVector(1, 1))
	ss := utility.GetScreenSize().ToVector()
	lv := utility.GetLevel()
	if len(lv.Players) == 0 {
		return
	}

	pll := lv.Players[0].GetLocation()
	for range a.SpawnRetryCount {
		l := utility.RandomVectorPtr().Mul(ss)
		p.SetLocation(l)

		// Constraints
		if ok, _, _ := utility.Intersect(lv.Colliders, p.GetRealFirstBounds(), nil); ok {
			continue
		}
		if pll.Sub(l).Length2() < a.InvalidPlayerDistance*a.InvalidPlayerDistance {
			continue
		}

		// When pass
		lv.Add(p)
		return
	}
}

func (a *EnemySpawner) Tick() {
	a.tickIndex++
	if a.tickIndex%int(utility.TickCount*a.SpawnSeconds) == 0 {
		a.Spawn()
	}
}
