package analytics

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type AnalyticsControllers struct {
	srv *AnalyticsService
}

func InitControllers(asrv *AnalyticsService) *AnalyticsControllers {
	return &AnalyticsControllers{
		srv: asrv,
	}
}

/*func (c *AnalyticsControllers) GetMultyHighCashback() http.HandlerFunc {
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

		slice, errr = c.srv.GetMostHihgCashBack(service.Count)
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
}*/

func (cSrv *AnalyticsControllers) UsersActivity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			a types.Answer
		)

		users, err := cSrv.srv.repo.UsersActivity()
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}

		a.Code = http.StatusOK
		a.Message = http.StatusText(a.Code)
		a.Info = nil
		json.NewEncoder(w).Encode(users)
	}
}
func (c *AnalyticsControllers) ServiceAllPurchases() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.ServiceProviderAllPurchase
			a   types.Answer
			pr  types.ServicesProviderDateRequest
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
		s, err = c.srv.repo.ServiceAllPurchase(&pr)
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

func (cSrv *AnalyticsControllers) ShowServiceProvidersRating() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			a                 types.Answer
			city              string
			err               error
			service_providers []types.ServiceProviderRating
		)

		city = r.URL.Query().Get("city")
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}

		service_providers, err = cSrv.srv.repo.ServiceProviderRating(city)
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}

		a.Code = http.StatusOK
		a.Message = http.StatusText(a.Code)
		a.Info = nil
		json.NewEncoder(w).Encode(service_providers)
	}
}

func (c *AnalyticsControllers) HandleGetServicePeroviderAnnualReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var (
			period              types.Period
			newServiceProviders []types.Row
		)

		err := json.NewDecoder(r.Body).Decode(&period)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		newServiceProviders, err = c.srv.repo.GetServiceProviderAnnualReport(period.Start, period.End)

		if err != nil {
			log.Fatal("Function NewUsersByPeriod no returned ")
		}
		err = json.NewEncoder(w).Encode(newServiceProviders)
		if err != nil {
			panic(err)
		}
	}
}
