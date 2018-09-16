package databaseinit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/AlifElectronicQueue/internal/pkg/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataAccess struct {
	ConfigFilePath string
	ConVar         *sqlx.DB
	DriverName     string
}

//
func (db *DataAccess) InitDataBase() error {
	var dbData types.DBData
	slbyte, _ := ioutil.ReadFile(db.ConfigFilePath)

	err := json.Unmarshal(slbyte, &dbData) // decoding JSON data into the struct
	if err != nil {
		return errors.New("Can not decode from JSON file!")
	}
	//dbInfo := dbData.DataBaseUser + `://` + dbData.DataBaseUser + `:` + dbData.DataBasePassword + `@127.0.0.1/` + dbData.DataBaseName + " sslmode=disable"
	dn := "user=postgres password=tomatotime1 dbname=alifqueue sslmode=disable"
	db.ConVar = sqlx.MustConnect("postgres", dn)

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
		ConfigFilePath: "./configs/dbDataAccess.json",
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

// func NewStorageContainer(defaultExpiration, cleanupInterval time.Duration) *Cache {

//     // инициализируем карту(map) в паре ключ(string)/значение(Item)
//     items := make(map[string]Item)

//     cache := Cache{
//         items:             items,
//         defaultExpiration: defaultExpiration,
//         cleanupInterval:   cleanupInterval,
//     }

//     // Если интервал очистки больше 0, запускаем GC (удаление устаревших элементов)
//     if cleanupInterval > 0 {
//         cache.StartGC() // данный метод рассматривается ниже
//     }

//     return &cache
// }
