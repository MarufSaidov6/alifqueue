package authentication

import (
	"crypto/md5"
	"fmt"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type AuthenticationService struct {
	repo types.IAuthenticationRepository
}

func InitService(aRep types.IAuthenticationRepository) *AuthenticationService {
	return &AuthenticationService{
		repo: aRep,
	}
}

func (srv *AuthenticationService) CheckHashPassword(email, password string) bool {
	h := md5.New()
	h.Write([]byte(password))
	newHashPassword := fmt.Sprintf("%x", h.Sum(nil))

	hashPassword := srv.repo.GetHashPassword(email)
	return newHashPassword == hashPassword
}

func (srv *AuthenticationService) CheckHashPasswordSP(email, password string) bool {
	h := md5.New()
	h.Write([]byte(password))
	newHashPassword := fmt.Sprintf("%x", h.Sum(nil))

	hashPassword := srv.repo.GetHashPasswordSP(email)
	return newHashPassword == hashPassword
}

func (srv *AuthenticationService) TestLogin(ou types.UserProviderAuthentication) (res bool, usInf types.ServiceProviderDeletePl) {
	if !srv.repo.EmailCheckSP(ou.Email) {
		return false, usInf
	}
	if !srv.CheckHashPasswordSP(ou.Email, ou.Password) {
		return false, usInf
	}
	res = true
	usInf = srv.repo.GetUserInfoSP(ou.Email)
	return false, usInf
}
