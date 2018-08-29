package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mavajee/todo/app"
	"github.com/mavajee/todo/app/db"
)

func main() {
	r := gin.Default()

	app.MakeRoutes(r)
	db.Connect(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
