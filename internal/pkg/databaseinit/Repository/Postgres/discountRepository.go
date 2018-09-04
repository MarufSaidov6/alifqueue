package databaseinit

import (
	"errors"
	"time"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type DiscountRepository struct {
	DB *sqlx.DB
}

func (repo *DiscountRepository) CreateDiscount(dis types.Discount) error {
	var (
		query = `INSERT INTO discounts(name, discount_percentage, service_provider_id, discount_state, start_date, end_date) 
			VALUES($1, $2, $3, $4, to_date($5, 'dd:mm:yyyy'), to_date($6, 'dd:mm:yyyy'))`
		err error
	)
	_, err = repo.DB.Exec(query, dis.Name, dis.DiscountPercentage, dis.ServiceProviderID, dis.DiscountState, dis.StartDate, dis.EndDate)

	return err
}

func (repo *DiscountRepository) DiscountExist(name string, serviceProviderID int) bool {
	var (
		query = `SELECT EXISTS(SELECT 1 FROM discounts WHERE name = $1 AND service_provider_id = $2)`
		exist bool
		row   *sqlx.Row
	)
	row = repo.DB.QueryRowx(query, name, serviceProviderID)
	row.Scan(&exist)

	return exist
}

func (repo *DiscountRepository) DiscountsList() ([]types.Discount, error) {
	var (
		query     = `SELECT * FROM discounts`
		discounts []types.Discount
		err       error
	)
	err = repo.DB.Select(&discounts, query)
	if err != nil {
		return nil, err
	}
	return discounts, err
}

func (repo *DiscountRepository) DiscountsActiveList() ([]types.Discount, error) {
	var (
		query     = `SELECT * FROM discounts WHERE discount_state = 'active'`
		discounts []types.Discount
		err       error
	)
	err = repo.DB.Select(&discounts, query)
	if err != nil {
		return nil, err
	}
	return discounts, err
}

func (repo *DiscountRepository) DiscountsSoonList() ([]types.Discount, error) {
	var (
		query     = `SELECT * FROM discounts WHERE discount_state = 'soon'`
		discounts []types.Discount
		err       error
	)
	err = repo.DB.Select(&discounts, query)
	if err != nil {
		return nil, err
	}
	return discounts, err
}

func (repo *DiscountRepository) DiscountsPastList(pastDis types.PastDiscount) ([]types.Discount, error) {
	var (
		query     = `SELECT * FROM discounts WHERE discount_state = 'past' AND start_date >= $1 AND end_date <= $2`
		discounts []types.Discount
		err       error
	)
	pastDis.FromDate += "T00:00:00.000Z"
	pastDis.ToDate += "T00:00:00.000Z"
	err = repo.ParseDate(pastDis)
	if err != nil {
		return nil, errors.New("INVALID DATES!")
	}
	fromdate, _ := time.Parse(time.RFC3339, pastDis.FromDate)
	todate, _ := time.Parse(time.RFC3339, pastDis.ToDate)
	err = repo.DB.Select(&discounts, query, fromdate, todate)
	if err != nil {
		return nil, err
	}
	return discounts, err
}

func (repo *DiscountRepository) ParseDate(pastDis types.PastDiscount) (err error) {
	_, err = time.Parse(time.RFC3339, pastDis.FromDate)
	if err != nil {
		return
	}
	_, err = time.Parse(time.RFC3339, pastDis.ToDate)
	if err != nil {
		return
	}
	return nil
}
