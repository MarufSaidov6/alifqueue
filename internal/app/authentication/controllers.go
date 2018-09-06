package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/AlifElectronicQueue/internal/pkg/types"
)

type AuthenticationControllers struct {
	srv *AuthenticationService
}

func InitControllers(asrv *AuthenticationService) *AuthenticationControllers {
	return &AuthenticationControllers{
		srv: asrv,
	}
}

//var
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// TODO Authentication for Admin
func (c *AuthenticationControllers) AdminSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var login types.AdminAuth

		json.NewDecoder(r.Body).Decode(&login)
		answer := c.srv.TestLogin(login)

		if answer {
			json.NewEncoder(w).Encode("*Redirect") ///TODO: Set session cookie
		} else {
			w.Write([]byte("Неверный пароль или логин"))
		}
	}
}
