package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"time"

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

func (wr *WalletRepository) UpdateBalanceInvestor(id int, value float64, idOperation int64,repositoryService *sql.Tx) error {
	tx := repositoryService

	query := `UPDATE tb_wallet SET balance = ? WHERE idInvestor = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(value, id)
	if err != nil {
		return err
	}

	query = `UPDATE tb_operationAsset SET isProcessedAlready = 1 WHERE idAsset = ?`
	stmt, err = tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(idOperation)
	if err != nil {
		return err
	}


	return nil
}

func (wr *WalletRepository) UpdateBalanceFromOperation(id int, value float64,repositoryService *sql.Tx) error {
	tx := repositoryService

	query := `UPDATE tb_wallet SET balance = ? WHERE idInvestor = ?`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(value, id)
	if err != nil {
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

	_, err = stmt.Exec(0,id)
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


func (wr *WalletRepository) FindIdWallet(idInvestor int) (int, error) {
	var idWallet int

	query := fmt.Sprintf(`SELECT idWallet FROM tb_wallet WHERE idInvestor = %d`, idInvestor)

	err := wr.db.QueryRow(query).Scan(&idWallet)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("nenhum dado encontrado")
		} else {
			return 0, errors.New("erro ao executar a consulta")
		}
	}

	return idWallet, nil
}


func (wr *WalletRepository) SearchDatasWallet(idInvestor int) ([]entities.Asset, error) {
	query := fmt.Sprintf(`SELECT wa.assetName,wa.assetType,wa.assetQuantity FROM tb_walletAsset AS wa INNER JOIN tb_wallet AS w ON wa.idWallet = w.idWallet WHERE w.idInvestor = %d`, idInvestor)
    rows, err := wr.db.Query(query)

	if err != nil {
		return []entities.Asset{}, errors.New("erro ao buscar ativos")
	}

	assets, err := buildAssets(rows)
	if err != nil {
		return []entities.Asset{}, errors.New("erro ao buscar ativos")
	}
	return assets, nil

}

func (wr *WalletRepository) SearchBalanceInvestor(idInvestor int) (float64, error) {
	var balance float64
	query := `SELECT balance FROM tb_wallet WHERE idInvestor = ?`
	err := wr.db.QueryRow(query, idInvestor).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("saldo inexiste")
		} 
		return -1, errors.New("erro ao buscar saldo")
	}

	return balance, nil
}

func (wr *WalletRepository) FetchDatasWalletOperation(idInvestor, offset int) ([]entities.WalletOperation, error) {
	offset *= 20
	query := fmt.Sprintf(`SELECT operationType,operationValue,operationDate FROM tb_walletOperation WHERE idInvestor = %d ORDER BY operationDate DESC LIMIT 20 OFFSET %d`, idInvestor,offset)
    rows, err := wr.db.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return []entities.WalletOperation{}, errors.New("nenhum histórico encontrado")
		} 
		return []entities.WalletOperation{}, errors.New("erro ao buscar ativos")
	}

	operationsWallet, err := buildoperationWallet(rows)
	if err != nil {
		return []entities.WalletOperation{}, errors.New("erro ao buscar ativos")
	}

	return operationsWallet, nil

}


func (wr *WalletRepository) FetchDayDeposits(idInvestor int) (float64, error) {
	query := `SELECT SUM(operationValue) FROM tb_walletOperation WHERE operationdate = DATE_FORMAT(NOW(),'%Y,%m,%d') AND operationType = 'DEPOSIT' AND idInvestor = ?`
	var balanceDay float64
	err := wr.db.QueryRow(query, idInvestor).Scan(&balanceDay)

	if err != nil {

		fmt.Println(err)
		if err.Error() == `sql: Scan error on column index 0, name "SUM(operationValue)": converting NULL to float64 is unsupported` {
			return 0, nil
		}
		return -1, errors.New("erro ao buscar saldo")
	}

	return balanceDay, nil
}

func (wr *WalletRepository) CreateWallet(idInvestor int,repositoryService *sql.Tx) error{
		tx := repositoryService
		query := ` INSERT INTO tb_wallet (balance, idInvestor) VALUES (?, ?) `
		stmt, err := tx.Prepare(query)
		if err != nil {
			return errors.New("ocorreu um erro, tente novamente")
		}
		defer stmt.Close()
		_, errExec := stmt.Exec(1000, idInvestor)
		if errExec != nil {
			fmt.Println(errExec)
			return errors.New("erro ao criar carteira")
		}
		return nil
	
}


func (wr *WalletRepository) InsertOperationWallet(operationType string,operationValue float64,operationDate string,idInvestor int,repositoryService *sql.Tx) error {
	tx := repositoryService

	query := `INSERT INTO tb_walletOperation(operationType,operationValue,operationDate,idInvestor) VALUES (?, ?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(operationType,operationValue,operationDate,idInvestor)
	if err != nil {
		return err
	}


	return nil
}

func buildoperationWallet(rows *sql.Rows) ([]entities.WalletOperation, error) {
	var operationsWallet []entities.WalletOperation

	for rows.Next() {
		var operation entities.WalletOperation
		var date time.Time
		err := rows.Scan(&operation.OperationType, &operation.OperationValue, &date)
		if err != nil {
			return nil, err
		}
		operation.OperationDate = date.Format("02-01-2006")
		operationsWallet = append(operationsWallet, operation)

	}

	return operationsWallet, nil

}

func buildAssets(rows *sql.Rows) ([]entities.Asset, error) {
	var assets []entities.Asset

	for rows.Next() {
		var asset entities.Asset
		err := rows.Scan(&asset.AssetName, &asset.AssetType, &asset.AssetQuantity)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)

	}

	return assets, nil

}
