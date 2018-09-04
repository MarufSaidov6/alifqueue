package users

import (
	"encoding/json"
	"net/http"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)


type UsersControllers struct {
	srv *UsersService
}

func InitControllers(usrv *UsersService) *UsersControllers {
	return &UsersControllers{
		srv: usrv,
	}
}

func (c *UsersControllers) StoryCashback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var (
			slice  []types.StoryCashback
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetStoryCash(entry.Id, entry.Count)
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

func (c *UsersControllers) GetAllActiveServiceProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice  []types.InfoServiceProvider
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetInfoService_Provider(entry.Count, entry.City)
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

func (c *UsersControllers) ShowInfoServices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice  []types.InfoServiceProvider
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetAllServices(entry.Id)
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

func (c *UsersControllers) ShowListInTouchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice  []types.ListInTouchProduct
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetListInTouchProduct(entry.Id, entry.City)
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

func (c *UsersControllers) GetAllActiveDiscounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice  []types.ActiveDiscounts
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetListActiveDiscounts(entry.City, entry.Id)
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

func (c *UsersControllers) GetListNewDiscounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			slice  []types.ActiveDiscounts
			errr   error
			answer types.Answer
			entry  types.EntryServiceProvider
		)
		if r.Method != http.MethodPost {
			answer.Code = http.StatusMethodNotAllowed
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			answer.Code = http.StatusBadRequest
			answer.Message = http.StatusText(answer.Code)
			answer.Info = nil
			json.NewEncoder(w).Encode(answer)
			return
		}

		slice, errr = c.srv.GetNewDiscounts(entry.Id, entry.City)
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
