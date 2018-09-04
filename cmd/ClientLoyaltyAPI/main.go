package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/admin"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/analytics"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/authentication"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/cashback"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/catalogue"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/discount"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/app/users"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/databaseinit"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLogging()
	log.Debug("Trying to initializa db connection!")

	dbProvider := "postgres"

	DataAccess := databaseinit.SetDriverName(dbProvider)
	defer DataAccess.Disconnect()

	aRepo := databaseinit.CreateAnalyticsRepository(dbProvider, DataAccess.ConVar)
	aSrv := analytics.InitService(aRepo)
	aContrl := analytics.InitControllers(aSrv)

	adRepo := databaseinit.CreateAdminRepository(dbProvider, DataAccess.ConVar)
	adSrv := admin.InitService(adRepo)
	adContrl := admin.InitControllers(adSrv)

	authRepo := databaseinit.CreateAuthenticationRepository(dbProvider, DataAccess.ConVar)
	authSrv := authentication.InitService(authRepo)
	authContrl := authentication.InitControllers(authSrv)

	cRepo := databaseinit.CreateCashbackRepository(dbProvider, DataAccess.ConVar)
	cSrv := cashback.InitService(cRepo)
	cContrl := cashback.InitControllers(cSrv)

	catRepo := databaseinit.CreateCatalogueRepository(dbProvider, DataAccess.ConVar)
	catSrv := catalogue.InitService(catRepo)
	catContrl := catalogue.InitControllers(catSrv)

	dRepo := databaseinit.CreateDiscountRepository(dbProvider, DataAccess.ConVar)
	dSrv := discount.InitService(dRepo)
	dContrl := discount.InitControllers(dSrv)

	uRepo := databaseinit.CreateUsersRepository(dbProvider, DataAccess.ConVar)
	uSrv := users.InitService(uRepo)
	uContrl := users.InitControllers(uSrv)

	router := http.NewServeMux()

	//Analytic
	router.HandleFunc("/v1/users/cashback/exchange", cContrl.BonusExchange())
	router.HandleFunc("/v1/users/cashback/tranpakupki", cContrl.TranInfoUserPakupki2())
	router.HandleFunc("/v1/users/cashback/tranperevodi", cContrl.TranInfoUserPerevodi2())

	router.HandleFunc("/v1/operator/userInfo/{id}", adContrl.GetUserInfo2())
	router.HandleFunc("/v1/bonus/purchase", cContrl.PurchaseByBonus())
	router.HandleFunc("/v1/cash/purchase", cContrl.PurchaseByCash())
	router.HandleFunc("/v1/analytics/serviceProvider/purchase/all", aContrl.ServiceAllPurchases())

	router.HandleFunc("/v1/catalogue/serviceProvider/rating", aContrl.ShowServiceProvidersRating())
	router.HandleFunc("/v1/analytics/serviceProvider/report/annual", aContrl.HandleGetServicePeroviderAnnualReport())

	//catalogue
	router.HandleFunc("/v1/catalogue/serviceProvider/signin", authContrl.ServiceProviderSignIn())
	router.HandleFunc("/v1/catalogue/serviceProvider/signup", catContrl.ServiceProviderSignUp())
	router.HandleFunc("/v1/catalogue/serviceProvider/update", catContrl.Update())
	router.HandleFunc("/v1/catalogue/serviceProvider/delete", catContrl.ServiceProviderDelete())

	router.HandleFunc("/v1/catalogue/serviceProvider/get/all", catContrl.GetAllServiceProvader())
	router.HandleFunc("/v1/catalogue/serviceProvider/get/id", catContrl.GetServiceProviderByID())

	router.HandleFunc("/v1/catalogue/serviceProvider/usersactivity", aContrl.UsersActivity())
	//Discounts
	router.HandleFunc("/v1/discount/list/all", dContrl.DiscountsList())
	router.HandleFunc("/v1/discount/list/active", dContrl.DiscountsActiveList())
	router.HandleFunc("/v1/discount/list/soon", dContrl.DiscountsSoonList())
	router.HandleFunc("/v1/discount/list/past", dContrl.DiscountsPastList())

	//?Users
	router.HandleFunc("/v1/users/catalogue/signin", authContrl.UserSignIn())
	router.HandleFunc("/v1/users/catalogue/signup", catContrl.UserSignUp())

	router.HandleFunc("/v1/users/get/all/serviceProvider", uContrl.GetAllActiveServiceProvider())
	router.HandleFunc("/v1/users/get/all/activediscounts", uContrl.GetAllActiveDiscounts())
	router.HandleFunc("/v1/users/get/list/newdiscounts", uContrl.GetListNewDiscounts())

	router.HandleFunc("/v1/users/show/all/services", uContrl.ShowInfoServices())
	router.HandleFunc("/v1/users/story/cashback", uContrl.StoryCashback())

	// Without logic or with error
	// router.HandleFunc("v1/users/list/InTouch/product", uContrl.ShowListInTouchProduct()) //!Non-working
	// router.HandleFunc("v1/cashback/get/tranpakupki", cContrl.GetTransactionListPakupki2())//?!
	// router.HandleFunc("v1/cashback/get/tranperevodi", cContrl.GetTransactionListPerevodi2())//?!
	// router.HandleFunc("v1/user/cashback/transfer/", cContrl.CashBounusTransfer())//?!
	//router.HandleFunc("/analytics/get/high/cashback", aContrl.GetMultyHighCashback())

	log.Info("Starting http server...")
	http.ListenAndServe(":80", router)

}

func initLogging() {

	file, err := os.OpenFile("../../logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
}
