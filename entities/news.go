package entities

import "time"

type News struct {
	Id         int
	News       string
	DateNews   time.Time
	IsApproved bool
}

type NewsStructure struct {
	Id   int
	News string
	DateNews string
}

type NewsDelete struct {
	TokenInvestor string
	IdNews 	      int
}