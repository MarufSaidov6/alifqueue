package authentication

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/AlifElectronicQueue/internal/pkg/types"
	// _ "github.com/AlifElectronicQueue/web/template"
	"github.com/gorilla/mux"
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
	"./web/template/login.html",
	"./web/template/admin.html",
	"./web/template/update.html",
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

func (c *AuthenticationControllers) OrderedApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		order, _ := strconv.Atoi(params["order"])
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

		users, err := c.srv.repo.GetPersonsOrdered(order)
		if err != nil {
			fmt.Println("Smth wrong in conntr->", err)
		}
		fmt.Println("ttt", users)
		err = templates.ExecuteTemplate(w, "admin.html", users)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func (c *AuthenticationControllers) SelectUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

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

		user, err := c.srv.repo.GetPersonById(id)
		if err != nil {
			fmt.Println("Smth wrong in conntr->", err)
		}
		fmt.Println(user, "here")
		err = templates.ExecuteTemplate(w, "update.html", user)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func (c *AuthenticationControllers) UpdateApplicationById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		status := r.PostFormValue("my_field")

		session, err := store.Get(r, "session")
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		err = c.srv.repo.UpdateApplicationStatusById(status, id)
		if err != nil {
			fmt.Println("Update Error:", err)
		}
		http.Redirect(w, r, "/admin/applications", 302)

	}
}

func (c *AuthenticationControllers) SelectUserByContact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		contact := params["contact"]
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
		fmt.Println("Whh")
		user, err := c.srv.repo.GetPersonByContact(contact)
		fmt.Println("check ID,", user)
		if err != nil {
			fmt.Println("Smth wrong in conntr->", err)
		}
		err = templates.ExecuteTemplate(w, "admin.html", user)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func (c *AuthenticationControllers) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./web/template/login.html") //!
		case "POST":

			//********AUTHENTICATION PROCESS*********//

			var login types.AdminAuth
			login.Login = r.FormValue("username")
			login.PasswordHash = r.FormValue("password")

			auth := c.srv.Authenticate(login)
			if auth {
				//*TODO:SET SESSION
				session, err := store.Get(r, "session")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
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
					http.Error(w, err.Error(), 500)
					return
				}

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

		// Revoke users authentication
		session.Values["authenticated"] = false
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
}
