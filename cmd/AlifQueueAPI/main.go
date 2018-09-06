package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlifElectronicQueue/internal/app/admin"
	"github.com/AlifElectronicQueue/internal/app/authentication"
	"github.com/AlifElectronicQueue/internal/app/users"
	"github.com/AlifElectronicQueue/internal/pkg/databaseinit"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLogging()
	log.Debug("Trying to initializa db connection!")

	//dbProvider := "???"

	DataAccess := databaseinit.SetDriverName(dbProvider)
	defer DataAccess.Disconnect()

	// adRepo := databaseinit.CreateAdminRepository(dbProvider, DataAccess.ConVar)
	// adSrv := admin.InitService(adRepo)
	// adContrl := admin.InitControllers(adSrv)

	authRepo := databaseinit.CreateAuthenticationRepository(dbProvider, DataAccess.ConVar)
	authSrv := authentication.InitService(authRepo)
	authContrl := authentication.InitControllers(authSrv)

	// uRepo := databaseinit.CreateUsersRepository(dbProvider, DataAccess.ConVar)
	// uSrv := users.InitService(uRepo)
	// uContrl := users.InitControllers(uSrv)
	// router.HandleFunc("/signup", catContrl.AdminSignUp())
	// router.HandleFunc("/update", catContrl.AdminUpdate())
	// router.HandleFunc("/delete", catContrl.AdminDelete())
	

	router := mux.NewServeMux()

	//admin
	router.HandleFunc("/admin/applications", authContrl.AdminSignIn())

	//Users
	router.HandleFunc("/v1/login/signin", authContrl.UserSignIn())
	//router.HandleFunc("/v1/users/catalogue/signup", catContrl.UserSignUp())

	//!**************************************************************************************************************/
	//? Public Routes
	//*Root Router
	router.HandleFunc("/",)

	//*Admin Login Router
	router.HandleFunc("/login").Methods("POST")

	//?Private Routes
	//*Admin page Router
	router.HandleFunc("/admin/applications")

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
