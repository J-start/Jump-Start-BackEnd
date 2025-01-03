package investor_service

import (
	"errors"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/security/jwt_security"
	
)


type InvestorService struct {
	repo *repository.InvestorRepository
}

func NewInvestorService(repo *repository.InvestorRepository) *InvestorService {
	return &InvestorService{repo: repo}
}

func (is *InvestorService) GetIdByToken(token string) (int,error) {
	email,err := jwt_security.ValidateToken(token)
	if err != nil {
		return -1,errors.New("token inválido")
	}
	id,err := is.repo.FetchIdInvestorByEmail(email.UserEmail)
	if err != nil {
		return -1,err
	}
	return id,nil
}

func (is *InvestorService) IsAdm(token string) (bool,error) {
	email,err := jwt_security.ValidateToken(token)
	if err != nil {
		return false,errors.New("token inválido")
	}
	role,err := is.repo.FetchRoleInvestor(email.UserEmail)
	if err != nil {
		return false,errors.New("erro ao buscar papel do investidor")
	}
	if role == "ADMIN" {
		return true,nil
	}
	return false,nil
}




