package main

import (
	"evaluate_backend/app"
)

func main() {
	app.Init()
	//路由
	r := routeInit()
	_ = r.Run()
}
