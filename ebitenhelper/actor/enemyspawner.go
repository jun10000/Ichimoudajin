package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type EnemySpawner struct {
	*component.ActorCom
	SpawnSeconds          float32
	SpawnRetryCount       int
	InvalidPlayerDistance float64
}

func (g ActorGeneratorStruct) NewEnemySpawner(options *NewActorOptions) *EnemySpawner {
	return &EnemySpawner{
		ActorCom:              component.NewActorCom(options.Name),
		SpawnSeconds:          3,
		SpawnRetryCount:       10,
		InvalidPlayerDistance: 300,
	}
}

func (a *EnemySpawner) Spawn() {
	po := NewNewActorOptions()
	po.Name = "SpawnedEnemy"
	p := ActorGenerator.NewAIPawn2(po)

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
	if utility.GetTickIndex()%int(utility.TickCount*a.SpawnSeconds) == 0 {
		a.Spawn()
	}
}
