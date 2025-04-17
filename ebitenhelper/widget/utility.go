package widget

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func GetWidgetObjectByName[T WidgetObjecter](actorName string, objectName string) (object T, ok bool) {
	a, ok := utility.GetFirstActorByName[*Widget](actorName)
	if !ok {
		return *new(T), false
	}

	o := a.GetWidgetObject(objectName)
	if o == nil {
		return *new(T), false
	}

	wo, ok := o.(T)
	if !ok {
		return *new(T), false
	}

	return wo, true
}
