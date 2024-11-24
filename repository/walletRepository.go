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
			return 0, errors.New("nenhum dado encontrado")
		} else {
			return 0, errors.New("erro ao executar a consulta")
		}
	}

	return balance, nil
}

func (wr *WalletRepository) UpdateBalanceInvestor(id int, value float64, idOperation int64) error {
	tx, err := wr.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE tb_wallet SET balance = ? WHERE idInvestor = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(value, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `UPDATE tb_operationAsset SET isProcessedAlready = 1 WHERE idAsset = ?`
	stmt, err = tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		errDelete := wr.deleteOpreation(int(idOperation))
		if errDelete != nil {
			return errDelete
		}
		return err
	}
	_, err = stmt.Exec(idOperation)
	if err != nil {
		tx.Rollback()
		errDelete := wr.deleteOpreation(int(idOperation))
		if errDelete != nil {
			return errDelete
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (wr *WalletRepository) IsBalanceExists(id int) (float64, error) {
	var balance float64

	query := fmt.Sprintf(`SELECT balance FROM tb_wallet WHERE idInvestor = %d`, id)

	err := wr.db.QueryRow(query).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("nenhum dado encontrado")
		} else {
			return 0, errors.New("erro ao executar a consulta")
		}
	}

	return balance, nil
}

func (wr *WalletRepository) CreateBalanceUser(id int) error {
	tx, err := wr.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO tb_wallet(balance, idInvestor) VALUES (?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(0, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (wr *WalletRepository) deleteOpreation(idOperation int) error {
	query := fmt.Sprintf(`DELETE FROM tb_operationAsset WHERE idAsset = %d `, idOperation)
	_, err := wr.db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}
