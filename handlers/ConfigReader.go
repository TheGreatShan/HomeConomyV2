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

	Panic(err)

	var db DbConfig

	err = json.Unmarshal(file, &db)

	Panic(err)

	return db
}
