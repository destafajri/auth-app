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

func(user userImplementation) Remindme(phone string) (*entity.UserEntity, error){
	var users entity.UserEntity
	query := `SELECT id, name, phone, role, password FROM users where phone=$1`
	
	err := user.db.QueryRow(query, phone).Scan(
		&users.ID,
		&users.Name,
		&users.Phone,
		&users.Role,
		&users.Password,
	)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		fmt.Println("error eksekusi query")
		fmt.Println(err)
		return nil, err
	}

	return &users, nil
}

func(user userImplementation) Login(phone string) (*entity.UserEntity, error){
	var users entity.UserEntity

	query := `SELECT
				id,
				name,
				phone,
				role,
				password
			FROM users
			WHERE
				phone=$1
			`
	err := user.db.QueryRow(query, phone).Scan(
		&users.ID,
		&users.Name,
		&users.Phone,
		&users.Role,
		&users.Password,
	)

	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		fmt.Println("error eksekusi query")
		fmt.Println(err)
		return nil, err
	}

	return &users, nil
}