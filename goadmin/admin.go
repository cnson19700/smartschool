package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	"github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be imported as a json file.
	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "root",
				Name:       "goadmin",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     "mysql",
			},
		},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath:   "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath:  "/var/logs/error.log",
		ColorScheme:   adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		// add generator, first parameter is the url prefix of table when visit.
		// example:
		//
		// "user" => http://localhost:9033/admin/info/user
		//
		AddGenerator("user", datamodel.GetUserTable).
		Use(r)

	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
