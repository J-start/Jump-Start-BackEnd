package main

import (
	"jumpStart-backEnd/controller"
	"jumpStart-backEnd/db"
	"jumpStart-backEnd/repository"
	"log"
	"net/http"

	"jumpStart-backEnd/useCase"
)


func main() {
	db := db.GetDB()
	newsRepository := repository.NewShareRepository(db)
	shareUseCase := usecase.NewShareUseCase(newsRepository)

	shareController := controller.NewShareController(shareUseCase)

	http.HandleFunc("/datas/shares", shareController.GetTodaySharesJSON)
	http.HandleFunc("/datas/shares/offset", shareController.GetSharesSpecifyOffSet)
	http.HandleFunc("/data/share/", shareController.GetShareById)
	http.HandleFunc("/datas/share/", shareController.GetShareList)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
