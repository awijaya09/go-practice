package main

import (
	"go-todo-practice1/models/db"
	"go-todo-practice1/routes/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.Connect()

	api.UserRouter(router)
	db.Connect()

	router.Run(":5000")
}
