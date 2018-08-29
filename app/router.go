package app

import (
	"github.com/gin-gonic/gin"

	"github.com/mavajee/todo/controllers"
)

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	r.Use(cors)

	// TODO ROUTING
	todosRouter := r.Group("/todos")
	{
		todoController := new(controllers.TodoController)

		todosRouter.GET("/", todoController.Index)
		todosRouter.POST("/", todoController.Create)
		todosRouter.GET("/:id", todoController.Show)
		todosRouter.PUT("/:id", todoController.Update)
		todosRouter.DELETE("/:id", todoController.Destroy)
	}
}