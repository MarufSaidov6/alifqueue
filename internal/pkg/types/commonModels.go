package types

type DBData struct {
	DriverName       string `json: "DriverName"`
	DataBaseUser     string `json: "DataBaseUser"`
	DataBaseName     string `json: "DataBaseName"`
	DataBasePassword string `json: "DataBasePassword"`
	SSLMode          string `json: "SSLMode"`
}

//!STORE USER
type UserAuth struct {
	Id               int    `json:"id"`
	FullName         string `json:"FIO"`
	Contact          string `json:"Contact"`
	SerialNumber     string `json:"SerialNumber"`
	RegistrationDate string `json:"Registrationdate"`
	PurchaseDateTime string `json:"PurchaseDateTime"`
	Services         []int  `json:"Services"` //! sql:"default: false"
}

// type Services struct {
// 	Installment bool `json:"nasiya"`
// 	Deposite    bool `json:"Installment"`
// 	CreditCard  bool `json:"CreditCard"`
// 	UsingApi    bool `json:"UsingApi"`
// 	None        bool `json:"None"`
// }

//!GET USERS
type GetUsers struct {
	Id               int    `json:"Id"`
	FullName         string `json:"FullName"`
	Contact          string `json:"Contact"`
	SerialNumber     string `json:"SerialNumber"`
	PurchaseDateTime string `json:"PurchaseDateTime"`
	Сhecked          bool   `json:"Checked" schema:"my_field" sql:"default: false"` //!ATTENTION!
}

//!

type AdminAuth struct {
	Login        string `json:"Login"`
	PasswordHash string `json:"PasswordHash"`
}

type Answer struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Info    interface{} `json:"info"`
}

type Application struct {
	Id          int                `json:"id"`
	Fio         string             `json:"fullname"`
	PhoneNumber string             `json:"phoneNumber"`
	SNumber     string             `json:"serialnumber"`
	RegDate     string             `json:"regDate"`
	Services    map[int]AppService `json:"services"` // map[id]appServ
}

type AppService struct {
	ServiceName string `json:serviceName`
}
