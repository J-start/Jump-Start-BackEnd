package main

import (
	"jumpStart-backEnd/controller"
	"jumpStart-backEnd/db"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase/buy"
	"log"
	"net/http"
	//"jumpStart-backEnd/entities"
	"jumpStart-backEnd/useCase"
	"jumpStart-backEnd/useCase/wallet"
)

func main() {

	//   asset := entities.AssetOperation{
	//   AssetName: "CHZ",
	//   AssetCode: "CHZ-BRL",
	//   AssetType: "CRYPTO",
	//   AssetAmount: 10,
	//   OperationType: "SELL",
	//   CodeInvestor: "123456",
	//    }

	//  asset := entities.AssetOperation{
	//  	AssetName: "PETROBRAS PN (PETR4.SA)",
	//  	AssetCode: "PETR4.SA",
	//  	AssetType: "SHARE",
	//  	AssetAmount: 10,
	//  	OperationType: "SELL",
	//  	CodeInvestor: "123456",
	//  }

	db := db.GetDB()
	shareRepository := repository.NewShareRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	operationAssetRepository := repository.NewOperationAssetRepository(db)

	walletUseCase := wallet.NewWalletUseCase(walletRepository)
	shareUsecase := usecase.NewShareUseCase(shareRepository)
	newBuyAssetsUseCase := buy.NewBuyAssetsUseCase(shareRepository, shareUsecase, walletUseCase, operationAssetRepository)
	BuyAssetController := controller.NewBuyAssetController(newBuyAssetsUseCase)

	shareController := controller.NewShareController(shareUsecase)

	http.HandleFunc("/datas/shares", shareController.GetTodaySharesJSON)
	http.HandleFunc("/datas/shares/offset", shareController.GetSharesSpecifyOffSet)
	http.HandleFunc("/data/share/", shareController.GetShareById)
	http.HandleFunc("/datas/share/", shareController.GetShareList)
	http.HandleFunc("/buy/", BuyAssetController.BuyAsset)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
