package controllers

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2/bson"


	"github.com/mavajee/todo/models"
)

type TodoController struct{}

type IndexSchema struct {
	// https://github.com/gin-gonic/gin#model-binding-and-validation
	Completed bool
}

func (t TodoController) Index(c *gin.Context) {
	tm := models.CreateTodoManager()
	// a.CompleteOnly()

	c.JSON(200, tm.All())
}

func (t TodoController) Create(c *gin.Context) {
	tm := models.CreateTodoManager()
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, tm.Add(todo))
}

func (t TodoController) Show(c *gin.Context) {
	tm := models.CreateTodoManager()
	id := c.Params.ByName("id")

	c.JSON(200, tm.GetById(id))
}

func (t TodoController) Update(c *gin.Context) {
	tm := models.CreateTodoManager()
	id := c.Params.ByName("id")

	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tm.UpdateById(id, todo))
}

func (t TodoController) Destroy(c *gin.Context) {
	tm := models.CreateTodoManager()
	id := c.Params.ByName("id")

	tm.RemoveById(id)

	c.JSON(200, gin.H{})
}
