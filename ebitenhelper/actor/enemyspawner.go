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

func (a *EnemySpawner) Spawn() {
	img, err := utility.GetImageFromFile("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ2.png")
	utility.PanicIfError(err)

	p := ActorGenerator.NewAIPawn(utility.ZeroVector(), 0, utility.NewVector(1, 1))
	p.Image = img
	p.MaxSpeed = 150

	ss := utility.GetScreenSize().ToVector()
	lv := utility.GetLevel()
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
