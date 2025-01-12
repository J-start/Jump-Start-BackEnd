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

type SendCodeEmail struct {
	Email string `json:"email"`
}

type CodeChangePassword struct {
	Email string `json:"email"`
	Code string `json:"code"`
	NewPassword string `json:"newPassword"`
}

type BalanceEmailInvestor struct {
	Name string `json:"name"`
	Balance string `json:"balance"`
}

type QuantityInvestorAsset struct {
	Quantity int `json:"quantity"`
}

