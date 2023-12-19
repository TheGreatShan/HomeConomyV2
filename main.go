package main

import (
	//"HomeConomyv2GO/handlers"
	"HomeConomyv2GO/version"
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
)

type Database struct {
	Db *sql.DB
}

func main()  {
	// config := handlers.GetDbConfig("config.json")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	// db, err := sql.Open("mysql", dsn)

	// if err != nil {
	//	panic(err)
	//}

	//defer func() {
	//	if err := db.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	// database := Database{Db: db}
	
	router := gin.Default()
	router.GET("/version", version.GetVersion)
	
	router.Run("localhost:8080")
}