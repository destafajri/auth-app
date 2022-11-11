package repository

import (
	"database/sql"
	"fmt"

	"github.com/destafajri/auth-app/applications/entity"
)

type userImplementation struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *userImplementation{
	return &userImplementation{
		db: db,
	}
}

func(user userImplementation) Register(users *entity.UserEntity) error{
	query :=`INSERT INTO users(
			id,
			name, 
			phone,
			role,
			password
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`
	err := user.db.QueryRow(query, users.ID, users.Name, users.Phone, users.Role , users.Password).Scan(&users.ID)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		fmt.Println("error eksekusi query")
		fmt.Println(err)
		return err
	}

	return nil
}