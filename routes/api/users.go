package api

import (
	"go-todo-practice1/models/db"
	"go-todo-practice1/models/receiver"
	"go-todo-practice1/routes/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter to setup routing on user management
func UserRouter(c *gin.Engine) {
	v1 := c.Group("api/v1/user")
	v1.GET("/login", loginUser)
	v1.POST("/register", registerUser)
}

func loginUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GET api/v1/user/login"})

}

func registerUser(c *gin.Context) {
	// Get request body for email password name
	var register receivers.RegisterForm
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := helper.HashPassword(register.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := db.User{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
	}

	data := db.DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "POST api/v1/user/register",
		"User":    data,
	})
}
