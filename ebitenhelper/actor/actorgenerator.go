package actor

import (
	"fmt"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type NewActorOptions struct {
	Name      string
	Location  utility.Vector
	Rotation  float64
	Scale     utility.Vector
	Size      utility.Vector
	IsVisible bool
	Text      *NewActorTextOptions
}

type NewActorTextOptions struct {
	Size  float64
	Text  string
	Color utility.RGB
}

func NewNewActorOptions() *NewActorOptions {
	return &NewActorOptions{
		Scale:     utility.DefaultScale(),
		IsVisible: true,
	}
}

var ActorGenerator = ActorGeneratorStruct{}

type ActorGeneratorStruct struct{}

func (g ActorGeneratorStruct) NewActorByTypeName(name string, options *NewActorOptions) (utility.Actor, error) {
	rets, err := utility.CallMethodByName(g, "New"+name, options)
	if err != nil {
		return nil, err
	}

	if len(rets) == 0 {
		return nil, fmt.Errorf("method 'New%s' does not return value", name)
	}

	ret, ok := rets[0].(utility.Actor)
	if !ok {
		return nil, fmt.Errorf("method 'New%s' does not return Actor", name)
	}

	return ret, nil
}
