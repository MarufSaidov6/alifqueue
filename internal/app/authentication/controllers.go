package authentication

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/AlifElectronicQueue/internal/pkg/types"
	// _ "github.com/AlifElectronicQueue/web/template"
	"github.com/gorilla/sessions"
)

type AuthenticationControllers struct {
	srv *AuthenticationService
}

func InitControllers(asrv *AuthenticationService) *AuthenticationControllers {
	return &AuthenticationControllers{
		srv: asrv,
	}
}

//var emailVal = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?)*$")

var templates = template.Must(template.ParseFiles(
	"C:/Projects/Go/src/github.com/AlifElectronicQueue/web/template/login.html",
	"C:/Projects/Go/src/github.com/AlifElectronicQueue/web/template/admin.html",
))

// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func init() {

	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 24 * 30, // 1 mounth
		HttpOnly: true,
	}
}

func (c *AuthenticationControllers) Application() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO:Show list of orders
		var data types.UserAuth
		err := json.NewDecoder(r.Body).Decode(&data)
		err = c.srv.CreateUser(data)
		if err != nil {
			//		panic(err)
		}

	}
}

// func initSession(r *http.Request) *sessions.Session {
// 	session, err := store.Get(r, "session")
// 	if err != nil {
// 		fmt.Println("Check initsession")
// 	}
// 	if session.IsNew {
// 		session.Options.Domain = "localhost"
// 		session.Options.HttpOnly = false
// 		session.Options.Secure = true
// 	}
// 	return session
// }
func (c *AuthenticationControllers) SelectUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")

		//templates.ExecuteTemplate(w, "admin.html", "newusers")
		session, err := store.Get(r, "session")
		fmt.Println("store", store)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		users, err := c.srv.repo.GetPersons()
		fmt.Println("check,", users)
		if err != nil {
			fmt.Println("Smth wrong in conntr->", err)
		}
		err = templates.ExecuteTemplate(w, "admin.html", users)
		if err != nil {
			fmt.Println(err)
		}
		//err = json.NewEncoder(w).Encode(users)
		// if err != nil {
		// 	panic(err)
		// }
	}
}

// func (c *AuthenticationControllers) Update() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		var (
// 			user   types.GetUsers
// 			answer types.Answer
// 		)
// 		if r.Method != http.MethodPut {
// 			answer.Code = http.StatusMethodNotAllowed
// 			answer.Message = http.StatusText(answer.Code)
// 			answer.Info = nil
// 			json.NewEncoder(w).Encode(answer)
// 			return
// 		}
// 		err := json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			answer.Code = http.StatusBadRequest
// 			answer.Message = http.StatusText(answer.Code)
// 			answer.Info = nil
// 			json.NewEncoder(w).Encode(answer)
// 			return
// 		}
// 		err := c.srv.UpdateUserStatus()
// 		if err != nil {
// 			answer.Code = http.StatusInternalServerError
// 			answer.Message = http.StatusText(answer.Code)
// 			answer.Info = nil
// 			json.NewEncoder(w).Encode(answer)
// 			return
// 		}
// 		str := strconv.FormatInt(count, 10)
// 		answer.Code = http.StatusOK
// 		answer.Message = http.StatusText(answer.Code)
// 		answer.Info = str
// 		json.NewEncoder(w).Encode(str)
// 	}
// }

func (c *AuthenticationControllers) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "text/html")

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "C:/Projects/Go/src/github.com/AlifElectronicQueue/web/template/login.html") //!
		case "POST":

			//********AUTHENTICATION PROCESS*********//

			var login types.AdminAuth
			login.Login = r.FormValue("username")
			login.PasswordHash = r.FormValue("password")
			answer := c.srv.Authenticate(login)

			if answer {
				//*****AUTHENTICATED******//

				//*TODO:SET SESSION
				session, err := store.Get(r, "session")
				if err != nil {
					fmt.Println("Check initsession")
				}

				//*TODO:SET SESSION STATUS
				session.Values["authenticated"] = true //Authenticated

				//*TODO:SET SessionUUID for specific user
				// sessionToken, _ := uuid.NewV4() //Generate session token
				// session.Values["userid"] = sessionToken.String()
				//session.Options.MaxAge =

				//w.Write([]byte(session.Values["userid"]))

				err = session.Save(r, w)
				if err != nil {
					fmt.Println("problem save")
					http.Error(w, err.Error(), 500)
					return
				}
				//w.Write([]byte(session.Values["userid"]))

				//Redirect
				http.Redirect(w, r, "/admin/applications", 302)

			} else {
				http.Redirect(w, r, "/login", 302)

			}
		}
	}
}

func (c *AuthenticationControllers) AdminLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, "session failed", http.StatusInternalServerError)
		}
		//Only Post Method

		// Revoke users authentication
		session.Values["authenticated"] = false
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
}
