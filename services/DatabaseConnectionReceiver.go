package services

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetDbConnection(c *gin.Context) (*sql.DB, error) {
	db, ok := c.Get("databaseConn")

	if !ok {
		return nil, fmt.Errorf("DB connection not valid")
	}

	database, ok := db.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("DB connection not valid")
	}

	return database, nil
}
