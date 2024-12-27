package main

import (
	"jumpStart-backEnd/controller"
	"jumpStart-backEnd/db"
	"jumpStart-backEnd/repository"
	"jumpStart-backEnd/service/investor_service"
	"jumpStart-backEnd/serviceRepository"
	"jumpStart-backEnd/useCase"
	"jumpStart-backEnd/useCase/assetwallet"
	"jumpStart-backEnd/useCase/buy"
	"jumpStart-backEnd/useCase/investor"
	"jumpStart-backEnd/useCase/listasset"
	"jumpStart-backEnd/useCase/news"
	"jumpStart-backEnd/useCase/operation"
	"jumpStart-backEnd/useCase/sell"
	"jumpStart-backEnd/useCase/wallet"
	"log"
	"net/http"
)

func main() {

	db := db.GetDB()

	shareRepository := repository.NewShareRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	operationAssetRepository := repository.NewOperationAssetRepository(db)
	assetWalletRepository := repository.NewWalletAssetRepository(db)
	serviceRepository := servicerepository.NewWServiceRepository(db)
	listAssetRepository := repository.NewListAssetRepository(db)
	newsRepository := repository.NewNewsRepository(db)
	investorRepository := repository.NewInvestorRepository(db)

	investorService := investor_service.NewInvestorService(investorRepository)

	listAssetUseCase := listasset.NewListAssetUseCase(listAssetRepository)
	
	newsUseCase := news.NewNewsUseCase(newsRepository,investorService)
	assetWalletUseCase := assetwallet.NewAssetWalletUseCase(assetWalletRepository)
	operationAssetUseCase := operation.NewOperationAssetUseCase(operationAssetRepository,investorService)
	walletUseCase := wallet.NewWalletUseCase(walletRepository,operationAssetUseCase,serviceRepository,investorService)
	investorUseCase := investor.NewInvestorUseCase(investorRepository,walletUseCase,serviceRepository)
	shareUsecase := usecase.NewShareUseCase(shareRepository)
	newBuyAssetsUseCase := buy.NewBuyAssetsUseCase(shareRepository, shareUsecase, walletUseCase, 
													  operationAssetUseCase,assetWalletUseCase,serviceRepository,investorService)
    NewSellAssetsUseCase := sell.NewSellAssetsUseCase(shareRepository, shareUsecase, walletUseCase, 
													  operationAssetUseCase,assetWalletUseCase,serviceRepository,investorService)
	
	investorController := controller.NewInvestorController(investorUseCase)	
	newsController := controller.NewNewsController(newsUseCase)
	listAssetController := controller.NewListAssetController(listAssetUseCase)
	operationAssetController := controller.NewOperationAssetController(operationAssetUseCase)
	BuyAssetController := controller.NewBuyAssetController(newBuyAssetsUseCase)
	sellAssetController := controller.NewSellAssetController(NewSellAssetsUseCase)
	shareController := controller.NewShareController(shareUsecase)
	walletController := controller.NewWalletController(walletUseCase)

	http.HandleFunc("/datas/shares", shareController.GetTodaySharesJSON)
	http.HandleFunc("/datas/shares/offset", shareController.GetSharesSpecifyOffSet)
	http.HandleFunc("/data/share/", shareController.GetShareById)
	http.HandleFunc("/datas/share/", shareController.GetShareList)
	http.HandleFunc("/buy/", BuyAssetController.BuyAsset)//*
	http.HandleFunc("/sell/", sellAssetController.SellAsset)//*
	http.HandleFunc("/details/asset/", listAssetController.ListAsset)
	http.HandleFunc("/asset/request/", listAssetController.ListAssetRequest)
	http.HandleFunc("/history/assets/", operationAssetController.FetchHistoryOperationInvestor)//*
	http.HandleFunc("/wallet/datas/", walletController.FetchDatasWallet)//*
	http.HandleFunc("/history/operations/", walletController.FetchOperationsWallet)//*
	http.HandleFunc("/withdraw/", walletController.WithDraw)//*
	http.HandleFunc("/deposit/", walletController.Deposit)//*
	http.HandleFunc("/news/", newsController.FetchNews)
	http.HandleFunc("/news/delete/", newsController.DeleteNews)
	http.HandleFunc("/investor/create/", investorController.CreateInvestor) //*
	http.HandleFunc("/investor/login/", investorController.Login)
	http.HandleFunc("/investor/password/code/", investorController.SendCodeEmailRecoverPassword)
	http.HandleFunc("/investor/password/update/", investorController.VerifyCodeEmail)

	log.Fatal(http.ListenAndServe(":8080", nil))


}
