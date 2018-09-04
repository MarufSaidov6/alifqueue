package databaseinit

import (
	"database/sql"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type AdminRepository struct {
	DB *sqlx.DB
}

//GetUserInfo returns data of user by name and surname
func (repo *AdminRepository) GetUserInfo(pr *types.UserName) ([]types.UserInfo, error) {
	var (
		rows *sql.Rows
		err  error
		ps   []types.UserInfo
		p    types.UserInfo
	)
	query := `
			SELECT u.id, u.name, u.surname, u.registration_date, u.adress, u.city, u.country,
				   u.email, u.phone_number 
			FROM users u WHERE u.name=$1 OR u.surname=$2
		    `
	rows, err = repo.DB.Query(query, pr.Name_user, pr.Surname_user)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Address, &p.Phone_number, &p.Email, &p.RegistrationDate, &p.City, &p.Country)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, err

}
