package entities

type InvestorInsert struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginInvestor struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenUser struct {
	Token string `json:"token"`
}