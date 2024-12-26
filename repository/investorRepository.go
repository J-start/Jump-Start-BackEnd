package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
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

func (ir *InvestorRepository) CreateInvestorDB(name, email, password string) error {
	const ROLE = "USER"
	query := `
		INSERT INTO tb_investor 
		(InvestorName, InvestorEmail, InvestorPassword, InvestorRole, operationCode, isAccountValid)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := ir.db.Exec(query, name, email, password, ROLE, "", false)
	if err != nil {
		if sqlErr, ok := err.(*mysql.MySQLError); ok {
			if sqlErr.Number == 1062 {
				return errors.New("e-mail já está em uso")
			}
		}
		return errors.New("erro ao criar novo investidor")
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

func (ir *InvestorRepository) UpdatePasswordInvestor(password string, idInvestor int) error {
	query := `
		UPDATE tb_investor
		SET investorPassword = ?
		WHERE idInvestor = ?
`
	_, err := ir.db.Exec(query, password, idInvestor)
	if err != nil {
		return errors.New("erro ao atualiza senha")
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