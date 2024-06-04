package main

import (
	_ "github.com/microsoft/go-mssqldb"
	"github.com/yanarowana123/kmf/internal/configs"
	"github.com/yanarowana123/kmf/internal/controller/http"
	"github.com/yanarowana123/kmf/internal/repository/mssql"
	"github.com/yanarowana123/kmf/internal/service"
	"github.com/yanarowana123/kmf/pkg/web"
	"log"
)

// @title currency
// @version		1.0
// @description	currency
func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	//connectionString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
	//	"localhost",
	//	1433,
	//	"sa",
	//	"MyPass@word",
	//	"f")
	//
	//db, err := sql.Open("sqlserver", connectionString)

	config, err := configs.New()

	if err != nil {
		return err
	}

	repo := mssql.NewRepository(config.DBConnectionString)
	taskService := service.NewCurrencyService(repo, "https://nationalbank.kz/rss/get_rates.cfm")
	router := http.NewRouter(http.NewController(taskService), config.WebServerPort)

	web.InitServer(router, *config)
	return nil
}
