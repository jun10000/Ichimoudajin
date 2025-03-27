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

type GenerateActorOptions struct {
	Name      string
	Location  utility.Vector
	Rotation  float64
	Scale     utility.Vector
	Size      utility.Vector
	IsVisible bool
	ExtraText *ExtraTextInfo
}

type ExtraTextInfo struct {
	Size  float64
	Text  string
	Color utility.RGB
}

func NewGenerateActorOptions() *GenerateActorOptions {
	return &GenerateActorOptions{
		Scale:     utility.DefaultScale(),
		IsVisible: true,
	}
}

type ActorGeneratorStruct struct {
	refValue reflect.Value
}

func NewActorGeneratorStruct() ActorGeneratorStruct {
	g := ActorGeneratorStruct{}
	g.refValue = reflect.ValueOf(g)
	return g
}

func (g ActorGeneratorStruct) NewActorByTypeName(name string, options *GenerateActorOptions) (utility.Actor, error) {
	m := g.refValue.MethodByName("New" + name)
	if !m.IsValid() {
		return nil, fmt.Errorf("method 'New%s' is not found", name)
	}

	argv := []reflect.Value{
		reflect.ValueOf(options.Name),
		reflect.ValueOf(options.Location),
		reflect.ValueOf(options.Rotation),
		reflect.ValueOf(options.Scale),
		reflect.ValueOf(options.Size),
		reflect.ValueOf(options.IsVisible),
		reflect.ValueOf(options.ExtraText),
	}

	argc := m.Type().NumIn()
	if argc > len(argv) {
		return nil, fmt.Errorf("method New%s has invalid argument counts: %d", name, argc)
	}

	ret := m.Call(argv[:argc])
	if len(ret) == 0 {
		return nil, fmt.Errorf("method New%s does not return value", name)
	}

	return ret[0].Interface().(utility.Actor), nil
}
