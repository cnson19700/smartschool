package main

import (
	"github.com/smartschool/api/routers"
	"github.com/smartschool/database"
)

func main() {
	database.Init()

	r := routers.Initialize()
	r.Run(":6969")
}
