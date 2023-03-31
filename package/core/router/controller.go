package router

import (
	"reflect"
)

type IRouter interface {
	InitRouter()
}

type Router[T any, I any] struct {
	controller I
}

func (b *Router[T, I]) InitRouter() {
	t := reflect.TypeOf((*T)(nil)).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		//log.Printf("%s.%s", t.Name(), t.Method(i).Name)
		reflect.ValueOf(b.controller).MethodByName(t.Method(i).Name).Call([]reflect.Value{})
	}
}
func NewBaseRouter[T any, I any](controller I) *Router[T, I] {
	return &Router[T, I]{
		controller: controller,
	}
}

func InitRouter(controllers []IRouter) {
	for _, controller := range controllers {

		controller.InitRouter()
	}
}
