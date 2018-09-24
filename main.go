package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlifElectronicQueue/internal/app/authentication"
	"github.com/AlifElectronicQueue/internal/pkg/databaseinit"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLogging()
	log.Debug("Trying to initializa db connection!")

	dbProvider := "postgres"

	DataAccess := databaseinit.SetDriverName(dbProvider)
	defer DataAccess.Disconnect()

	authRepo := databaseinit.CreateAuthenticationRepository(dbProvider, DataAccess.ConVar)
	authSrv := authentication.InitService(authRepo)
	authContrl := authentication.InitControllers(authSrv)
	authMiddle := authentication.InitMiddlewares(authContrl) //? CREATE MIDDLEWARE

	router := mux.NewRouter()
	//?/REFACTOR ERRORS,SQL,MIDDLEWARE
	//!------------------------------------------------------
	router.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", http.FileServer(http.Dir("./web/static/"))))
	//!--------------------------------------------------/
	//router.HandleFunc("/")
	router.HandleFunc("/login", authContrl.AdminLogin()).Methods("GET")
	router.HandleFunc("/login", authContrl.AdminLogin()).Methods("POST") //TODO:Authentication Process

	router.HandleFunc("/logout", authMiddle.RequiresLogin(authContrl.AdminLogout())).Methods("POST") //TODO: Destroy COOKIE

	router.HandleFunc("/admin/applications", authMiddle.RequiresLogin(authContrl.SelectUsers())) //TODO:Middleware->SecretPage*
	//!!!!!!!!!!!!!!!!!!!!!!!ASK SENIOR
	router.HandleFunc("/admin/applications/{id:[0-9]+}", authMiddle.RequiresLogin(authContrl.SelectUserById())).Methods("GET")                                //TODO:Middleware->SecretPage*
	router.HandleFunc("/admin/applications/{contact}", authMiddle.RequiresLogin(authContrl.SelectUserByContact())).Queries("getby", "contact").Methods("GET") //TODO:Middleware->SecretPage*

	router.HandleFunc("/admin/applications/", authMiddle.RequiresLogin(authContrl.OrderedApplication())).Queries("orderby", "{order}").Methods("GET")
	//router.HandleFunc("/admin/applications/users", authContrl.SelectUsers()).Methods("GET")      //!CHECK

	router.HandleFunc("/", authContrl.Application())
	// router.HandleFunc("/admin/applications", authMiddle.RequiresLogin(authContrl.Application()))//.Methods("POST")
	//!--------------------------------------------------/
	log.Info("Starting http server...")
	http.ListenAndServe(":80", router)

}

func initLogging() {

	file, err := os.OpenFile("./logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //TODO: Changed path!
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
}
