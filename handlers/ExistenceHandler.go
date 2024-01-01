package handlers

import (
	"HomeConomyv2GO/services"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func CheckIfExistsById(c *gin.Context, table string) bool {
	connection, err := services.GetDbConnection(c)

	if err != nil {
		panic(err)
	}
	id, err := hex.DecodeString(c.Param("id"))
	if err != nil {
		panic(err)
	}
	item, err := connection.Query("SELECT * FROM "+table+" WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	if item.Next() {
		return true
	} else {
		return false
	}
}

func CheckIfExistsByName(c *gin.Context, table string, name string) bool {
	connection, err := services.GetDbConnection(c)

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	item, err := connection.Query("SELECT * FROM "+table+" WHERE name = ?", name)

	if err != nil {
		panic(err)
	}

	if item.Next() {
		return true
	} else {
		return false
	}
}
