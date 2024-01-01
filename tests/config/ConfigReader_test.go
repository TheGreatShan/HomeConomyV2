package main

import (
	"HomeConomyv2GO/handlers"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	t.Run("Valid JSON", func(t *testing.T) {
		temp, err := os.CreateTemp("", "dbconfig.json")
		if err != nil {
			t.Fatalf("Cannot create temporary file: %v", err)
		}

		defer os.Remove(temp.Name())

		dbConfig := handlers.DbConfig{
			Host:     "value",
			Port:     "123",
			Database: "value",
			Username: "value",
			Password: "value",
		}

		marshal, err := json.Marshal(dbConfig)
		if err != nil {
			t.Fatalf("Cannot marshal DbConfig: %v", err)
		}

		os.WriteFile(temp.Name(), marshal, 0644)

		result := handlers.GetDbConfig(temp.Name())

		if !reflect.DeepEqual(dbConfig, result) {
			t.Errorf("Expected %+v, got %+v", dbConfig, result)
		}
	})
}
