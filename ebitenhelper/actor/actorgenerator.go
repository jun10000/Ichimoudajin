package actor

import (
	"fmt"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

var ActorGenerator ActorGeneratorStruct

func init() {
	ActorGenerator = NewActorGeneratorStruct()
}

type ActorGeneratorStruct struct {
	refValue reflect.Value
}

func NewActorGeneratorStruct() ActorGeneratorStruct {
	g := ActorGeneratorStruct{}
	g.refValue = reflect.ValueOf(g)
	return g
}

func (g ActorGeneratorStruct) NewActorByName(name string, location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector) (any, error) {
	m := g.refValue.MethodByName("New" + name)
	if !m.IsValid() {
		return nil, fmt.Errorf("method 'New%s' is not found", name)
	}

	argc := m.Type().NumIn()
	if argc > 4 {
		return nil, fmt.Errorf("method New%s has invalid argument counts: %d", name, argc)
	}

	argv := []reflect.Value{
		reflect.ValueOf(location),
		reflect.ValueOf(rotation),
		reflect.ValueOf(scale),
		reflect.ValueOf(size),
	}
	ret := m.Call(argv[:argc])
	if len(ret) == 0 {
		return nil, fmt.Errorf("method New%s does not return value", name)
	}

	return ret[0].Interface(), nil
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

func (g ActorGeneratorStruct) NewAnimatedActor(location utility.Vector, rotation float64, scale utility.Vector) *AnimatedActor {
	a := &AnimatedActor{}
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	return a
}

func (g ActorGeneratorStruct) NewBlockingArea(location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector) *BlockingArea {
	t := utility.NewStaticTransform(location, rotation, scale)

	a := &BlockingArea{}
	a.StaticColliderComponent = component.NewStaticColliderComponent(t, a.GetRectangleBounds)
	a.size = size

	a.UpdateBounds()
	return a
}

func (g ActorGeneratorStruct) NewDestroyer() *Destroyer {
	return &Destroyer{
		status: DestroyerStatusDisable,
		circle: utility.NewCircleF(0, 0, 0),

		GrowSpeed:   1,
		ShrinkSpeed: 2,
		MaxRadius:   120,
		BorderWidth: 2,
		BorderColor: utility.ColorLightBlue.ToRGBA(0xff),
		FillColor:   utility.ColorLightBlue.ToRGBA(0x20),
	}
}

func (g ActorGeneratorStruct) NewEnemySpawner() *EnemySpawner {
	return &EnemySpawner{
		SpawnSeconds:          3,
		SpawnRetryCount:       10,
		InvalidPlayerDistance: 300,
	}
}

func (g ActorGeneratorStruct) NewImageActor(location utility.Vector, rotation float64, scale utility.Vector) *ImageActor {
	a := &ImageActor{}
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawImageComponent = component.NewDrawImageComponent(a)
	return a
}

func (g ActorGeneratorStruct) NewPawn(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &Pawn{}
	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(t, a.GetCircleBounds)

	a.destroyer = g.NewDestroyer()

	a.UpdateBounds()
	return a
}

// NewPawn1 creates another version of Pawn
func (g ActorGeneratorStruct) NewPawn1(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	a := g.NewPawn(location, rotation, scale)
	a.Image = utility.GetImageFromFileP("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	a.MaxSpeed = 200
	return a
}
