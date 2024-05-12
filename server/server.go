package server

import (
	"base-golang-rest-api-docker-container/db"
	"base-golang-rest-api-docker-container/routes"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbConn, err := db.InitDBConnection()
	if err != nil {
		fmt.Printf("DB connection error: %s", err)
		log.Fatal()
	}

	if err := db.InitTables(dbConn); err != nil {
		fmt.Printf("DB intialise tables error: %s", err)
		log.Fatal()
	}

	routes.InitTodoRouter(e, dbConn)

	// Start server
	e.Logger.Fatal(e.Start(":4200"))
}
