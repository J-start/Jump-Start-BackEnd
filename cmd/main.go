package main

import (
	"jumpStart-backEnd/controller"
	"jumpStart-backEnd/db"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/useCase/assetwallet"
	"jumpStart-backEnd/useCase/buy"
	"jumpStart-backEnd/useCase/sell"
	"log"
	"net/http"
	"jumpStart-backEnd/useCase"
	"jumpStart-backEnd/useCase/operation"
	"jumpStart-backEnd/useCase/wallet"
	"jumpStart-backEnd/serviceRepository"
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
	assetWalletRepository := repository.NewWalletAssetRepository(db)
	serviceRepository := servicerepository.NewWServiceRepository(db)

	assetWalletUseCase := assetwallet.NewAssetWalletUseCase(assetWalletRepository)
	operationAssetUseCase := operation.NewOperationAssetUseCase(operationAssetRepository)
	walletUseCase := wallet.NewWalletUseCase(walletRepository,operationAssetUseCase)
	shareUsecase := usecase.NewShareUseCase(shareRepository)
	newBuyAssetsUseCase := buy.NewBuyAssetsUseCase(shareRepository, shareUsecase, walletUseCase, operationAssetUseCase,assetWalletUseCase,serviceRepository)
	NewSellAssetsUseCase := sell.NewSellAssetsUseCase(shareRepository, shareUsecase, walletUseCase, operationAssetUseCase,assetWalletUseCase,serviceRepository)
	
	BuyAssetController := controller.NewBuyAssetController(newBuyAssetsUseCase)
	sellAssetController := controller.NewSellAssetController(NewSellAssetsUseCase)
	shareController := controller.NewShareController(shareUsecase)

	http.HandleFunc("/datas/shares", shareController.GetTodaySharesJSON)
	http.HandleFunc("/datas/shares/offset", shareController.GetSharesSpecifyOffSet)
	http.HandleFunc("/data/share/", shareController.GetShareById)
	http.HandleFunc("/datas/share/", shareController.GetShareList)
	http.HandleFunc("/buy/", BuyAssetController.BuyAsset)
	http.HandleFunc("/sell/", sellAssetController.SellAsset)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
