package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


type WalletRepository struct {
	db *sql.DB
}


func NewWalletRepoository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}