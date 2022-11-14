package repository

import "database/sql"

type CurrencyInterface interface {
	
}

type currencyImplementation struct {
	db *sql.DB
}

func NewCurrency(db *sql.DB) *currencyImplementation {
	return &currencyImplementation{
		db: db,
	}
}