package repository

import "database/sql"

type FetchInterface interface {
	
}

type fetchImplementation struct {
	db *sql.DB
}

func NewFetch(db *sql.DB) *fetchImplementation {
	return &fetchImplementation{
		db: db,
	}
}