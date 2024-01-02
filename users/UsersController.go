package users

import (
	"HomeConomyv2GO/handlers"
	"HomeConomyv2GO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user User

	connection, err := services.GetDbConnection(c)
	handlers.Panic(err)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	exists := handlers.CheckIfExistsByName(c, "users", user.Name)
	if exists == true {
		c.IndentedJSON(http.StatusConflict, gin.H{"Message": "user already exists"})
		return
	}

	newUser, err := CreateNewUser(connection, user)
	handlers.Panic(err)

	c.IndentedJSON(http.StatusCreated, newUser)
}
