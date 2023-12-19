package handlers

import (
	"encoding/json"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func GetDbConfig(path string) DbConfig {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var db DbConfig

	err = json.Unmarshal(file, &db)

	if err != nil {
		panic(err)
	}

	return db
}
