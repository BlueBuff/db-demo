package model

import "github.com/kataras/iris"

type Route struct {
	Root string
	Rout map[string]RouteStyle
}

type HAND_TYPE string

const (
	GET  HAND_TYPE = "GET"
	POST HAND_TYPE = "POST"
)

type RouteStyle struct {
	Type   HAND_TYPE
	Handle iris.HandlerFunc
}

func NewRouteStyle(t HAND_TYPE, h iris.HandlerFunc) *RouteStyle {
	r:=new(RouteStyle)
	r.Handle=h
	r.Type=t
	return r
}

func NewRoute(root string) *Route {
	route := new(Route)
	route.Rout = make(map[string]RouteStyle)
	route.Root = root
	return route
}

func (r *Route)AppendRouteHandle(route string,style *RouteStyle) *Route{
	r.Rout[route]=*style
	return r
}

type RouteApi interface {
	MakeRoute() *Route
	//AppendRoute(route *Route)
}
