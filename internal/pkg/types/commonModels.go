package types

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type DBData struct {
	DriverName       string `json: "DriverName"`
	DataBaseUser     string `json: "DataBaseUser"`
	DataBaseName     string `json: "DataBaseName"`
	DataBasePassword string `json: "DataBasePassword"`
	SSLMode          string `json: "SSLMode"`
}

type Users struct {
	Id                int    `json:"id"`
	Name              string `json:"Name"`
	Surname           string `json:"Surname"`
	Adress            string `json:"Adress"`
	Registration_date string `json:"Registration_date"`
	Email             string `json:"Email"`
	Company_name      string `json:"Company_name"`
	Password_hash     string `json:"Password_hash "`
	Is_deleted        bool   `json:"IsDeleted "`
	Phone_number      string `json:"PhoneNumber"`
	City              string `json:"City"`
	Country           string `json:"Country"`
}

//
type Service_provider struct {
	Id                   int    `json:"id"`
	Contact_name         string `json:"ContactName"`
	Adress               string `json:"Adress"`
	Registration_date    string `json:"RegistrationDate"`
	Cashback_percent     int    `json:"cashbackPercent"`
	Alifs_cashback_share int    `json:"alifsCashbackShare"`
	Users_cashback_share int    `json:"users_cashbackShare"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Company_name         string `json:"CompanyName"`
	Is_deleted           bool   `json:"IsDeleted"`
	City                 string `json:"City"`
	Country              string `json:"Country"`
}
type BalanceTransfer struct {
	Balance         int           `json:"Balance"`
	SenderAccount   AccountStruct `json:"SenderAccount"`
	ReceiverAccount AccountStruct `json:"ReceiverAccount"`
}
type AccountStruct struct {
	Id           int    `json:"id" db:"id"`
	Balance      int    `json:"balance" db:"balance"`
	Owner_id     int    `json:"ownerId" db:"owner_id"`
	Account_type string `json:"accountType" db:"account_type"`
	Owner_type   string `json:"ownerType" db:"owner_type"`
}

type Cashbacks struct {
	id             int `json:"id"`
	transaction_id int `json:"transactionId"`
	amount         int `json:"amount"`
}

type Transactions struct {
	id                int    `json:"id"`
	amount            int    `json:"amount"`
	transaction_type  string `json:"transactionType"`
	date              string `json:"date"`
	source_account_id int    `json:"sourceAccountId"`
	target_account_id int    `json:"targetAccountId"`
	service_id        int    `json:"serviceId "`
}

type Services struct {
	Id                  int    `json:"id"`
	Service_provider_id int    `ServiceProviderId"`
	Name_of_service     string `json:"nameOfService"`
	Price               int    `json:"price"`
	Category            string `json:"category"`
}
type UserAccount struct {
	UserId    int `json:"userId"`
	UserCash  int `json:"userCash"`
	UserBonus int `json:"userBonus"`
}

type ServiceProviderDeletePl struct {
	Id                 int    `json:"id"`
	Name               string `json:"Name"`
	Adress             string `json:"Adress"`
	RegistrationDate   string `json:"RegistrationDate"`
	CashbackPercent    int    `json:"cashbackPercent"`
	AlifsCashbackShare int    `json:"alifsCashbackShare"`
	UsersCashbackShare int    `json:"usersCashbackShare"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	CompanyName        string `json:"CompanyName"`
	Segment            string `json:"segment"`
	Active             bool   `json:"active"`
}

type Discount struct {
	ID                 int    `json:"id" db:"id"`
	Name               string `json:"name" db:"name"`
	DiscountPercentage int    `json:"DiscountPercentage" db:"discount_percentage"`
	ServiceProviderID  int    `json:"serviceProviderID" db:"service_provider_id"`
	DiscountState      string `json:"discountState" db:"discount_state"`
	StartDate          string `json:"startDate" db:"start_date"`
	EndDate            string `json:"endDate" db:"end_date"`
}

type PastDiscount struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type TransactionAmount struct {
	ServiceProviderName string
	Segment             string
	Address             string
	AllTransactions     int64
	AverageTransaction  int32
	MinTransaction      int32
	MaxTransaction      int32
}

type TransactionPerevodi struct {
	Id               int    `json:"id"`
	Date             string `json:"date"`
	Transaction_type string `json:"transactionType"`
	Amount           string `json:"amount"`
	Name             string `json:"Name1"`
	Surname          string `json:"Surname1"`
	Name_u2          string `json:"Name2"`
	Surname_u2       string `json:"Surname2"`
}

type NewServiceProvider struct {
	ID                   int    `json: "ID"`
	Contact_Name         string `json: "ContactName"`
	Adress               string `json: "Adress"`
	Registration_Date    string `json: "RegistrationDate"`
	CashBack_Percent     int    `json: "CashBackPercent"`
	Alifs_CashBack_Share int    `json: "AlifsCashBackShare"`
	Users_CashBack_Share int    `json: "UsersCashBackShare"`
	Email                string `json: "Email"`
	Password             string `json: "Password"`
	Company_Name         string `json:"CompanyName"`
	Is_deleted           bool   `json:"Isdeleted"`
	City                 string `json:"City"`
	Country              string `json:"Country"`
}
type Answer struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Info    interface{} `json:"info"`
}

type UserTransactionRequest struct {
	Id       int    `json:"id"`
	Count    int    `json:"count"`
	DateFrom string `json:"datefrom"`
	DateTo   string `json:"dateto"`
}

type ForPurchaseByCash struct {
	UserId    int `json:"UserId"`
	ServiseId int `json:"ServiseId"`
}

type UserProviderAuthentication struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type NewUsersTable struct {
	Name         string `json:"Firstname"`
	Surname      string `json:"Lastname"`
	Adress       string `json:"Adress"`
	RegDate      string `json:"RegDate"`
	Email        string `json:"Email"`
	HashPassword string `json:"HashPassword"`
	IsDeleted    bool   `json:"IsDeleted"`
	PhoneNumber  string `json:"PhoneNumber"`
	City         string `json:"City"`
	Country      string `json:"Country"`
}

type UserAuthenticationAnswer struct {
	Email        string        `json:"Email"`
	Name         string        `json:"UserName"`
	Surname      string        `json:"Surname"`
	AccountBonus AccountStruct `json:"UserAccountBonus"`
	AccountCash  AccountStruct `json:"UserAccountCash"`
}

type UserForAnalitics struct {
	Name                  string `json:"name"`
	Surname               string `json:"surname"`
	Adress                string `json:"Adress"`
	City                  string `json:"city"`
	Registration_date     string `json:"registration_date"`
	Max_transaction       int    `json:"max_transaction"`
	Avg_transaction       int    `json:"avg_transaction"`
	Count_of_transactions int    `json:"count_of_transactions"`
}

type UserInfo struct {
	ID               int            `json: "ID"`
	Name             string         `json: "Name"`
	Surname          sql.NullString `json: "Surname"`
	Address          sql.NullString `json: "Address"`
	Phone_number     string         `json: "Phone_number"`
	Email            sql.NullString `json: "Email"`
	RegistrationDate pq.NullTime    `json: "RegistrationDate"`
	City             sql.NullString `json: "City"`
	Country          sql.NullString `json: "Country"`
}

type UserName struct {
	Name_user    string `json: "Name_user"`
	Surname_user string `json: "SurName_user"`
}
type UserTranInfoRequest struct {
	Id       int    `json:"id"`
	DateFrom string `json:"datefrom"`
	DateTo   string `json:"dateto"`
}

type TranUserInfoPakupki struct {
	Amount           string `json:"amount"`
	Cashback         string `json:"cashback"`
	Type             string `json:"type"`
	Date             string `json:"date"`
	Name             string `json:"Name"`
	Surname          string `json:"Surname"`
	Company_name     string `json:"company_name"`
	Cashback_percent string `json:"cashback_percent"`
}

type TransactionPakupki struct {
	Id               int    `json:"id"`
	Amount           string `json:"amount"`
	Transaction_type string `json:"transactionType"`
	Date             string `json:"date"`
	Name             string `json:"Name"`
	Surname          string `json:"Surname"`
	Company_name     string `json:"CompanyName"`
	Cashback_percent string `json:"CashbackPercent"`
}

type TranUserInfoPerevodi struct {
	Amount     string `json:"amount"`
	Cashback   string `json:"cashback"`
	Type       string `json:"type"`
	Date       string `json:"date"`
	Name       string `json:"name1"`
	Surname    string `json:"surname1"`
	Name_u2    string `json:"name2"`
	Surname_u2 string `json:"surname2"`
}

type EntryServiceProvider struct {
	Id         int    `json:"Id"`
	Count      int    `json:"Count"`
	Is_deleted bool   `json:"Is_deleted"`
	City       string `json:"City"`
}

type StoryCashback struct {
	ID               int       `json:"ID"`
	Name             string    `json:"Name"`
	Surname          string    `json:"Surname"`
	Transaction_type string    `json:"Transaction_type"`
	Date             time.Time `json:"Data"`
	Amount           int       `json:"Amount"`
}

type InfoServiceProvider struct {
	ID                   int    `json:"ID"`
	Company_Name         string `json:"CompanyName"`
	Adress               string `json:"Address,omitempty"`
	City                 string `json:"City,omitempty"`
	Users_Cashback_Share int    `json:"UsersCashbackShare,omitempty"`
	Name_Of_Service      string `json:"NameOfService,omitempty"`
	Price                int    `json:"Price,omitempty"`
	Category             string `json:"Category,omitempty"`
}

type ListInTouchProduct struct {
	ID              int    `json:"ID"`
	Company_name    string `json:"CompanyName"`
	City            string `json:"City"`
	Name_Of_Product string `json:"NameOfProduct"`
	Category        string `json:"Category"`
	Price           int    `json:"Price"`
	Balance         int    `json:"Balance"`
}

type ServicesProviderDateRequest struct {
	DateFrom string `json:"datefrom"`
	DateTo   string `json:"dateto"`
}

type ServiceProviderAllPurchase struct {
	Id               int    `json:"id"`
	Company_name     string `json:"companyName"`
	Cashback_percent int    `json:"cashbackPercent"`
	SumAmount        int    `json:"sumamount"`
	SumCashback      int    `json:"sumcashback"`
	Count            int    `json:"count"`
	DateFrom         string `json:"datefrom"`
	DateTo           string `json:"dateto"`
}

type ActiveDiscounts struct {
	Company_name        string    `json:"CompanyName"`
	City                string    `json:"City"`
	Name_of_discount    string    `json:"NameOfDiscount"`
	Discount_percentage int       `json:"DiscountPercentage"`
	Discount_state      string    `json:"DiscountState"`
	Start_date          time.Time `json:"StartDate"`
	End_date            time.Time `json:"EndDate"`
}

type ServiceProviderRating struct {
	Contact_Name                string `json:"ContactName"`
	Category                    string `json:"Category"`
	Count_of_bonus_transactions int    `json:"CountOfBonusTransactions"`
	Count_of_cash_transactions  int    `json:"CountOfCashTransactions"`
}

//Annual Report
type Period struct {
	Start string `json: "start"`
	End   string `json: "end"`
}

type Row struct {
	City  string         `json:"City"`
	Count pq.StringArray `json:"Count"`
}

type ServiceProviderForController struct {
	City     string `json:"City"`
	Purchase string `json:"Purchase"`
}
