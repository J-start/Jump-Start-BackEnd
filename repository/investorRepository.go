package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


type InvestorRepository struct {
	db *sql.DB
}


func NewInvestorRepository(db *sql.DB) *InvestorRepository {
	return &InvestorRepository{db: db}
}