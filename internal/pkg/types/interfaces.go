package types

//!Blacknote
type IAuthenticationRepository interface {
	VerifyLogin(string) bool
	VerifyPasswordHash(string, string) bool
}
