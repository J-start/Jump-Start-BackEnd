package investor

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/security/encryption"
	"jumpStart-backEnd/security/jwt"
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
	if !isEmailValid(investor.Email) {
		return errors.New("email inválido")
	}
	if err := isNameValid(investor.Name); err != nil {
		return err
	}
	if err := isPasswordValid(investor.Password); err != nil {
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

func (iu *InvestorUseCase) LoginInvestor(investor entities.LoginInvestor) (entities.TokenUser,error) {
	if !isEmailValid(investor.Email){
		return entities.TokenUser{},errors.New("email inválido")
	}
	if err := isPasswordValid(investor.Password); err != nil {
		return entities.TokenUser{},err
	}

	
    passwordDataBase,err := iu.repo.FetchPasswordInvestorByEmail(investor.Email)
	if err != nil {
		if err.Error() == "e-mail não encontrado"{
			return entities.TokenUser{},err
		}
		return entities.TokenUser{},errors.New("erro ao realizar o login")
	}

	key := getSecretyKey()
	decryptedPassword,err := encryption.DecryptMessage(key,passwordDataBase.Password)

	if err != nil {
		return entities.TokenUser{},errors.New("erro ao realizar o login")
	}

	if decryptedPassword != investor.Password{
		return entities.TokenUser{},errors.New("senha incorreta")
	}

	if investor.Email != passwordDataBase.Email{
		return entities.TokenUser{},errors.New("email incorreto")
	}

	token,errToken := jwt.GenerateToken(investor.Email)

	if errToken != nil {
		fmt.Println(errToken)
		return entities.TokenUser{},errors.New("erro ao realizar o login")
	}

	var tokenInvestor = entities.TokenUser{Token: token}

	return tokenInvestor,nil
	
}


func isEmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func isNameValid(name string) error {
	if name == ""{
		return errors.New("nome vazio")
	}
	if len(name) < 3 || len(name) > 50 {
		return errors.New("o nome deve ter entre 3 e 50 caracteres")
	}
	if strings.Trim(name, " ") == "" {
		return errors.New("o nome deve ter entre 3 e 50 caracteres")
	}
	return nil
}

func isPasswordValid(password string) error {
	if password == ""{
		return errors.New("senha vazia")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("senha deve ter entre 8 e 30 caracteres")
	}
	if strings.Trim(password, " ") == "" {
		return errors.New("senha deve ter entre 8 e 30 caracteres")
	}
	return nil
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