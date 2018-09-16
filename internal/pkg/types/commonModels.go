package types

type DBData struct {
	DriverName       string `json: "DriverName"`
	DataBaseUser     string `json: "DataBaseUser"`
	DataBaseName     string `json: "DataBaseName"`
	DataBasePassword string `json: "DataBasePassword"`
	SSLMode          string `json: "SSLMode"`
}

// type Server struct {
// 	router *Routers
// }

//
type Services struct {
	Installment bool `json:"id"`
	Deposite    bool `json:"id"`
	CreditCard  bool `json:"id"`
	UsingApi    bool `json:"id"`
	None        bool `json:"id"`
}

type UserAuth struct {
	Id                int    `json:"id"`
	RegistrationDate  string `json:"Registration_date"`
	FullName          string `json:"Name"`
	Contacts          string `json:"Contacts"`
	SerialNumber      string `json:SerialNumber`
	Registration_date string
	Services          Services `json:SerialNumber`
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
