package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type EnemySpawner struct {
	tickIndex int

	SpawnSeconds float32
}

func NewEnemySpawner() *EnemySpawner {
	return &EnemySpawner{
		SpawnSeconds: 3,
	}
}

func (a *EnemySpawner) Spawn() {
	img, err := utility.GetImageFromFile("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ2.png")
	utility.PanicIfError(err)

	l := utility.RandomVectorPtr().Mul(utility.GetScreenSize().ToVector())
	p := NewAIPawn(l, 0, utility.NewVector(1, 1))
	p.Image = img
	p.MaxSpeed = 200
	utility.GetLevel().Add(p)
}

func (a *EnemySpawner) Tick() {
	a.tickIndex++
	if a.tickIndex%int(utility.TickCount*a.SpawnSeconds) == 0 {
		a.Spawn()
	}
}
