package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mavajee/todo/app/db"
)

type Todo struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `form:"title json:"title" binding:"required"`
	Description string `form:"description json:"description" binding:"required"`
	Completed bool `form:"completed json:"completed" binding:"required"`
}

type TodoManager struct {
	C *mgo.Collection
}

func (tm *TodoManager) All() []Todo {
	todoList := []Todo{}
	tm.C.Find(nil).All(&todoList)

	return todoList
}

func (tm *TodoManager) Add(todo Todo) Todo {
	tm.C.Insert(todo)

	return todo
}

func (tm *TodoManager) GetById(id string) Todo {
	todo := Todo{}
	tm.C.FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo
}

func (tm *TodoManager) UpdateById(id string, todo Todo) Todo {
	tm.C.UpdateId(bson.ObjectIdHex(id), todo)
	return todo
}

func (tm *TodoManager) RemoveById(id string) {
	tm.C.RemoveId(bson.ObjectIdHex(id))
}


func CreateTodoManager() *TodoManager {
	tm := new(TodoManager)
	tm.C = db.DB.C("todo")

	return tm
}
