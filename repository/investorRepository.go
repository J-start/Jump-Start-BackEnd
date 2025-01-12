package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"

	"github.com/go-sql-driver/mysql"

)

type InvestorRepository struct {
	db *sql.DB
}

func NewInvestorRepository(db *sql.DB) *InvestorRepository {
	return &InvestorRepository{db: db}
}

func (ir *InvestorRepository) FetchIdInvestorByEmail(email string) (int, error) {
	var id int
	query := fmt.Sprintf(`SELECT idInvestor FROM tb_investor WHERE investorEmail = '%s' `, email)
	err := ir.db.QueryRow(query).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (ir *InvestorRepository) FetchRoleInvestor(email string) (string, error){
	var role string
	query := fmt.Sprintf(`SELECT investorRole FROM tb_investor WHERE investorEmail = '%s' `, email)
	err := ir.db.QueryRow(query).Scan(&role)

	if err != nil {
		return "", err
	}

	return role, nil
}


func (ir *InvestorRepository) FetchPasswordInvestorByEmail(email string) (entities.LoginInvestor, error) {
	var investor entities.LoginInvestor
	query := fmt.Sprintf(`SELECT investorPassword,investorEmail FROM tb_investor WHERE investorEmail = '%s' `, email)
	err := ir.db.QueryRow(query).Scan(&investor.Password, &investor.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.LoginInvestor{}, errors.New("e-mail não encontrado")
		}
		return entities.LoginInvestor{}, err
	}

	return investor, nil
}

func (ir *InvestorRepository) IsEmailExists(email string) (string, error) {
	var emailInvestor string
	query := fmt.Sprintf(`SELECT investorEmail FROM tb_investor WHERE investorEmail = '%s' `, email)
	err := ir.db.QueryRow(query).Scan(&emailInvestor)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("e-mail não encontrado")
		}
		return "", err
	}

	return emailInvestor, nil
}



func (ir *InvestorRepository) FetchCodeInvestorByEmail(email string) (string, error) {
	var code string
	query := fmt.Sprintf(`SELECT operationCode FROM tb_investor WHERE investorEmail = '%s' `, email)
	err := ir.db.QueryRow(query).Scan(&code)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("e-mail não encontrado")
		}
		return "", err
	}

	return code, nil
}

func (ir *InvestorRepository) UpdatePasswordInvestor(email, newPassword string) error {
	query := `
		UPDATE tb_investor
		SET investorPassword = ?
		WHERE investorEmail = ?`

	_, err := ir.db.Exec(query, newPassword, email)
	if err != nil {
		fmt.Println("error ",err)
		return errors.New("erro ao atualizar a senha")
	}
	return nil
}

func (ir *InvestorRepository) CreateInvestorDB(name, email, password string,repositoryService *sql.Tx) (int,error) {
	tx := repositoryService
	const ROLE = "USER"
	query := ` INSERT INTO tb_investor (InvestorName, InvestorEmail, InvestorPassword, InvestorRole, operationCode, isAccountValid) VALUES (?, ?, ?, ?, ?, ?) `
	
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1,err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(name, email, password, ROLE, "", true)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			if sqlErr.Number == 1062 {
				return -1,errors.New("e-mail já está em uso")
			}
		}
		return -1,errors.New("erro ao criar novo investidor")
	}

	idReturn , errLastId := rows.LastInsertId()
	if errLastId != nil {
		return -1,errLastId
	}

	return int(idReturn),nil
}


func (ir *InvestorRepository) UpdateCodeInvestor(email, code string) error {
	query := `
		UPDATE tb_investor
		SET operationCode = ?
		WHERE investorEmail = ?
`
	_, err := ir.db.Exec(query, code, email)
	if err != nil {
		fmt.Println(err)
		return errors.New("erro ao atualizar código")
	}
	return nil
}

func (ir *InvestorRepository) UpdateDatasInvestor(name string, idInvestor int) error {
	query := `
		UPDATE tb_investor
		SET investorName = ?
		WHERE idInvestor = ?
`
	_, err := ir.db.Exec(query, name, idInvestor)
	if err != nil {
		return errors.New("erro ao criar novo investidor")
	}
	return nil
}


func (ir *InvestorRepository) ChangeAccountStatusInvestor(isAccountValid bool,idInvestor int) error {

	query := ` UPDATE tb_investor SET isAccountValid = ? WHERE idInvestor = ?`
	_, err := ir.db.Exec(query, isAccountValid, idInvestor)
	if err != nil {
		return errors.New("erro ao atualizar status da conta")
	}
	return nil
}

func (ir *InvestorRepository) FetchInvestorEmailAndBalance(id int) (entities.BalanceEmailInvestor, error) {
	var datas entities.BalanceEmailInvestor
	query := fmt.Sprintf(`SELECT ti.investorName,tw.balance FROM tb_wallet AS tw INNER JOIN tb_investor AS ti ON ti.idInvestor = tw.idInvestor WHERE ti.idInvestor = %d; `, id)
	err := ir.db.QueryRow(query).Scan(&datas.Name, &datas.Balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.BalanceEmailInvestor{}, errors.New("investidor não encontrado")
		}
		return entities.BalanceEmailInvestor{}, err
	}

	return datas, nil
}

func (ir *InvestorRepository) FetchAssetQuantity(idInvestor int, assetName string) (int, error) {
	var quantity int
	query := fmt.Sprintf(`SELECT twa.assetQuantity FROM tb_walletAsset AS twa INNER JOIN tb_wallet AS tw ON twa.idWallet = tw.idWallet WHERE tw.idInvestor = %d AND twa.assetName = '%s'; `, idInvestor, assetName)
	err := ir.db.QueryRow(query).Scan(&quantity)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return quantity, nil
}