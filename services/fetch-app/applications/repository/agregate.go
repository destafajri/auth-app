package repository

import "database/sql"

type AgregateInterface interface {
	
}

type agregateImplementation struct {
	db *sql.DB
}

func NewAgregate(db *sql.DB) *agregateImplementation {
	return &agregateImplementation{
		db: db,
	}
}