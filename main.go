package main

import (
	"bitroll/codefest1-users/config"
	"bitroll/codefest1-users/controller"
	"bitroll/codefest1-users/database"
	"bitroll/codefest1-users/router"
	"embed"
	"fmt"
	"log"
)

//go:embed sql/migrations/*.sql
var MigrationsFS embed.FS

const MigrationsPath = "sql/migrations"

func main() {
	log.Println("Starting...")
	config := config.ReadEnvVars()

	// setup database

	db, err := database.DBConnect(config.DBConnStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.RunMigrations(db, MigrationsFS, MigrationsPath); err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations succeded")

	// controller

	conn := controller.SetupController(db)

	// router

	router := router.SetupRouter(&conn)
	err = router.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		log.Fatal(err)
	}
}
