package main

import "github.com/smartschool/api/routers"

func main() {
	//database.Init()

	r := routers.Initialize()
	r.Run(":6969")
}
