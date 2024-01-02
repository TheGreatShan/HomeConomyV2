package users

import (
	"HomeConomyv2GO/services"
	"database/sql"
	"encoding/hex"
)

func CreateNewUser(connection *sql.DB, user User) (User, error) {
	id := services.CreateUuid()


	decodedId, err := hex.DecodeString(services.CreateUuid())
	if err != nil {
		return user, err
	}

	user.Id = id

	query, err := connection.Query("INSERT INTO users (id, name) VALUES (?, ?)", decodedId, user.Name)

	if err != nil {
		return user, err
	}

	defer query.Close()

	return user, nil
}
