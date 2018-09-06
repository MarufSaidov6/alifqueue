package authentication

import (
	"crypto/md5"
	"fmt"

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

func (srv *AuthenticationService) CheckHashPassword(login, password string) bool {
	h := md5.New()
	h.Write([]byte(password))
	newHashPassword := fmt.Sprintf("%x", h.Sum(nil))

	hashPassword := srv.repo.GetHashPassword(login)
	return newHashPassword == hashPassword
}

func (srv *AuthenticationService) TestLogin(ou types.AdminAuth) (result bool) {
	if !srv.repo.LoginCheck(ou.Login) {
		return false
	}
	if !srv.CheckHashPassword(ou.Login, ou.Password) {
		return false
	}
	result = true

	return
}
