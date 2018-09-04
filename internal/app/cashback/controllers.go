package cashback

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type CashbackControllers struct {
	srv *CashbackService
}

func InitControllers(asrv *CashbackService) *CashbackControllers {
	return &CashbackControllers{
		srv: asrv,
	}
}

// TODO: when want change bonus to money
// POST method {"Id", "Bonus"}
// will get ur accounts bonus and money from DB by Id
// If there enough bonu in ur account will exchange
// Ur Bonus to Money and upgrades ur account in DB
func (c *CashbackControllers) BonusExchange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			User types.UserAccount
		)
		err := json.NewDecoder(r.Body).Decode(&User)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		c.srv.DoExchange(User.UserId, User.UserBonus)
		json.NewEncoder(w).Encode(User)
	}
}

func (c *CashbackControllers) GetTransactionListPakupki2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.TransactionPakupki
			a   types.Answer
			pr  types.UserTransactionRequest
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
		s, err = c.srv.repo.GetTransactionListPakupki(&pr)
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

func (c *CashbackControllers) GetTransactionListPerevodi2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.TransactionPerevodi
			a   types.Answer
			pr  types.UserTransactionRequest
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
		s, err = c.srv.repo.GetTransactionListPerevodi(&pr)
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

func (c *CashbackControllers) TranInfoUserPakupki2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.TranUserInfoPakupki
			a   types.Answer
			pr  types.UserTranInfoRequest
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
		s, err = c.srv.repo.TranInfoUserPakupki(&pr)
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

func (c *CashbackControllers) TranInfoUserPerevodi2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			err error
			s   []types.TranUserInfoPerevodi
			a   types.Answer
			pr  types.UserTranInfoRequest
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
		s, err = c.srv.repo.TranInfoUserPerevodi(&pr)
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

// функция для перевода денег и баллов
// ждем tr инстанса BalanceTransfer которое должно иметь:
// Balance - размер перевода
// SenderAccount - аккаунт отправителя
// ReceiverAccount - аккаунт получателя
func (c *CashbackControllers) CashBounusTransfer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			a   types.Answer
			tr  types.BalanceTransfer
			err error
		)
		json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
		}
		err, tr = c.srv.DoTransfer(tr)
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = err
			json.NewEncoder(w).Encode(a)
		}
		json.NewEncoder(w).Encode(tr)
	}
}

func (cSrv *CashbackControllers) PurchaseByCash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			jq types.ForPurchaseByCash
			a  types.Answer
		)

		if r.Method != http.MethodPut {
			a.Code = http.StatusMethodNotAllowed
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			fmt.Println("Wrong method")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&jq)
		fmt.Println(jq)
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}

		fmt.Println("Got inside controller, json values id: ", jq)

		err = cSrv.srv.PurchaseByCash(jq.UserId, jq.ServiseId)
		if err != nil {
			if err == NotEnoughAmountErr {
				a.Code = http.StatusInternalServerError
				a.Message = "Not enough money"
				a.Info = nil
				json.NewEncoder(w).Encode(a)
			} else {
				a.Code = http.StatusInternalServerError
				a.Message = http.StatusText(a.Code)
				a.Info = nil
				json.NewEncoder(w).Encode(a)
				fmt.Println("Error inside  PurchaseByCash controller: ", err)
				return
			}
		} else {
			a.Code = http.StatusOK
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
		}
	}
}

func (cSrv *CashbackControllers) PurchaseByBonus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var (
			jq types.ForPurchaseByCash
			a  types.Answer
		)

		if r.Method != http.MethodPut {
			a.Code = http.StatusMethodNotAllowed
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			fmt.Println("Wrong method")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&jq)
		fmt.Println(jq)
		if err != nil {
			a.Code = http.StatusBadRequest
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
			return
		}

		fmt.Println("Got inside controller, json values id: ", jq)

		err = cSrv.srv.PurchaseByBonus(jq.UserId, jq.ServiseId)
		if err != nil {
			if err == NotEnoughAmountErr {
				a.Code = http.StatusInternalServerError
				a.Message = "Not enough money"
				a.Info = nil
				json.NewEncoder(w).Encode(a)
			} else {
				a.Code = http.StatusInternalServerError
				a.Message = http.StatusText(a.Code)
				a.Info = nil
				json.NewEncoder(w).Encode(a)
				fmt.Println("Error inside  PurchaseByBonus controller: ", err)
				return
			}
		} else {
			a.Code = http.StatusOK
			a.Message = http.StatusText(a.Code)
			a.Info = nil
			json.NewEncoder(w).Encode(a)
		}
	}
}
