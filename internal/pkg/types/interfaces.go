package types

//!Blacknote
type IAuthenticationRepository interface {
	VerifyLogin(string) bool
	VerifyPasswordHash(string, string) bool
	InsertUser(user UserAuth) error
	GetPersons() ([]GetUsers, error)
	// UpdateServiceProvider(status *bool, id int) (err error)
}
