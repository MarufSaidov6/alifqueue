package databaseinit

import (
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository struct {
	DB *sqlx.DB
}

// checking  email for existing
func (ar *AuthenticationRepository) EmailCheck(email string) (res bool) {
	inq := `select exists(select 1 from users where email=$1)`
	ar.DB.QueryRowx(inq, email).Scan(&res)
	return res
}

func (ar *AuthenticationRepository) EmailCheckSP(email string) (res bool) {
	inq := `select exists(select 1 from service_provider where email=$1)`
	res = false
	ar.DB.QueryRowx(inq, email).Scan(&res)
	return res
}

func (ar *AuthenticationRepository) GetHashPassword(email string) (hashPassword string) {
	inq := `select password_hash from users where email=$1`
	ar.DB.QueryRowx(inq, email).Scan(&hashPassword)
	return
}

func (ar *AuthenticationRepository) GetHashPasswordSP(email string) (hashPassword string) {
	inq := `select password_hash from service_provider where email=$1`
	ar.DB.QueryRowx(inq, email).Scan(&hashPassword)
	return
}

// getting account & username email ..
func (cr *AuthenticationRepository) GetUserInfo(email string) *types.UserAuthenticationAnswer {
	var (
		usInf types.UserAuthenticationAnswer
		id    int
	)
	inq := `select id,name,surname,email from users where email=$1`
	cr.DB.QueryRowx(inq, email).Scan(&id, &usInf.Name, &usInf.Surname, &usInf.Email)
	inq = `select * from accounts where owner_id=$1 and account_type='bonus'`
	cr.DB.QueryRowx(inq, id).StructScan(&usInf.AccountBonus)
	inq = `select * from accounts where owner_id=$1 and account_type='cash'`
	cr.DB.QueryRowx(inq, id).StructScan(&usInf.AccountCash)
	return &usInf

}

func (cr *AuthenticationRepository) GetUserInfoSP(email string) (usInf types.ServiceProviderDeletePl) {
	inq := `select * from service_provider where email=$1`
	cr.DB.QueryRowx(inq, email).StructScan(&usInf)
	return usInf
}

func (ar *AuthenticationRepository) SelectAll() []string {
	return nil
}
