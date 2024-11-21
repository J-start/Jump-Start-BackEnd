package main

import (
	// "jumpStart-backEnd/controller"
	  "jumpStart-backEnd/db"
	  "jumpStart-backEnd/repository"
	// "log"
	// "net/http"

	"jumpStart-backEnd/entities"
	"jumpStart-backEnd/useCase"
	"fmt"
	"time"
)


func main() {
    now := time.Now()
    fmt.Println(now)
	fmt.Println("dia: ", now.Weekday())
	

	    asset := entities.AssetOperation{
	    AssetName: "CHZ",
	   	AssetCode: "CHZ-BRL",
 	 	AssetType: "CRYPTO",
	    AssetAmount: 10,
	    OperationType: "SELL",
	    CodeInvestor: "123456",
	     }

	//    asset := entities.AssetOperation{
	//    	AssetName: "PETROBRAS PN (PETR4.SA)",
	//    	AssetCode: "PETR4.SA",
	//    	AssetType: "SHARE",
	//    	AssetAmount: 10,
	//    	OperationType: "SELL",
	//    	CodeInvestor: "123456",
	//    }

	  db := db.GetDB()
      shareRepository := repository.NewShareRepository(db)
	  shareUseCase := usecase.NewSellAssetsUseCase(shareRepository)
	  shareUseCase.ManipulationAsset(asset)

	//  shareController := controller.NewShareController(shareUseCase)

	//  http.HandleFunc("/datas/shares", shareController.GetTodaySharesJSON)
	//  http.HandleFunc("/datas/shares/offset", shareController.GetSharesSpecifyOffSet)
	//  http.HandleFunc("/data/share/", shareController.GetShareById)
	//  http.HandleFunc("/datas/share/", shareController.GetShareList)
	//  log.Fatal(http.ListenAndServe(":8080", nil))

}
