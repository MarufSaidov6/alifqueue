package databaseinit

import (
	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository struct {
	DB *sqlx.DB
}

// checking  login for existing
func (ar *AuthenticationRepository) LoginCheck(login string) (result bool) {
	Query := `select exists(select login from users where login=$1)`
	ar.DB.QueryRowx(Query, login).Scan(&result)
	return
}

func (ar *AuthenticationRepository) GetHashPassword(login string) (hashPassword string) {
	Query := `select password_hash from users where login=$1`
	ar.DB.QueryRowx(Query, login).Scan(&hashPassword)
	return
}
