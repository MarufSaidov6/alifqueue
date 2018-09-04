package discount

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	log "github.com/sirupsen/logrus"
)

type DiscountControllers struct {
	srv *DiscountService
}

func InitControllers(asrv *DiscountService) *DiscountControllers {
	return &DiscountControllers{
		srv: asrv,
	}
}

func (c *DiscountControllers) CreateDiscount() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		dis := types.Discount{}
		json.NewDecoder(request.Body).Decode(&dis)

		fmt.Println(dis)
		err := c.srv.CreateDiscount(dis)
		if err != nil {
			log.Info("NO CREATED NEW DISCOUNT: " + err.Error())
			fmt.Fprintln(response, "NO CREATED NEW DISCOUNT! \nPLEASE TRY AGAIN")
		}
	}
}

func (c *DiscountControllers) DiscountsList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		discounts, err := c.srv.DiscountsList()

		if err != nil {
			fmt.Fprintln(response, "CANNOT SELECT DISCONTS!")
		}
		json.NewEncoder(response).Encode(discounts)
	}
}

func (c *DiscountControllers) DiscountsActiveList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		discounts, err := c.srv.DiscountsActiveList()

		if err != nil {
			fmt.Fprintln(response, "CANNOT SELECT DISCONTS!")
		}
		json.NewEncoder(response).Encode(discounts)
	}
}

func (c *DiscountControllers) DiscountsSoonList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		discounts, err := c.srv.DiscountsSoonList()

		if err != nil {
			fmt.Fprintln(response, "CANNOT SELECT DISCONTS!")
		}
		json.NewEncoder(response).Encode(discounts)
	}
}

func (c *DiscountControllers) DiscountsPastList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		var pastDis types.PastDiscount
		json.NewDecoder(request.Body).Decode(&pastDis)

		discounts, err := c.srv.DiscountsPastList(pastDis)

		if err != nil {
			fmt.Fprintln(response, "CANNOT SELECT DISCONTS!", err)
		}
		json.NewEncoder(response).Encode(discounts)
	}
}
