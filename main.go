package  main

import (
	"bee-keeper/module"
	"bee-keeper/pages"
	"bee-keeper/tables"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                         // ui theme
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	module.InitDatabaseConn(eng.PostgresqlConnection())

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	_ = r.Run(":9000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.PostgresqlConnection().Close()
}

