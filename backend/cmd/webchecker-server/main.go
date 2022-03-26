package main

import (
	"exam.com/webchecker/internal/controller"
	"exam.com/webchecker/internal/core"
	"exam.com/webchecker/internal/gateway"
	"exam.com/webchecker/server"
)

func main() {
	// init server
	s := server.NewServer()

	// init gateway
	gateway := gateway.NewGateway()

	// init core
	core := core.NewCore(gateway)

	// init controller with this core
	controller := controller.NewController(core)

	// init route with controller
	s.InitRoute(controller)

	// run server
	s.Run()
}
