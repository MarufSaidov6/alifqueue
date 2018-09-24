package databaseinit

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/AlifElectronicQueue/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository struct {
	DB *sqlx.DB
}

//!WORKS
// checking  login for existing
func (ar *AuthenticationRepository) VerifyLogin(login string) (result bool) {
	Query := `select exists(select login from admin where login=$1)`
	ar.DB.QueryRowx(Query, login).Scan(&result)
	return result
}

//TODO:INPUT APPLICATION TO DB
func (repo *AuthenticationRepository) InsertUser(user types.UserAuth) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err //
	}
	txId := 0
	query := `insert into application (fullname,contact,serialnumber,registrationdate,purchasedate) values($1,$2,$3,$4,$5) returning id;`
	err = tx.QueryRow(query, user.FullName, user.Contact, user.SerialNumber, user.RegistrationDate, user.PurchaseDateTime).Scan(&txId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `insert into application_service values($1,$2)`

	for _, Services := range user.Services {
		_, err = tx.Exec(query, txId, Services)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (repo *AuthenticationRepository) GetPersons() ([]types.GetUsers, error) {
	var (
		rows   *sql.Rows
		err    error
		output []types.GetUsers
		input  types.GetUsers
	)
	//?Changed ID
	query := "select fullname,contact,serialnumber,purchasedate,checked from application;"
	rows, err = repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
		if err != nil {
			return nil, err
		}
		output = append(output, input)
	}
	fmt.Println(output)
	return output, err
}

func (repo *AuthenticationRepository) GetPersonsOrdered(ordered int) ([]types.GetUsers, error) {
	var (
		rows   *sql.Rows
		err    error
		output []types.GetUsers
		input  types.GetUsers
	)
	fmt.Println(ordered, "lllllllll")
	//?Changed ID

	query := fmt.Sprintf("select fullname,contact,serialnumber,purchasedate,checked from application order by %d", ordered)

	rows, err = repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)
	for rows.Next() {
		err = rows.Scan(&input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
		if err != nil {
			return nil, err
		}
		output = append(output, input)
	}
	fmt.Println("check", output)
	return output, err
}

func (repo *AuthenticationRepository) GetPersonById(id int) ([]types.GetUsers, error) {
	var (
		query  = `SELECT fullname,contact,serialnumber,purchasedate,checked FROM application WHERE id = $1;`
		output []types.GetUsers
		input  types.GetUsers
		err    error
		rows   *sqlx.Rows
	)
	rows, _ = repo.DB.Queryx(query, id)
	for rows.Next() {
		err = rows.Scan(&input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
		if err != nil {
			return nil, err
		}
		output = append(output, input)
	}
	return output, err
}

func (repo *AuthenticationRepository) GetPersonByContact(contact string) ([]types.GetUsers, error) {
	var (
		query  = `SELECT fullname,contact,serialnumber,purchasedate,checked FROM application WHERE contact = $1;`
		output []types.GetUsers
		input  types.GetUsers
		err    error
		rows   *sqlx.Rows
	)
	rows, _ = repo.DB.Queryx(query, contact)
	for rows.Next() {
		err = rows.Scan(&input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
		if err != nil {
			return nil, err
		}
		output = append(output, input)
	}
	fmt.Println("heeeeeee", output)
	return output, err
}

// func (repo *AuthenticationRepository) GetPersonByContact(id int) ([]types.GetUsers, error) {
// 	var (
// 		query  = `SELECT fullname,contact,serialnumber,purchasedate,checked FROM application WHERE contact = $1;`
// 		output []types.GetUsers
// 		input  types.GetUsers
// 		err    error
// 		row    *sql.Row
// 	)
// 	row = repo.DB.QueryRow(query, id)
// 	err = row.Scan(&input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
// 	return output, err
// }

//
// func (repo *AuthenticationRepository) UpdateServiceProvider(status *bool, id int) (err error) {
// 	Qupdate := `UPDATE application SET checked=$1 where id=$1;`
// 	_, err = repo.DB.Exec(Qupdate, status, id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return err
// }

func (ar *AuthenticationRepository) VerifyPasswordHash(login, password string) bool {

	//*STEP1:GENERATE HASH&SALT FROM GIVEN PASSWORD
	// var hashPassword string
	// h := md5.New()
	// h.Write([]byte(password))
	// newHashPassword := fmt.Sprintf("%x", h.Sum(nil))

	//* Generate "hash" to store from user password
	// recenthashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(recenthashPassword), "gen from", password)
	recenthashPassword := password
	//*STEP2:GET HASH&SALT FROM DB
	var DBHashPassword []byte
	Query := `SELECT password FROM admin WHERE login=$1;`
	ar.DB.QueryRowx(Query, login).Scan(&DBHashPassword)
	fmt.Println(string(DBHashPassword))

	//*STEP3:COMPARE HASHES&SALT
	err := bcrypt.CompareHashAndPassword([]byte(DBHashPassword), []byte(recenthashPassword))
	fmt.Println(string(DBHashPassword))
	if err != nil { //?
		return false
	}
	return true
}

// func (ar *AuthenticationRepository) GetHashPassword(login string) (hashPassword string) {
// 	Query := `select password from admin where login='saidov';`
// 	ar.DB.QueryRowx(Query, login).Scan(&hashPassword)
// 	return hashPassword
// }
