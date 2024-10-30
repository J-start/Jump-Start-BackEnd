package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ShareRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *ShareRepository {
	return &ShareRepository{db: db}
}

func (repo *ShareRepository) FindAllShares() {
	//TODO make a query to find all shares
}