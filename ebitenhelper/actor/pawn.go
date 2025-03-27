package actor

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type PawnStates int

const (
	PawnStateNormal PawnStates = iota
	PawnStateInvincible
	PawnStateYararechatta
)

type Pawn struct {
	*component.MovementCom
	*component.DrawAnimationCom
	*component.ControllerCom
	*component.ColliderCom[*utility.CircleF]

	state           PawnStates
	invincibleTimer *utility.CallTimer
	currentHP       int
	destroyer       *Destroyer
	hpWidget        *TextWidget
	gameoverWidget  *TextWidget

	MaxHP                 int
	InvincibleSeconds     float32
	InvincibleDrawSeconds float32
}

func (g ActorGeneratorStruct) NewPawn(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &Pawn{}
	a.MovementCom = component.NewMovementCom(a)
	a.DrawAnimationCom = component.NewDrawAnimationCom(a)
	a.ControllerCom = component.NewControllerCom(a)
	a.ColliderCom = component.NewColliderCom(t, a.GetCircleBounds)

	a.state = PawnStateNormal
	a.invincibleTimer = utility.NewCallTimer()
	a.currentHP = 3

	a.MaxHP = 3
	a.InvincibleSeconds = 3
	a.InvincibleDrawSeconds = 0.06

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

func (a *Pawn) BeginPlay() {
	if d, ok := utility.GetFirstActor[*Destroyer](); ok {
		a.destroyer = d
	} else {
		log.Panicln("actor 'Destroyer' is not found")
	}

	if w, ok := utility.GetFirstActorByName[*TextWidget]("HPWidget"); ok {
		a.hpWidget = w
	} else {
		log.Panicln("actor 'HPWidget' is not found")
	}

	if w, ok := utility.GetFirstActorByName[*TextWidget]("GameOverWidget"); ok {
		a.gameoverWidget = w
	} else {
		log.Panicln("actor 'GameOverWidget' is not found")
	}
}

func (a *Pawn) ReceiveMouseButtonInput(button ebiten.MouseButton, state utility.PressState, pos utility.Point) {
	a.ControllerCom.ReceiveMouseButtonInput(button, state, pos)
	if button != ebiten.MouseButtonLeft {
		return
	}

	switch state {
	case utility.PressStatePressed:
		a.destroyer.Start(pos.ToVector())
	case utility.PressStateReleased:
		a.destroyer.Finish()
	}
}

func (a *Pawn) Tick() {
	a.invincibleTimer.Tick()
	a.MovementCom.Tick()
	a.ApplyHPToWidget()
}

func (a *Pawn) Draw(screen *ebiten.Image) {
	switch a.state {
	case PawnStateNormal:
		a.DrawAnimationCom.Draw(screen)
	case PawnStateInvincible:
		ni := int(float32(utility.GetTickIndex()) / (float32(utility.TickCount) * a.InvincibleDrawSeconds))
		if ni%2 == 0 {
			a.DrawAnimationCom.Draw(screen)
		}
	}
}

func (a *Pawn) ReceiveHit(result *utility.TraceResult[utility.Collider]) {
	if _, ok := result.HitCollider.(*AIPawn); ok {
		a.AddHP(-1)
	}
}

func (a *Pawn) ApplyHPToWidget() {
	a.hpWidget.Text = fmt.Sprintf("HP %d", a.currentHP)
}

func (a *Pawn) AddHP(delta int) {
	switch a.state {
	case PawnStateNormal:
		a.currentHP += delta
		if a.currentHP > a.MaxHP {
			a.currentHP = a.MaxHP
		}

		if a.currentHP <= 0 {
			a.state = PawnStateYararechatta
			a.ReceiveDeath()
		} else if delta < 0 {
			a.state = PawnStateInvincible
			a.invincibleTimer.StartCallTimer(func() {
				a.state = PawnStateNormal
			}, a.InvincibleSeconds)
		}
	case PawnStateInvincible:
		if delta > 0 {
			a.currentHP += delta
			if a.currentHP > a.MaxHP {
				a.currentHP = a.MaxHP
			}
		}
	}
}

func (a *Pawn) ReceiveDeath() {
	utility.GetLevel().Remove(a)
	a.hpWidget.IsVisible = false
	a.gameoverWidget.IsVisible = true
}
