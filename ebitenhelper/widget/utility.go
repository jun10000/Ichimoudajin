package widget

import (
	"log"

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

func GetWidgetObjectByNameP[T WidgetObjecter](actorName string, objectName string) T {
	o, ok := GetWidgetObjectByName[T](actorName, objectName)
	if !ok {
		log.Panicf("widget object '%s' is not found", objectName)
	}

	return o
}
