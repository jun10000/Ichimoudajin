package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIPawn struct {
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.AIControllerComponent
	*component.ColliderComponent[*utility.CircleF]
}

func (g ActorGeneratorStruct) NewAIPawn(location utility.Vector, rotation float64, scale utility.Vector) *AIPawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &AIPawn{}
	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.AIControllerComponent = component.NewAIControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(t, a.GetCircleBounds)
	a.UpdateBounds()
	return a
}

// NewAIPawn1 creates another version of AIPawn
func (g ActorGeneratorStruct) NewAIPawn1(location utility.Vector, rotation float64, scale utility.Vector) *AIPawn {
	a := g.NewAIPawn(location, rotation, scale)
	a.Image = utility.GetImageFromFileP("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ1.png")
	a.MaxSpeed = 150
	return a
}

// NewAIPawn2 creates another version of AIPawn
func (g ActorGeneratorStruct) NewAIPawn2(location utility.Vector, rotation float64, scale utility.Vector) *AIPawn {
	a := g.NewAIPawn(location, rotation, scale)
	a.Image = utility.GetImageFromFileP("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ2.png")
	a.MaxSpeed = 100
	return a
}

func (a *AIPawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}

func (a *AIPawn) ReceiveHit(result *utility.TraceResult[utility.Collider]) {
	if p, ok := result.HitCollider.(*Pawn); ok {
		p.AddHP(-1)
	}
}
