package authentication

import (
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

func (srv *AuthenticationService) Authenticate(ou types.AdminAuth) (result bool) {

	if !srv.repo.VerifyLogin(ou.Login) {
		return false
	}
	if !srv.repo.VerifyPasswordHash(ou.Login, ou.PasswordHash) {
		return false
	}

	return true
}
