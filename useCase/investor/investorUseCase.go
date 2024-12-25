package investor

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/security/encryption"
	"net/mail"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type InvestorUseCase struct {
	repo *repository.InvestorRepository
}

func NewInvestorUseCase(repo *repository.InvestorRepository) *InvestorUseCase {
	return &InvestorUseCase{repo: repo}
}

func (iu *InvestorUseCase) CreateInvestor(investor entities.InvestorInsert) error {
	if err := validateFields(investor.Name, investor.Email, investor.Password); err != nil {
		return err
	}
	key := getSecretyKey()
	encryptedPassword,err := encryption.EncryptMessage(key,investor.Password)
	if err != nil {
		return err
	}
	errCreate := iu.repo.CreateInvestorDB(investor.Name, investor.Email, encryptedPassword)
	if errCreate != nil {
		return errCreate
	}
	return nil
}


func validateFields(name,email,password string) error {
	if name == "" || email == "" || password == "" {
		return errors.New("todos os campos devem ser preenchidos")
	}
	if strings.Trim(name, " ") == "" || strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
		return errors.New("todos os campos devem ser preenchidos")
	}
	if len(name) < 3 || len(name) > 50 {
		return errors.New("nome deve ter entre 3 e 50 caracteres")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("senha deve ter entre 6 e 50 caracteres")
	}
	if !isEmailValid(email) {
		return errors.New("email inválido")
	}
	return nil
}

func isEmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func getSecretyKey() []byte{
	err2 := godotenv.Load()
    if err2 != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
    }
	PASSWORD := os.Getenv("ENCRYPT_KEY")
	jwtSecret := []byte(PASSWORD)

	return jwtSecret
}