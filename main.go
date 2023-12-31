package main

import (
	"HomeConomyv2GO/company"
	"HomeConomyv2GO/handlers"
	"HomeConomyv2GO/version"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Db *sql.DB
}

func ApiMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func main() {
	config := handlers.GetDbConfig("config.json")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	database := Database{Db: db}

	router := gin.Default()

	router.Use(ApiMiddleware(database.Db))
	router.GET("/version", version.GetVersion)

	router.GET("/companies", company.GetCompanies)
	router.GET("/companies/:id", company.GetCompany)
	router.POST("/companies", company.CreateCompany)

	router.Run("localhost:8080")
}
