package catalogue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type CatalogueControllers struct {
	srv *CatalogueService
}

func InitControllers(asrv *CatalogueService) *CatalogueControllers {
	return &CatalogueControllers{
		srv: asrv,
	}
}

func (c *CatalogueControllers) SelectAll() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

	}
}

// JSON:Firstname,Lastname,Adress,RegDate,Email,HashPassword,IsDeleted,PhoneNumber,City,Country
func (c *CatalogueControllers) UserSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newUser types.NewUsersTable
		json.NewDecoder(r.Body).Decode(&newUser)
		err, us := c.srv.CreateMyUser(&newUser)
		if err == nil {
			json.NewEncoder(w).Encode(us)
		} else {
			w.Write([]byte(err.Error()))
		}
	}
}

func (c *CatalogueControllers) ServiceProviderSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var (
			answer types.Answer
			us     types.NewServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		json.NewDecoder(r.Body).Decode(&us)
		fmt.Println("afas", us)
		err := c.srv.CreateUser(&us)
		fmt.Println(err)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		answer.Code = http.StatusOK
		answer.Message = http.StatusText(answer.Code)
		answer.Info = us
		json.NewEncoder(w).Encode(answer)

	}
}

func (c *CatalogueControllers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			service types.NewServiceProvider
			answer  types.Answer
		)
		if r.Method != http.MethodPut {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		count, err := c.srv.UpdateServiceProvider(service.ID, service.Contact_Name, service.Adress, service.Registration_Date, service.CashBack_Percent, service.Alifs_CashBack_Share, service.Users_CashBack_Share, service.Email, service.Password, service.Company_Name, service.City, service.Country)
		if err != nil {
			answer.Code = http.StatusInternalServerError
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		str := strconv.FormatInt(count, 10)
		answer.Code = http.StatusOK
		answer.Message = http.StatusText(answer.Code)
		answer.Info = str
		json.NewEncoder(w).Encode(str)
	}
}

func (c *CatalogueControllers) GetAllServiceProvader() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice   []types.NewServiceProvider
			errr    error
			answer  types.Answer
			service types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&service)

		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetMultiServiceProvider(service.Count)
		if errr != nil {
			answer.Code = http.StatusInternalServerError
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		answer.Code = http.StatusOK
		answer.Message = http.StatusText(answer.Code)
		answer.Info = slice

		json.NewEncoder(w).Encode(answer)
	}
}

func (c *CatalogueControllers) GetServiceProviderByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			answer       types.Answer
			entryservice types.EntryServiceProvider
			service      types.NewServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entryservice)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		service, err = c.srv.GetSingleServiceProvider(entryservice.Id)
		if err != nil {
			answer.Code = http.StatusInternalServerError
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		answer.Code = http.StatusOK
		answer.Message = http.StatusText(answer.Code)
		answer.Info = service
		json.NewEncoder(w).Encode(answer)

	}

}

func (c *CatalogueControllers) ServiceProviderDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			answer  types.Answer
			service types.NewServiceProvider
		)

		if r.Method != http.MethodPut {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&service)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		count, err := c.srv.DeleteServiceProvider(service.ID, service.Is_deleted)
		if err != nil {
			answer.Code = http.StatusInternalServerError
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		answer.Code = http.StatusOK
		answer.Message = http.StatusText(answer.Code)
		answer.Info = count
		json.NewEncoder(w).Encode(answer)
	}

}
