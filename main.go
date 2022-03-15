package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/service/excel"
	"github.com/smartschool/tables"
	"os"
	"os/signal"
)

func main() {
	database.Init()

	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be imported as a json file.
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				File:       "./admin.db",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverSqlite,
			},
		},
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Debug:     true,
		Language:  language.EN,
		Theme:     "adminlte",
	}

	// Add configuration and plugins, use the Use method to mount to the web framework.
	_ = eng.AddConfig(&cfg).
		AddGenerators(tables.Generators).
		Use(r)

	r.Static("/public", "./public")

	eng.HTML("GET", "/admin", DashboardPage)

	r.GET("/summary", excel.ExportSummary)
	r.POST("/course", excel.ImportCourse)

	go func() {
		r.Run(":9035")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func DashboardPage(ctx *context.Context) (types.Panel, error) {
	return types.Panel{
		Content:     "hello world",
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}
