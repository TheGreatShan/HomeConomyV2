package services

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUuid() string {
	uuid := uuid.New()
	id := strings.Replace(uuid.String(), "-", "", -1)
	return id
}