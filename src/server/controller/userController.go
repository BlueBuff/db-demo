package controller

import (
	"hdg.com/db-demo/src/server/service"
	"github.com/kataras/iris"
	"hdg.com/db-demo/src/server/model"
)

type UserController struct {
	service service.UserService
}

func (con *UserController) MakeRoute() *model.Route{
	routes:= model.NewRoute("/user").AppendRouteHandle("/test",model.NewRouteStyle(model.GET,con.GetUserById))
	routes.AppendRouteHandle("/cache",model.NewRouteStyle(model.GET,con.GetCacheUserById))
	return routes
}

func NewUserController() *UserController{
	userController:=new(UserController)
	userController.service=service.NewUserServiceImpl()
	return userController
}

func (con *UserController) GetUserById(ctx *iris.Context) {
	id,err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(503, "参数错误")
		return
	}
	user := con.service.GetUser(id)
	if user == nil {
		ctx.JSON(403, "not found")
		return
	}
	ctx.JSON(200, user)
}

func (con *UserController) GetCacheUserById(ctx *iris.Context){
	id,err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(503, "参数错误")
		return
	}
	if id == 0 {
		ctx.JSON(502, "ID无效")
		return
	}
	user:=con.service.GetCacheUser(id)
	if user == nil {
		ctx.JSON(403, "not found")
		return
	}
	ctx.JSON(200, user)
}
