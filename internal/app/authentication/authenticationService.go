package authentication

import (
	"strings"
	"time"

	"github.com/AlifElectronicQueue/internal/pkg/types"
)

type AuthenticationService struct {
	repo types.IAuthenticationRepository
}

func InitService(aRep types.IAuthenticationRepository) *AuthenticationService {
	return &AuthenticationService{
		repo: aRep,
	}
}

//
func (srv *AuthenticationService) CreateUser(user types.UserAuth) error {
	//TODO: ADD TIME
	layout := "2006-01-02"
	currentTime := time.Now()
	user.RegistrationDate = currentTime.Format(layout)

	//TODO: ADD PURCHASEDATE
	user.PurchaseDateTime = currentTime.Add(time.Hour * 24 * 3).Format(layout)

	//P12345678
	//A12345678
	//RT23456Y7
	//!VALIDATION PROCESS!!!
	//TODO: "A" may be cyrillic!
	xbytes := []rune(user.SerialNumber)
	xbytes[0] = 'A'
	user.SerialNumber = string(xbytes)
	//TODO: Parse Phone number
	user.Contact = strings.Replace(user.Contact, " ", "", -1)
	//!/

	//TODO:INSERT USER INTO DB
	err := srv.repo.InsertUser(user)
	return err
}

// func (srv *AuthenticationService) UpdateUserStatus(users types.GetUsers) error {
// 	err := srv.repo.UpdateUserStatus(users.Checked, users.Id)
// 	return err
// }

func (srv *AuthenticationService) Authenticate(ou types.AdminAuth) (result bool) {

	if !srv.repo.VerifyLogin(ou.Login) {
		return false
	}
	if !srv.repo.VerifyPasswordHash(ou.Login, ou.PasswordHash) {
		return false
	}

	return true
}
