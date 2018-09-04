package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type AuthenticationControllers struct {
	srv *AuthenticationService
}

func InitControllers(asrv *AuthenticationService) *AuthenticationControllers {
	return &AuthenticationControllers{
		srv: asrv,
	}
}

// TODO Authentication for User
func (c *AuthenticationControllers) UserSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var login types.UserProviderAuthentication
		json.NewDecoder(r.Body).Decode(&login)
		ans, usInfo := c.srv.TestLogin(login)

		if ans {
			json.NewEncoder(w).Encode(usInfo)
		} else {
			w.Write([]byte("ERROR:Wrong email or Password"))
		}
	}
}

// TODO Authentication for ServiceProvider
func (c *AuthenticationControllers) ServiceProviderSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var login types.UserProviderAuthentication

		json.NewDecoder(r.Body).Decode(&login)
		ans, usInfo := c.srv.TestLogin(login)
		if ans == true {
			json.NewEncoder(w).Encode(usInfo)
		} else {
			w.Write([]byte("ERROR:Wrong email or Password"))
		}
	}
}
