package main

import (
	route "backend-starter-project/delivery/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// app := bootstrap.App()

	// env := app.Env

	// db := app.Mongo.Database(env.DBName)
	// defer app.CloseDBConnection()

	// Do something with the database
	// _ = db
	route.Setup(gin.Default())
}
