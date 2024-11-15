package entities

import "time"

type News struct {
	Id         int
	News       string
	DateNews   time.Time
	IsApproved bool
}