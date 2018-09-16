package authentication

import (
	"fmt"
	"net/http"
)

type AuthenticationMiddleWares struct {
	mdl *AuthenticationControllers
}

func InitMiddlewares(amdl *AuthenticationControllers) *AuthenticationMiddleWares {
	return &AuthenticationMiddleWares{
		mdl: amdl,
	}
}

//TODO: USER TRAP
func (m *AuthenticationMiddleWares) RequiresLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//!DUDE,IT DOES NOT RETURN VALUES!
		session, _ := store.Get(r, "session") //TODO: Returns a session for the given name after adding it to the registry
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		fmt.Println("Flag middleSESSION!", session) //!
		session.Values["authenticated"] = true
		//*CHECKS IF YOU ARE AUTH
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		//TODO: RESTRICTED FOR USERS
		//TODO: IF COOKIE IS EMPTY

		handler(w, r)
	}
}
