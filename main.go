package main

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/service"
)

func main() {
	database.Init()
	var dummy dto.DeviceSignal
	service.CheckIn(dummy)
	// r := routers.Initialize()
	// r.Run(":6969")
}
