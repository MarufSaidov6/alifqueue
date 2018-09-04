package databaseinit

import (
	"fmt"
	"time"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	DB *sqlx.DB
}

func (repo *UsersRepository) GetStoryCash(id, count int) ([]types.StoryCashback, error) {
	var (
		Qselect = `SELECT users.id,users.name,users.surname,transactions.date,transactions.transaction_type,cashbacks.amount 
		FROM users  LEFT OUTER JOIN transactions ON (users.id = transactions.users_id) LEFT OUTER 
		JOIN cashbacks ON(transactions.id = transaction_id)  WHERE users.id =$1 and cashbacks.amount>0 
		ORDER BY transactions.date DESC LIMIT $2;`
		slice []types.StoryCashback
		story types.StoryCashback
	)

	rows, err := repo.DB.Queryx(Qselect, id, count)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&story)
		if err != nil {
			panic(err)
		}
		slice = append(slice, story)

	}
	return slice, err
}

func (repo *UsersRepository) GetInfoService_Provider(count int, city string) ([]types.InfoServiceProvider, error) {
	var (
		Qselect = `SELECT service_provider.id,service_provider.company_name,service_provider.adress,service_provider.city 
		from service_provider where is_deleted = false and city = $2 LIMIT $1;`
		slice []types.InfoServiceProvider
		info  types.InfoServiceProvider
	)

	rows, err := repo.DB.Queryx(Qselect, count, city)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&info)
		if err != nil {
			panic(err)
		}
		slice = append(slice, info)

	}
	return slice, err
}

func (repo *UsersRepository) GetAllServices(id int) ([]types.InfoServiceProvider, error) {
	var (
		Qselect = `SELECT service_provider.id,service_provider.company_name,service_provider.users_cashback_share,services.name_of_service,services.price,services.category 
		FROM service_provider LEFT OUTER JOIN services ON (service_provider.id=services.service_provider_id) WHERE service_provider.id=$1;`
		slice []types.InfoServiceProvider
		info  types.InfoServiceProvider
	)

	rows, err := repo.DB.Queryx(Qselect, id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&info)
		if err != nil {
			panic(err)
		}
		slice = append(slice, info)

	}
	return slice, err
}

func (repo *UsersRepository) GetListInTouchProduct(id int, city string) ([]types.ListInTouchProduct, error) {
	var (
		Qselect = `SELECT users.id,service_provider.company_name,service_provider.city,products.name_of_product,products.category,products.price,accounts.balance
		FROM accounts LEFT OUTER JOIN users ON (accounts.owner_id = users.id)
		LEFT OUTER JOIN service_provider ON (users.city = service_provider.city)
		LEFT OUTER JOIN products ON (service_provider.id= products.service_provider_id)
		WHERE products.price <= accounts.balance and users.id= $1 and users.city=$2;`
		slice []types.ListInTouchProduct
		list  types.ListInTouchProduct
	)

	rows, err := repo.DB.Queryx(Qselect, id, city)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&list)
		if err != nil {
			panic(err)
		}
		slice = append(slice, list)

	}
	return slice, err
}

func (repo *UsersRepository) GetListActiveDiscounts(city string, id int) ([]types.ActiveDiscounts, error) {
	var (
		Qselect = `SELECT service_provider.company_name,service_provider.city,discounts.name_of_discount,discounts.discount_percentage,discounts.discount_state,discounts.start_date,discounts.end_date
		FROM users LEFT OUTER JOIN service_provider ON (users.city=service_provider.city)
		LEFT OUTER JOIN discounts ON (service_provider.id=discounts.service_provider_id)
		WHERE service_provider.is_deleted =false and discounts.discount_state='active' and users.city =$1 and users.id =$2;`
		slice  []types.ActiveDiscounts
		active types.ActiveDiscounts
	)

	rows, err := repo.DB.Queryx(Qselect, city, id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&active)
		if err != nil {
			panic(err)
		}
		slice = append(slice, active)

	}
	return slice, err
}

func (repo *UsersRepository) GetNewDiscounts(id int, city string) ([]types.ActiveDiscounts, error) {
	var (
		Qselect = `SELECT service_provider.company_name,service_provider.city,discounts.name_of_discount,discounts.discount_percentage,discounts.discount_state,discounts.start_date,discounts.end_date
		FROM users LEFT OUTER JOIN service_provider ON (users.city=service_provider.city)
		LEFT OUTER JOIN discounts ON (service_provider.id=discounts.service_provider_id)
		WHERE service_provider.is_deleted = false  and users.city =$1 and users.id =$2 and discounts.start_date between $3 and $4;`
		slice  []types.ActiveDiscounts
		active types.ActiveDiscounts
		time1  = time.Now()
		time2  = time.Now().AddDate(1, 0, 0)
	)
	fmt.Println(time1, time2)
	rows, err := repo.DB.Queryx(Qselect, city, id, time1, time2)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&active)
		if err != nil {
			panic(err)
		}
		slice = append(slice, active)

	}
	return slice, err

}
