package investor

import (
	"errors"
	"fmt"
	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/security/encryption"
	"jumpStart-backEnd/security/jwt_security"
	"jumpStart-backEnd/service/email_service"
	"jumpStart-backEnd/service/investor_service"
	"jumpStart-backEnd/serviceRepository"
	"jumpStart-backEnd/useCase/wallet"
	"net/mail"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type InvestorUseCase struct {
	repo              *repository.InvestorRepository
	walletUseCase     *wallet.WalletUseCase
	repositoryService *servicerepository.ServiceRepository
	investorService   *investor_service.InvestorService
}

type Role struct {
	IsAdm bool `json:"isAdm"`
}

func NewInvestorUseCase(repo *repository.InvestorRepository, walletUseCase *wallet.WalletUseCase, repositoryService *servicerepository.ServiceRepository, investorService *investor_service.InvestorService) *InvestorUseCase {
	return &InvestorUseCase{repo: repo, walletUseCase: walletUseCase, repositoryService: repositoryService, investorService: investorService}
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

	repositoryService, err := iu.repositoryService.StartTransaction()
	if err != nil {
		return errors.New("erro ao processar requisição, tente novamente")
	}

	id, errCreate := iu.repo.CreateInvestorDB(investor.Name, investor.Email, encryptedPassword, repositoryService)
	if errCreate != nil {
		repositoryService.Rollback()
		return errCreate
	}

	errWallet := iu.walletUseCase.CreateWallet(id, repositoryService)
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

func (iu *InvestorUseCase) SendUrlToRecoverPassword(email string) error {
	if !isEmailValid(email) {
		return errors.New("email invalido")
	}
	_, errEmailDB := iu.repo.IsEmailExists(email)
	if errEmailDB != nil {
		if errEmailDB.Error() == "e-mail não encontrado" {
			return errors.New("email não encontrado")
		} else {
			return errors.New("ocoreu um erro, tente novamente")
		}
	}

	credentials, errCredentials := recoverCredentialsEmail()
	if errCredentials != nil {
		return errors.New("ocoreu um erro, tente novamente" + errCredentials.Error())
	}

	token, errToken := jwt_security.GenerateTokenWithNMinutes(email, 15)

	if errToken != nil {
		return errors.New("erro ao gerar url para recuperação de senha")
	}

	url := fmt.Sprintf("https://jumpstart.dev.br/recoverPassword.html?token=%s", token)

	bodyEmail := fmt.Sprintf("Para atualizar sua senha acesse o seguinte link para a plataforma jumpStart: %s", url)

	errEmail := email_service.SendEmail(email, credentials[0], credentials[1], "Jump start - Recuperação de senha", bodyEmail)
	if errEmail != nil {
		return errors.New("ocoreu um erro, tente novamente")
	}
	return nil
}

func (iu *InvestorUseCase) VerifyToken(token string) error {
	_, err := jwt_security.ValidateToken(token)
	if err != nil {
		return err
	}

	return nil
}

func (iu *InvestorUseCase) UpdatePasswordInvestor(token string, newPassword string) error {
	email, err := jwt_security.ValidateToken(token)
	if err != nil {
		return err
	}
	if !isEmailValid(email.UserEmail) {
		return errors.New("email invalido")
	}
	_, errEmailDB := iu.repo.IsEmailExists(email.UserEmail)
	if errEmailDB != nil {
		if errEmailDB.Error() == "e-mail não encontrado" {
			return errors.New("usuário não encontrado")
		} else {
			return errors.New("ocoreu um erro, tente novamente")
		}
	}

	key, errKey := getKeyEncryption()
	if errKey != nil {
		return errors.New("ocoreu um erro, tente novamente" + errKey.Error())
	}

	passwoesEncrypted, errCrypto := encryption.EncryptMessage(key, newPassword)
	if errCrypto != nil {
		return errors.New("ocoreu um erro, tente novamente" + errCrypto.Error())
	}

	errDb := iu.repo.UpdatePasswordInvestor(email.UserEmail, passwoesEncrypted)
	if errDb != nil {
		return errors.New("ocoreu um erro ao atualiza senha, tente novamente")
	}

	return nil

}

func (iu *InvestorUseCase) NameAndBalanceInvestor(token string) (entities.BalanceEmailInvestor, error) {
	idInvestor, err := iu.investorService.GetIdByToken(token)
	if err != nil {
		return entities.BalanceEmailInvestor{}, errors.New("token inválido, realize o login novamente")
	}
	datas, errDb := iu.repo.FetchInvestorEmailAndBalance(idInvestor)
	if errDb != nil {
		return entities.BalanceEmailInvestor{}, errors.New("erro ao buscar dados do investidor")
	}
	return datas, nil
}
func (iu *InvestorUseCase) GetAssetsQuantity(token string, nameAsset string) (entities.QuantityInvestorAsset, error) {
	idInvestor, err := iu.investorService.GetIdByToken(token)
	if err != nil {
		return entities.QuantityInvestorAsset{}, errors.New("token inválido, realize o login novamente")
	}
	quantity, errDb := iu.repo.FetchAssetQuantity(idInvestor, nameAsset)
	if errDb != nil {
		return entities.QuantityInvestorAsset{}, errors.New(`erro ao buscar quantidade do ativo ` + nameAsset)
	}
	datas := entities.QuantityInvestorAsset{
		Quantity: quantity,
	}
	return datas, nil
}

func (iu *InvestorUseCase) GetdatasInvestor(token string) (entities.DatasInvestor, error) {
	idInvestor, err := iu.investorService.GetIdByToken(token)
	if err != nil {
		return entities.DatasInvestor{}, errors.New("token inválido, realize o login novamente")
	}

	datas, errDb := iu.repo.FetchDatasInvestor(idInvestor)
	if errDb != nil {
		return entities.DatasInvestor{}, errors.New("erro ao buscar dados do investidor")
	}
	return datas, nil
}

func (iu *InvestorUseCase) UpdateDatasInvestor(token string, datas entities.DatasInvestor) error {
	idInvestor, err := iu.investorService.GetIdByToken(token)
	if err != nil {
		return errors.New("token inválido, realize o login novamente")
	}
	if !isEmailValid(datas.Email) {
		return errors.New("email inválido")
	}
	if err := isNameValid(datas.Name); err != nil {
		return err
	}

	errDb := iu.repo.UpdateDatasInvestor(datas, idInvestor)
	if errDb != nil {
		return errors.New("erro ao atualizar dados do investidor")
	}
	return nil
}

func (iu *InvestorUseCase) IsAdm(token string) (Role, error) {
	isAdm, err := iu.investorService.IsAdm(token)
	if err != nil {
		return Role{}, errors.New("erro ao verificar permissão")
	}
	var role Role = Role{
		IsAdm: isAdm,
	}

	return role, nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isNameValid(name string) error {
	if name == "" {
		return errors.New("nome vazio")
	}
	if len(name) < 3 || len(name) > 30 {
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

func recoverCredentialsEmail() ([]string, error) {
	err2 := godotenv.Load()
	if err2 != nil {
		fmt.Println(err2) 
	}
	PASSWORD_EMAIL := os.Getenv("PASSWORD_EMAIL")
	ADRESS_EMAIL := os.Getenv("ADRESS_EMAIL")

	credentials := []string{ADRESS_EMAIL, PASSWORD_EMAIL}

	return credentials, nil
}

func getKeyEncryption() ([]byte, error) {
	err2 := godotenv.Load()
	if err2 != nil {
		fmt.Println(err2) 
	}
	PASSWORD := os.Getenv("ENCRYPT_KEY")
	key := []byte(PASSWORD)

	return key, nil
}
