package main

import (
	"os"
	"os/signal"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // sql driver
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/service/excel"
	"github.com/smartschool/tables"
)

func main() {
	locale, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = locale

	database.Init()

	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be imported as a json file.
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "13.228.244.196",
				Port:       "5432",
				User:       "busmapdb",
				Pwd:        "frjsdfhaflpzlcdzgnfvuxkdwiiiiklpojzowxajmendeeoqtbzyrgi",
				Name:       "phenikaamaas_attendancedb",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverPostgresql,
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

	r.GET("/admin", func(context *gin.Context) {
		context.Redirect(301, "/admin/info/courses")
	})

	r.GET("/summary", excel.ExportSummary)
	r.POST("/course", excel.ImportCourse)
	r.POST("/user", excel.ImportUser)

	go func() {
		r.Run(":6001")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func DashboardPage(ctx *context.Context) (types.Panel, error) {
	return types.Panel{
		Content:     "",
		Title:       "Dashboard",
		Description: "",
	}, nil
}
