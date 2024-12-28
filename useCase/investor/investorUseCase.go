package investor

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository/investor_repository"
	"jumpStart-backEnd/security/encryption"
	"jumpStart-backEnd/security/jwt_security"
	 "jumpStart-backEnd/serviceRepository"
	"jumpStart-backEnd/useCase/wallet"
	"jumpStart-backEnd/service/email_service"
	"net/mail"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

type InvestorUseCase struct {
	repo *investor_repository.InvestorRepository
	walletUseCase  *wallet.WalletUseCase
	repositoryService *servicerepository.ServiceRepository
}

func NewInvestorUseCase(repo *investor_repository.InvestorRepository,walletUseCase  *wallet.WalletUseCase,repositoryService *servicerepository.ServiceRepository) *InvestorUseCase {
	return &InvestorUseCase{repo: repo,walletUseCase:walletUseCase,repositoryService:repositoryService}
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
	encryptedPassword, err := encryption.EncryptMessage(key, investor.Password)
	if err != nil {
		return err
	}

	repositoryService,err := iu.repositoryService.StartTransaction()
	if err != nil {
		return errors.New("erro ao processar requisição, tente novamente")
	}

	id,errCreate := iu.repo.CreateInvestorDB(investor.Name, investor.Email, encryptedPassword,repositoryService)
	if errCreate != nil {
		repositoryService.Rollback()
		return errCreate
	}

	errWallet := iu.walletUseCase.CreateWallet(id,repositoryService)
	if errWallet != nil {
		repositoryService.Rollback()
		return errWallet
	}

	errService := repositoryService.Commit()
	if errService != nil {
		repositoryService.Rollback()
		return errors.New("erro ao processar requisição, tente novamente")
	}
	return nil
}

func (iu *InvestorUseCase) LoginInvestor(investor entities.LoginInvestor) (entities.TokenUser, error) {
	if !isEmailValid(investor.Email) {
		return entities.TokenUser{}, errors.New("email inválido")
	}
	if err := isPasswordValid(investor.Password); err != nil {
		return entities.TokenUser{}, err
	}

	passwordDataBase, err := iu.repo.FetchPasswordInvestorByEmail(investor.Email)
	if err != nil {
		if err.Error() == "e-mail não encontrado" {
			return entities.TokenUser{}, err
		}
		return entities.TokenUser{}, errors.New("erro ao realizar o login")
	}

	key := getSecretyKey()
	decryptedPassword, err := encryption.DecryptMessage(key, passwordDataBase.Password)

	if err != nil {
		return entities.TokenUser{}, errors.New("erro ao realizar o login")
	}

	if decryptedPassword != investor.Password {
		return entities.TokenUser{}, errors.New("senha incorreta")
	}

	token, errToken := jwt_security.GenerateToken(investor.Email)

	if errToken != nil {
		return entities.TokenUser{}, errors.New("erro ao realizar o login")
	}

	var tokenInvestor = entities.TokenUser{Token: token}

	return tokenInvestor, nil

}

func (iu *InvestorUseCase) SendCodeToRecoverPassword(email string) error {
	_,errEmailDB := iu.repo.IsEmailExists(email)
	if errEmailDB != nil {
		if errEmailDB.Error() == "e-mail não encontrado" {
			return errors.New("email não encontrado")
		}else{
			return errors.New("ocoreu um erro, tente novamente")
		}
	}
	if !isEmailValid(email){
		return errors.New("email invalido")
	}
	code,errCode := generateRandomString(3)
	if errCode != nil {
		return errors.New("ocoreu um erro, tente novamente"+errCode.Error())
	}

	credentials,errCredentials := recoverCredentialsEmail()
	if errCredentials != nil {
		return errors.New("ocoreu um erro, tente novamente"+errCredentials.Error())
	}

	bodyEmail := "Código para recuperação: "+ code
    
	key,errKey := getKeyEncryption()
	if errKey != nil {
		return errors.New("ocoreu um erro, tente novamente"+errKey.Error())
	}

	codeEncryption,errCrypto := encryption.EncryptMessage(key,code)
	if errCrypto != nil{
		return errors.New("ocoreu um erro, tente novamente"+errCrypto.Error())
	}

	errUpdate := iu.repo.UpdateCodeInvestor(email,codeEncryption)
	if errUpdate != nil {
		return errors.New("ocoreu um erro, tente novamente"+errUpdate.Error())
	}

	errEmail := email_service.SendEmail(email,credentials[0],credentials[1],"Jump start - Código recuperação de senha",bodyEmail)
	if errEmail != nil {
		return errors.New("ocoreu um erro, tente novamente")
	}
	return nil
}

func (iu *InvestorUseCase) VerifyCode(email,code,newPassword string) error {
	_,errEmailDB := iu.repo.IsEmailExists(email)
	if errEmailDB != nil {
		if errEmailDB.Error() == "e-mail não encontrado" {
			return errors.New("email não encontrado")
		}else{
			return errors.New("ocoreu um erro, tente novamente")
		}
	}

	if !isEmailValid(email){
		return errors.New("email invalido")
	}

	codeEncrypted,errDb := iu.repo.FetchCodeInvestorByEmail(email)
	if errDb != nil {
		return errors.New("aconteceu algum erro, tente novamente")
	}

	key,errKey := getKeyEncryption()
	if errKey != nil {
		return errors.New("ocoreu um erro, tente novamente"+errKey.Error())
	}

	codeDescrypted, errDecryp := encryption.DecryptMessage(key,codeEncrypted)
	if errDecryp != nil {
		return errors.New("ocoreu um erro, tente novamente"+errDecryp.Error())
	}
	if codeDescrypted != code {
		return errors.New("código incorreto")
	}
	passwordEncrypted, errEncrypted :=  encryption.EncryptMessage(key,newPassword)
	if errEncrypted != nil {
		return errors.New("ocoreu um erro, tente novamente"+errEncrypted.Error())
	}

	err := iu.repo.UpdatePasswordInvestor(email,passwordEncrypted)
	if err != nil {
		return errors.New("ocoreu um erro, tente novamente"+err.Error())
	}

	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isNameValid(name string) error {
	if name == "" {
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
	if password == "" {
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

func getSecretyKey() []byte {
	err2 := godotenv.Load()
	if err2 != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}
	PASSWORD := os.Getenv("ENCRYPT_KEY")
	jwtSecret := []byte(PASSWORD)

	return jwtSecret
}



func generateRandomString(length int) (string,error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
	   return "",err
	}
	return base64.StdEncoding.EncodeToString(b),nil
 }

 func recoverCredentialsEmail() ([]string,error){
	err2 := godotenv.Load()
	if err2 != nil {
		return nil,err2
	}
	PASSWORD_EMAIL := os.Getenv("PASSWORD_EMAIL")
	ADRESS_EMAIL := os.Getenv("ADRESS_EMAIL")
	
	credentials := []string{ADRESS_EMAIL,PASSWORD_EMAIL}

	return credentials,nil
 }

 func getKeyEncryption() ([]byte,error) {
	err2 := godotenv.Load()
	if err2 != nil {
		return nil,errors.New("ocorreu um erro")
	}
	PASSWORD := os.Getenv("ENCRYPT_KEY")
	key := []byte(PASSWORD)

	return key,nil
 }

