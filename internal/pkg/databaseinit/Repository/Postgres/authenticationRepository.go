package databaseinit

import (
	"github.com/AlifElectronicQueue/internal/pkg/types"

	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func (repo *AuthenticationRepository) VerifyPasswordHash(login, password string) bool {

	var DBpassword []byte
	query := `SELECT password FROM admin WHERE login=$1;`
	repo.DB.QueryRowx(query, login).Scan(&DBpassword)

	err := bcrypt.CompareHashAndPassword([]byte(DBpassword), []byte(password))

	if err != nil {
		log.Error("Something failed, login or password")
		return false
	}
	return true
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
	query := "select id,fullname,contact,serialnumber,purchasedate,checked from application;"
	rows, err = repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&input.Id, &input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
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

	query := fmt.Sprintf("select id,fullname,contact,serialnumber,purchasedate,checked from application order by %d", ordered)

	rows, err = repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(rows)
	for rows.Next() {
		err = rows.Scan(&input.Id, &input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
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
		query  = `SELECT id,fullname,contact,serialnumber,purchasedate,checked FROM application WHERE id = $1;`
		output []types.GetUsers
		input  types.GetUsers
		err    error
		rows   *sqlx.Rows
	)
	rows, _ = repo.DB.Queryx(query, id)
	for rows.Next() {
		err = rows.Scan(&input.Id, &input.FullName, &input.Contact, &input.SerialNumber, &input.PurchaseDateTime, &input.Сhecked)
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

func (repo *AuthenticationRepository) UpdateApplicationStatusById(checked string, id int) (err error) {
	fmt.Println("CGCV", checked, id)
	if checked != "" {
		checked = "true"
	} else {
		checked = "false"
	}
	Qupdate := fmt.Sprintf("UPDATE application SET checked=%s where id=%d;", checked, id)
	fmt.Println("check repo", Qupdate, checked)
	_, err = repo.DB.Exec(Qupdate)
	if err != nil {
		panic(err)
	}

	return err
}
