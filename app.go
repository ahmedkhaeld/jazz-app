package main

import (
	"github.com/ahmedkhaeld/jazz"
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
)

type application struct {
	*jazz.Jazz
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func run() *application {
	//get the root path of the application
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//define a jazz instance
	j := &jazz.Jazz{}

	//init the jazz
	err = j.New(rootPath)
	if err != nil {
		j.ErrorLog.Fatal(err)
	}

	m := &middleware.Middleware{
		Jazz: j,
	}

	h := &handlers.Handlers{
		Jazz: j,
	}
	app := &application{
		Jazz:       j,
		Handlers:   h,
		Middleware: m,
	}
	app.Models = data.New(app.Jazz.DB.Pool)
	app.Middleware.Models = app.Models
	//add new routes with default ones
	app.Jazz.Routes = app.routes()
	return app
}
