package databaseinit

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository struct {
	DB *sqlx.DB
}

//!WORKS
// checking  login for existing
func (ar *AuthenticationRepository) VerifyLogin(login string) (result bool) {
	Query := `select exists(select login from admin where login=$1)`
	ar.DB.QueryRowx(Query, login).Scan(&result)
	return result
}

func (ar *AuthenticationRepository) VerifyPasswordHash(login string, password string) bool {

	//*STEP1:GENERATE HASH&SALT FROM GIVEN PASSWORD
	// var hashPassword string
	// h := md5.New()
	// h.Write([]byte(password))
	// newHashPassword := fmt.Sprintf("%x", h.Sum(nil))

	//* Generate "hash" to store from user password
	// recenthashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(recenthashPassword), "gen from", password)
	recenthashPassword := password
	//*STEP2:GET HASH&SALT FROM DB
	var DBHashPassword []byte
	Query := `SELECT password FROM admin WHERE login=$1;`
	ar.DB.QueryRowx(Query, login).Scan(&DBHashPassword)
	fmt.Println(string(DBHashPassword))

	//*STEP3:COMPARE HASHES&SALT
	err := bcrypt.CompareHashAndPassword([]byte(DBHashPassword), []byte(recenthashPassword))
	fmt.Println(string(DBHashPassword))
	if err != nil { //?
		return false
	}
	return true
}

// func (ar *AuthenticationRepository) GetHashPassword(login string) (hashPassword string) {
// 	Query := `select password from admin where login='saidov';`
// 	ar.DB.QueryRowx(Query, login).Scan(&hashPassword)
// 	return hashPassword
// }
