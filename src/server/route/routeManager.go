package route

import (
	"hdg.com/db-demo/src/server/controller"
)

type RouteRegister interface {
	Register()
}

type RouteManager struct {
	RouteRegister
}

func (*RouteManager) Register(){
	userController:=controller.NewUserController()
	AppendRoute(userController.MakeRoute())
}

func AddRoute(){
	register:=new(RouteManager)
	register.Register()
}