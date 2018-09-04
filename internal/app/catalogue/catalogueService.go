package catalogue

import (
	"crypto/md5"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

var emailVal = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?)*$")
var nameValidation = regexp.MustCompile("^[a-zA-Z]{3,255}$") // username validation
var phoneNumberVlidation = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)// format phone number (992)4444444444

type CatalogueService struct {
	repo types.ICatalogueRepository
}

func InitService(cRep types.ICatalogueRepository) *CatalogueService {
	return &CatalogueService{
		repo: cRep,
	}
}

func (srv *CatalogueService) CreateUser(us *types.NewServiceProvider) error {
	if !emailVal.MatchString(us.Email) {
		return errors.New("Wrong email")
	}
	if ans, err := srv.repo.Exists(us.Email); ans {
		fmt.Println(ans, err)
		return errors.New("Email Exists!")
	}
	if !nameValidation.MatchString(us.Contact_Name) {
		return errors.New("Wrong name")
	}
	us.Password = srv.HashIt(us.Password)
	us.Registration_Date = time.Time.Format(time.Now(), "01-02-2006")
	us.Is_deleted = false

	srv.repo.InsertRegistretedUser(us)
	us.Password = ""
	us.Registration_Date = ""

	return nil
}

func (srv *CatalogueService) HashIt(someStr string) string {
	h := md5.New()
	h.Write([]byte(someStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (srv *CatalogueService) UpdateServiceProvider(id int, contactName, adress, registrationDate string, cashBackPercent, alifsCashBackShare, usersCashBackShare int, email, password, companyName, city, country string) (count int64, err error) {
	count, err = srv.repo.UpdateServiceProvider(id, contactName, adress, registrationDate, cashBackPercent, alifsCashBackShare, usersCashBackShare, email, password, companyName, city, country)
	return count, err
}

func (srv *CatalogueService) GetMultiServiceProvider(count int) ([]types.NewServiceProvider, error) {
	return srv.repo.GetMultiServiceProvider(count)
}

func (srv *CatalogueService) GetSingleServiceProvider(id int) (types.NewServiceProvider, error) {
	return srv.repo.GetSingleServiceProvider(id)
}
func (srv *CatalogueService) DeleteServiceProvider(id int, isdeleted bool) (count int64, err error) {
	count, err = srv.repo.DeleteServiceProvider(id, isdeleted)
	return
}

func (srv *CatalogueService) CreateMyUser(us *types.NewUsersTable) (err error, userInf *types.UserAuthenticationAnswer) {
	if !nameValidation.MatchString(us.Name) || !nameValidation.MatchString(us.Surname) ||
		!nameValidation.MatchString(us.City) || !nameValidation.MatchString(us.Country) ||
		!phoneNumberVlidation.MatchString(us.PhoneNumber) { 
		return errors.New("Feilds filled wrong!"), userInf
	}
	if !emailVal.MatchString(us.Email) {
		return errors.New("Wrong Email!"), userInf
	}
	if srv.repo.ExistsUser(us.Email) {
		return errors.New("User Exists!"), userInf
	}
	us.RegDate = time.Time.Format(time.Now(), "01-02-2006")
	us.HashPassword = srv.HashIt(us.HashPassword)
	srv.repo.InsertMyUser(us)
	fmt.Println("ok1")
	userInf1 := srv.repo.GetUserInfo(us.Email)
	fmt.Println("userInf:", userInf1)
	return err, &userInf1
}
