package entities

type InvestorInsert struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}