package databaseinit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataAccess struct {
	ConfigFilePath string
	ConVar         *sqlx.DB
	DriverName     string
}

func (db *DataAccess) InitDataBase() error {
	var dbData types.DBData
	slbyte, _ := ioutil.ReadFile(db.ConfigFilePath)

	err := json.Unmarshal(slbyte, &dbData) // decoding JSON data into the struct
	if err != nil {
		return errors.New("Can not decode from JSON file!")
	}
	dbInfo := dbData.DataBaseUser + `://` + dbData.DataBaseUser + `:` + dbData.DataBasePassword + `@192.168.202.10/` + dbData.DataBaseName
	db.ConVar = sqlx.MustConnect("postgres", dbInfo)

	if err = db.ConVar.Ping(); err != nil {
		return err
	}
	fmt.Println("\n PINGED!")

	return nil
}

func (db *DataAccess) Disconnect() {
	fmt.Println("DISCONNECTED!")
	db.ConVar.Close()
}

func (db *DataAccess) GetDriverName() string {
	return db.DriverName
}

func SetDriverName(DriverName string) *DataAccess {
	var db = DataAccess{
		ConfigFilePath: "../../configs/dbDataAccess.json",
		DriverName:     DriverName,
	}
	switch DriverName {
	case "postgres":
		err := db.InitDataBase()
		if err != nil {
			fmt.Println(err)
		}
		// case "mysql":
	default:
		fmt.Println("Driver name undefined")
	}
	return &db
}
