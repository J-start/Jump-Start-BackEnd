package main

import (
	"jumpStart-backEnd/db"
	"jumpStart-backEnd/repository"
	
)

func main() {
	db := db.GetDB()
	newsRepository := repository.NewNewsRepository(db)
	newsRepository.FindAllNews()
	
}

