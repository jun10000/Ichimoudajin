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
