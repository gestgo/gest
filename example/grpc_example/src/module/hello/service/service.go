package service

import "go.uber.org/fx"

type IHelloService interface {
}

type HelloService struct {
}

type HelloServiceParams struct {
	fx.In
}

func NewHelloService(params HelloServiceParams) IHelloService {
	return &HelloService{}

}
