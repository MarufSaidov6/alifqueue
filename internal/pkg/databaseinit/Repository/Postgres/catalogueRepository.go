package databaseinit

import (
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type CatalogueRepository struct {
	DB *sqlx.DB
}

// to get account & username, email ..
func (cr *CatalogueRepository) GetUserInfo(email string) types.UserAuthenticationAnswer {
	var (
		id    int
		inq   = `select id, name, surname, email from users where email=$1`
		usInf types.UserAuthenticationAnswer
	)
	usInf.Email = email
	cr.DB.QueryRow(inq, usInf.Email).Scan(&id, &usInf.Name, &usInf.Surname, &usInf.Email)

	inq = `select * from accounts where owner_id=$1 and account_type='bonus'`
	cr.DB.QueryRowx(inq, id).StructScan(&usInf.AccountBonus)
	inq = `select * from accounts where owner_id=$1 and account_type='cash'`
	cr.DB.QueryRowx(inq, id).StructScan(&usInf.AccountCash)
	return usInf

}

// Iserts New User
func (cr *CatalogueRepository) InsertMyUser(us *types.NewUsersTable) {
	var (
		id  = 0
		inq = `insert into users(name, surname, adress, registration_date, email, password_hash, 
		is_deleted, phone_number, city, country) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id`
		tx = cr.DB.MustBegin()
	)
	tx.QueryRowx(inq, us.Name, us.Surname, us.Adress, us.RegDate, us.Email, us.HashPassword, us.IsDeleted,
		us.PhoneNumber, us.City, us.Country).Scan(&id)

	inq = `insert into accounts values(default, 0, $1, 'cash', 'user')`
	tx.MustExec(inq, &id)
	inq = `insert into accounts values(default, 0, $1, 'bonus', 'user')`
	tx.MustExec(inq, &id)

	tx.Commit()
}

// check new email in database
// id email exists returns true
func (cotr *CatalogueRepository) ExistsUser(email string) (res bool) {
	inq := `select exists(select 1 from users where email=$1)`
	cotr.DB.QueryRowx(inq, email).Scan(&res)
	return res
}

func (cr *CatalogueRepository) InsertRegistretedUser(us *types.NewServiceProvider) {
	tx := cr.DB.MustBegin()

	inq := `insert into service_provider(contact_name,adress,registration_date,cashback_percent,
		alifs_cashback_share,users_cashback_share,email,password,company_name,is_deleted,city,
		country) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	tx.MustExec(inq, us.Contact_Name, us.Adress, us.Registration_Date, us.CashBack_Percent,
		us.Alifs_CashBack_Share, us.Users_CashBack_Share, us.Email, us.Password, us.Company_Name,
		us.Is_deleted, us.City, us.Country)
	tx.Commit()

}

func (cotr *CatalogueRepository) Exists(email string) (res bool, err error) {
	inq := `select exists(select 1 from service_provider where email=$1)`
	err = cotr.DB.QueryRowx(inq, email).Scan(&res)
	return res, err
}

func (repo *CatalogueRepository) UpdateServiceProvider(id int, contactName, adress, registrationDate string, cashBackPercent, alifsCashBackShare, usersCashBackShare int, email, password, companyName, city, country string) (count int64, err error) {
	Qupdate := `UPDATE service_provider SET contact_name=$2, adress=$3, registration_date=$4, cashback_percent=$5, alifs_cashback_share=$6, users_cashback_share=$7, email=$8, password=$9, company_name=$10, city=$11, country=$12 WHERE id=$1;`
	resl, err := repo.DB.Exec(Qupdate, id, contactName, adress, registrationDate, cashBackPercent, alifsCashBackShare, usersCashBackShare, email, password, companyName, city, country)
	if err != nil {
		panic(err)
	}
	count, err = resl.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, err
}

func (repo *CatalogueRepository) GetMultiServiceProvider(count int) ([]types.NewServiceProvider, error) {
	var (
		Qselect = `SELECT * FROM service_provider LIMIT $1;`
		slice   []types.NewServiceProvider
		service types.NewServiceProvider
	)

	rows, err := repo.DB.Queryx(Qselect, count)
	if err != nil {
		panic(err)
	}
	for rows.Next() {

		err := rows.StructScan(&service)
		if err != nil {
			panic(err)
		}
		slice = append(slice, service)

	}
	return slice, err
}

func (repo *CatalogueRepository) GetSingleServiceProvider(id int) (types.NewServiceProvider, error) {
	var (
		QselectById = `SELECT *FROM service_provider WHERE id = $1;`
		service     types.NewServiceProvider
		err         error
	)
	err = repo.DB.QueryRowx(QselectById, id).StructScan(&service)
	return service, err
}

func (repo *CatalogueRepository) DeleteServiceProvider(id int, isdeleted bool) (count int64, err error) {

	Qupdate := `UPDATE service_provider SET is_deleted = $2 WHERE id = $1;`
	resl, err := repo.DB.Exec(Qupdate, id, isdeleted)
	if err != nil {
		panic(err)
	}
	count, err = resl.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, err
}
