package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIPawn struct {
	*component.MovementCom
	*component.DrawAnimationCom
	*component.AIControllerCom
	*component.ColliderCom[*utility.CircleF]
}

func (g ActorGeneratorStruct) NewAIPawn(location utility.Vector, rotation float64, scale utility.Vector) *AIPawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &AIPawn{}
	a.MovementCom = component.NewMovementCom(a)
	a.DrawAnimationCom = component.NewDrawAnimationCom(a)
	a.AIControllerCom = component.NewAIControllerCom(a)
	a.ColliderCom = component.NewColliderCom(t, a.GetCircleBounds)
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
	a.MovementCom.Tick()
	a.DrawAnimationCom.Tick()
}

func (a *AIPawn) ReceiveHit(result *utility.TraceResult[utility.Collider]) {
	if p, ok := result.HitCollider.(*Pawn); ok {
		p.AddHP(-1)
	}
}
