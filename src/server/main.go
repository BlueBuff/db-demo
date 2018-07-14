package main

import (
	_ "hdg.com/db-demo/src/server/common"
	"hdg.com/db-demo/src/server/route"
)

func main() {
	route.InitRoute()
}
