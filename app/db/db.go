package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	DB *mgo.Database
)

const(
	// load with environment
	MongoServerURL = "mongodb://localhost:27017"
)

func Connect(r *gin.Engine) {
	session, err := mgo.Dial(MongoServerURL)

	if err != nil {
		fmt.Printf("Can't connect to mongoDB. %v\n", err)
		panic(err.Error())
	}

	session.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", MongoServerURL)

	db := session.DB("todo")

	Session = session
	DB = db
}
