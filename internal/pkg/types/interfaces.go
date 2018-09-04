package types

//!Blacknote
type IAdminRepository interface {
	GetUserInfo(pr *UserName) ([]UserInfo, error)
}

//!Blacknote
type IAuthenticationRepository interface {
	EmailCheck(string) bool
	EmailCheckSP(string) bool
	GetHashPassword(string) string
	GetHashPasswordSP(string) string
	GetUserInfo(string) *UserAuthenticationAnswer
	GetUserInfoSP(string) ServiceProviderDeletePl
}
