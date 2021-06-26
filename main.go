package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/gorm"

	// internal packages
	"main/routes"
	"main/binance"
)

// create echo framework
var e *echo.Echo

// db Object
var db *gorm.DB

func main() {

	fmt.Println("Server starting...")

	////////////////////////
	// Setup
	////////////////////////

	// .env file
	godotenv.Load(".env")
	enviro := os.Getenv("enviro")
	
	// initialise the binance API
	binance.Init()

	// connect to mysql
	initDb()

	// start echo framework
	e = echo.New()

	// load templates and connect them to echo
	setupTemplater(e)

	// refresh templates on each request while developing?
	if enviro == "local" {
		e.Use(middlewareReprocessTemplates)
	}

	// serve static assets to requests to /assets/*
	e.Static("/assets", "assets")

	////////////////////////
	// routes
	////////////////////////

	routes.SystemRoutes(db, e)
	routes.MainRoutes(db, e)
	routes.AuthRoutes(db, e)
	routes.InvestorRoutes(db, e)
	

	////////////////////////
	// Start Server
	////////////////////////
	server := e.Start("localhost:3000")
	e.Logger.Fatal(server)

	defer e.Close()

}
