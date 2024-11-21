package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


type WalletRepository struct {
	db *sql.DB
}


func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (wr *WalletRepository) FindBalanceInvestor(id int) (float64, error) {
	var balance float64
	
	query := fmt.Sprintf(`SELECT balance FROM tb_wallet WHERE idInvestor = %d`, id)
	
	err := wr.db.QueryRow(query).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0,errors.New("nenhum dado encontrado")
		} else {
			return 0,errors.New("erro ao executar a consulta")
		}
	}

	return balance, nil
}