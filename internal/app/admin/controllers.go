package admin

import (
	"encoding/json"
	"net/http"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type AdminControllers struct {
	srv *AdminService
}

func InitControllers(adSrv *AdminService) *AdminControllers {
	return &AdminControllers{
		srv: adSrv,
	}
}

func (c *AdminControllers) GetUserInfo2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.UserInfo
			a   types.Answer
			pr  types.UserName
		)
		err = json.NewDecoder(r.Body).Decode(&pr)
		if err != nil {
			return
		}
		if r.Method != http.MethodPost {
			a.Code = http.StatusMethodNotAllowed
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}
		s, err = c.srv.repo.GetUserInfo(&pr)
		if err != nil {
			a.Code = http.StatusInternalServerError
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}
		a.Code = http.StatusOK
		a.Message = http.StatusText(a.Code)
		a.Info = s
		json.NewEncoder(w).Encode(a)
	}
}
