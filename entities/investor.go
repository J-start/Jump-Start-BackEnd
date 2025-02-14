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
	Token string `json:"token"`
}

type UpdatePassword struct {
	Token string `json:"token"`
	NewPassword string `json:"newPassword"`
}

type BalanceEmailInvestor struct {
	Name string `json:"name"`
	Balance string `json:"balance"`
}

type QuantityInvestorAsset struct {
	Quantity float64 `json:"quantity"`
}
type DatasInvestor struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

