package handlers

import (
	"HomeConomyv2GO/services"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func CheckIfExistsById(c *gin.Context, table string) bool {
	connection, err := services.GetDbConnection(c)

	Panic(err)

	id, err := hex.DecodeString(c.Param("id"))
	Panic(err)
	item, err := connection.Query("SELECT * FROM "+table+" WHERE id = ?", id)

	Panic(err)

	if item.Next() {
		return true
	} else {
		return false
	}
}

func CheckIfExistsByName(c *gin.Context, table string, name string) bool {
	connection, err := services.GetDbConnection(c)

	Panic(err)

	item, err := connection.Query("SELECT * FROM "+table+" WHERE name = ?", name)

	Panic(err)

	if item.Next() {
		return true
	} else {
		return false
	}
}
