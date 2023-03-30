package controller

import (
	"log"
	"reflect"
)

type IController interface {
	InitRouter()
}

type Controller[T any, I any] struct {
	controller I
}

func (b *Controller[T, I]) InitRouter() {
	t := reflect.TypeOf((*T)(nil)).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		log.Printf("%s.%s", t.Name(), t.Method(i).Name)
		reflect.ValueOf(b.controller).MethodByName(t.Method(i).Name).Call([]reflect.Value{})
	}
}
func NewBaseController[T any, I any](controller I) *Controller[T, I] {
	return &Controller[T, I]{
		controller: controller,
	}
}

func InitControllers(controllers []IController) {
	for _, controller := range controllers {

		controller.InitRouter()
	}
}
