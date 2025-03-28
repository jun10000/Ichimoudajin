package actor

import (
	"fmt"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

var ActorGenerator ActorGeneratorStruct

func init() {
	ActorGenerator = NewActorGeneratorStruct()
}

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

type ActorGeneratorStruct struct {
	selfRef reflect.Value
}

func NewActorGeneratorStruct() ActorGeneratorStruct {
	g := ActorGeneratorStruct{}
	g.selfRef = reflect.ValueOf(g)
	return g
}

func (g ActorGeneratorStruct) NewActorByTypeName(name string, options *NewActorOptions) (utility.Actor, error) {
	ret, err := utility.CallMethodByName(g.selfRef, "New"+name, options)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf("method New%s does not return value", name)
	}

	return ret[0].(utility.Actor), nil
}
