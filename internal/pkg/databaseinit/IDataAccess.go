package databaseinit

type IDataAccess interface {
	InitDataBase() error
	Disconnect()
	GetDriverName()
	SetDriverName()
}
