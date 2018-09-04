package databaseinit

import (
	"database/sql"
	"time"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AnalyticsRepository struct {
	DB *sqlx.DB
}

//
//
// func (repo *AnalyticsRepository) TransactionAmount() error {
// 	// nedodelanniy
// 	var (
// 		query = `SELECT sp.name, sp.segment,
// 				 (SELECT COUNT(*)
// 				 FROM transactions t, service_provider sp
// 				 WHERE t.target_account_id = sp.id) as all_transactions,
// 				 (SELECT ) as average,
// 				 () as minim,
// 				 () as maxim
// 				FROM service_provider sp, `
// 	)

// 	_, err := repo.DB.Queryx(query)

// 	return err
// }

// func (repo *AnalyticsRepository) GetMostHihgCashBack(count int) ([]types.NewServiceProvider, error) {
// 	var (
// 		Qselect = `SELECT * FROM service_provider ORDER BY cashback_percent DESC LIMIT $1 ;`
// 		slice   []types.NewServiceProvider
// 		service types.NewServiceProvider
// 	)
// 	rows, err := repo.DB.Queryx(Qselect, count)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for rows.Next() {
// 		err := rows.StructScan(&service)
// 		if err != nil {
// 			panic(err)
// 		}
// 		slice = append(slice, service)
// 	}

// 	return slice, err
// }
func (repo *AnalyticsRepository) ServiceProviderRating(city string) ([]types.ServiceProviderRating, error) {
	// nedodelanniy
	var (
		service_providers []types.ServiceProviderRating
		query             = `select s1.contact_name, s1.category,  COALESCE(s1.COUNT_OF_CASH_TRANSACTIONS, 0) AS COUNT_OF_CASH_TRANSACTIONS,  COALESCE(s2.COUNT_OF_BONUS_TRANSACTIONS, 0) AS COUNT_OF_BONUS_TRANSACTIONS
		from
			(
			SELECT sp.contact_name, s.category, COUNT (t.amount) AS COUNT_OF_CASH_TRANSACTIONS
			FROM service_provider sp
			JOIN services s ON sp.id = s.service_provider_id
			JOIN transactions t ON sp.id = t.target_account_id
			WHERE sp.city = $1 and t.transaction_type = 'purchase'
			GROUP BY sp.contact_name, s.category
		) s1
		left join
			(
			SELECT sp.contact_name ,COUNT (t.amount) AS COUNT_OF_BONUS_TRANSACTIONS
			FROM service_provider sp
			JOIN services s ON sp.id = s.service_provider_id
			JOIN transactions t ON sp.id = t.target_account_id
			WHERE sp.city = $1 and t.transaction_type = 'purchase_bonus'
			GROUP BY sp.contact_name, s.category
		) s2

		on
			s1.contact_name = s2.contact_name
			GROUP BY s1.COUNT_OF_CASH_TRANSACTIONS ,s1.contact_name, s1.category, s2.COUNT_OF_BONUS_TRANSACTIONS, s2.contact_name
			ORDER BY s1.COUNT_OF_CASH_TRANSACTIONS DESC
		;`
		err  error
		rows *sqlx.Rows
	)

	rows, err = repo.DB.Queryx(query, city)
	if err != nil {
		logrus.Warn("Error: ", err)
		return service_providers, err
	}

	for rows.Next() {
		var service_provider types.ServiceProviderRating
		err := rows.Scan(&service_provider.Contact_Name, &service_provider.Category, &service_provider.Count_of_cash_transactions, &service_provider.Count_of_bonus_transactions)
		if err != nil {
			logrus.Warn("Error: ", err)
			return service_providers, err
		}

		service_providers = append(service_providers, service_provider)
	}

	return service_providers, err
}

func (repo *AnalyticsRepository) ServiceProviderRatingBonus(city string) ([]types.ServiceProviderRating, error) {
	// nedodelanniy
	var (
		service_providers []types.ServiceProviderRating
		query             = `select s1.contact_name, s1.category,  COALESCE(s1.COUNT_OF_CASH_TRANSACTIONS, 0) AS COUNT_OF_CASH_TRANSACTIONS,  COALESCE(s2.COUNT_OF_BONUS_TRANSACTIONS, 0) AS COUNT_OF_BONUS_TRANSACTIONS
		from
			(
			SELECT sp.contact_name, s.category, COUNT (t.amount) AS COUNT_OF_CASH_TRANSACTIONS
			FROM service_provider sp
			JOIN services s ON sp.id = s.service_provider_id
			JOIN transactions t ON sp.id = t.target_account_id
			WHERE sp.city = $1 and t.transaction_type = 'purchase'
			GROUP BY sp.contact_name, s.category
		) s1
		left join
			(
			SELECT sp.contact_name ,COUNT (t.amount) AS COUNT_OF_BONUS_TRANSACTIONS
			FROM service_provider sp
			JOIN services s ON sp.id = s.service_provider_id
			JOIN transactions t ON sp.id = t.target_account_id
			WHERE sp.city = $1 and t.transaction_type = 'purchase_bonus'
			GROUP BY sp.contact_name, s.category
		) s2

		on
			s1.contact_name = s2.contact_name
			GROUP BY s1.COUNT_OF_CASH_TRANSACTIONS ,s1.contact_name, s1.category, s2.COUNT_OF_BONUS_TRANSACTIONS, s2.contact_name
			ORDER BY s2.COUNT_OF_BONUS_TRANSACTIONS
		;`
		err  error
		rows *sqlx.Rows
	)

	rows, err = repo.DB.Queryx(query, city)
	if err != nil {
		logrus.Warn("Error: ", err)
		return service_providers, err
	}

	for rows.Next() {
		var service_provider types.ServiceProviderRating
		err := rows.Scan(&service_provider.Contact_Name, &service_provider.Category, &service_provider.Count_of_cash_transactions, &service_provider.Count_of_bonus_transactions)
		if err != nil {
			logrus.Warn("Error 3: ", err)
			return service_providers, err
		}

		service_providers = append(service_providers, service_provider)
	}

	return service_providers, err
}

func (repo *AnalyticsRepository) UsersActivity() ([]types.UserForAnalitics, error) {
	var (
		query = `SELECT u.name, u.surname, u.city, u.adress, u.registration_date,
		COUNT (t.amount) AS COUNT_OF_TRANSACTIONS, ROUND(AVG (t.amount))AS Average_transaction,
		MAX (t.amount) AS Maximum_transaction
		FROM users u
		INNER JOIN transactions t ON u.id = t.source_account_id
		WHERE t.transaction_type = 'purchase' AND t.date BETWEEN '2018-08-01' AND '2018-08-31'
		GROUP BY u.id
		ORDER BY COUNT_OF_TRANSACTIONS DESC;`
		users []types.UserForAnalitics
		err   error
		rows  *sqlx.Rows
	)

	rows, err = repo.DB.Queryx(query)
	if err != nil {
		logrus.Warn("Error: ", err)
		return users, err
	}

	for rows.Next() {
		var user types.UserForAnalitics
		err := rows.Scan(&user.Name, &user.Surname, &user.City, &user.Adress, &user.Registration_date, &user.Count_of_transactions, &user.Avg_transaction, &user.Max_transaction)
		if err != nil {
			logrus.Warn("Error: ", err)
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}

func (repo *AnalyticsRepository) ServiceAllPurchase(pr *types.ServicesProviderDateRequest) ([]types.ServiceProviderAllPurchase, error) {
	var (
		rows *sql.Rows
		err  error
		ps   []types.ServiceProviderAllPurchase
		p    types.ServiceProviderAllPurchase
	)

	query := `select  sr.id, sr.company_name, sr.cashback_percent,  sum(t.amount), sum(c.amount), 
	count(t.id), $1 as datefrom, $2 as dateto from transactions t
	join accounts a on t.target_account_id = a.id
	join service_provider sr on a.owner_id=sr.id and owner_type='service_provider'
	join cashbacks c on t.id = c.transaction_id
	where t.date between $1 and $2 group by sr.id;`
	rows, err = repo.DB.Query(query, pr.DateFrom, pr.DateTo)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Company_name, &p.Cashback_percent, &p.SumAmount, &p.SumCashback, &p.Count, &p.DateFrom, &p.DateTo)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	return ps, err
}

func (repo *AnalyticsRepository) GetServiceProviderAnnualReport(start, end string) ([]types.Row, error) {

	Begin, err := time.Parse("2006-01-02", start)
	Last, err := time.Parse("2006-01-02", end)
	if err != nil {
		logrus.Error("date formating ERR")
	}

	var (
		rows   *sqlx.Rows
		output []types.Row
		in     types.Row
	)

	Qselect := `
				with s_table as (
					SELECT c.city, gs.month as month,
					COUNT(u.city) as count
					FROM generate_series($1::date, $2::date , interval '1 month') as gs(month)
					CROSS JOIN(select distinct u.city from "service_provider" u) c 
					LEFT JOIN "service_provider" u
					ON date_trunc('month', u.registration_date) = date_trunc('month',gs.month::date) and
					u.city = c.city
					GROUP BY c.city, gs.month
				)
				SELECT city, array_agg(count order by month) as count
				FROM s_table
				GROUP BY s_table.city
				ORDER BY s_table.city;`

	rows, err = repo.DB.Queryx(Qselect, Begin, Last)
	if err != nil {
		logrus.Error("Wrong GetServicePeroviderAnnualReport query ")
	}

	for rows.Next() {
		err = rows.Scan(&in.City, &in.Count)
		if err != nil {
			logrus.Error("Cannot scan rows")
		}

		output = append(output, in)
	}
	return output, err

}
