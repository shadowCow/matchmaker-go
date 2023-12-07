package main

import (
	"github.com/shadowCow/matchmaker-go/adapters/driving/api"
	"github.com/shadowCow/matchmaker-go/domain/services"
)

func main() {
	// create the domain service
	mm := services.NewMatchmakingService()

	// create the driving adapter
	a := api.NewApiHttp(mm)

	// start the server
	a.Start(8888)
}