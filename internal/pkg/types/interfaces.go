package types

type IAdminRepository interface {
	GetUserInfo(pr *UserName) ([]UserInfo, error)
}

type IAnalyticsRepository interface {
	//GetMostHihgCashBack(count int) ([]NewServiceProvider, error)//TODO:DO NOT WORK
	ServiceAllPurchase(pr *ServicesProviderDateRequest) ([]ServiceProviderAllPurchase, error)
	ServiceProviderRating(city string) ([]ServiceProviderRating, error)
	// TransactionAmount() error
	UsersActivity() ([]UserForAnalitics, error)
	GetServiceProviderAnnualReport(start, end string) ([]Row, error)
}

type IAuthenticationRepository interface {
	EmailCheck(string) bool
	EmailCheckSP(string) bool
	GetHashPassword(string) string
	GetHashPasswordSP(string) string
	GetUserInfo(string) *UserAuthenticationAnswer
	GetUserInfoSP(string) ServiceProviderDeletePl
}

type ICashbackRepository interface {
	GetAccounts(int, string, string) (AccountStruct, error)
	GetUserAccount(BalanceTransfer) BalanceTransfer
	GetService(int) (Services, error)
	GetServiceProvider(int) (Service_provider, error)
	GetTransactionId() (int, error)
	GetTransactionListPakupki(*UserTransactionRequest) ([]TransactionPakupki, error)
	GetTransactionListPerevodi(*UserTransactionRequest) ([]TransactionPerevodi, error)
	SaveTransactionInfoIntoDBTable(int, Services, string) error
	SaveTransactionToDB(AccountStruct, AccountStruct, AccountStruct, AccountStruct) error
	SaveTransactionByBonusToDB(AccountStruct, AccountStruct) error
	SaveExchange(int, int, int)
	SaveCashbackInfo(int, int) error
	TranInfoUserPakupki(*UserTranInfoRequest) ([]TranUserInfoPakupki, error)
	TranInfoUserPerevodi(*UserTranInfoRequest) ([]TranUserInfoPerevodi, error)
	SaveTransferTransaction(BalanceTransfer) (err error)
}

type ICatalogueRepository interface {
	GetUserInfo(string) UserAuthenticationAnswer
	InsertMyUser(*NewUsersTable)
	ExistsUser(string) bool
	InsertRegistretedUser(*NewServiceProvider)
	Exists(string) (bool, error)
	UpdateServiceProvider(int, string, string, string, int, int, int, string, string, string, string, string) (int64, error)
	GetMultiServiceProvider(int) ([]NewServiceProvider, error)
	GetSingleServiceProvider(int) (NewServiceProvider, error)
	DeleteServiceProvider(int, bool) (int64, error)
}

type IDiscountRepository interface {
	CreateDiscount(Discount) error
	DiscountExist(string, int) bool
	DiscountsList() ([]Discount, error)
	DiscountsActiveList() ([]Discount, error)
	DiscountsSoonList() ([]Discount, error)
	DiscountsPastList(PastDiscount) ([]Discount, error)
	ParseDate(PastDiscount) error
}

type IUsersRepository interface {
	GetStoryCash(id, count int) ([]StoryCashback, error)
	GetInfoService_Provider(count int, city string) ([]InfoServiceProvider, error)
	GetAllServices(id int) ([]InfoServiceProvider, error)
	GetListInTouchProduct(id int, city string) ([]ListInTouchProduct, error)
	GetListActiveDiscounts(city string, id int) ([]ActiveDiscounts, error)
	GetNewDiscounts(id int, city string) ([]ActiveDiscounts, error)
}
