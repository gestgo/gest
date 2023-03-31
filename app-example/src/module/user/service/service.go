package service

import "go.uber.org/fx"

type IUserService interface {
	Log()
}
type userService struct {
}

func (u *userService) Log() {
	//TODO implement me
	panic("implement me")
}

type UserServiceParams struct {
	fx.In
}

func NewUserService(params UserServiceParams) IUserService {
	return &userService{}

}
