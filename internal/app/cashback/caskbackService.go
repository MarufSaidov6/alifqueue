package cashback

import (
	"errors"
	"fmt"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/sirupsen/logrus"
)

const BonusCoef = 1

type CashbackService struct {
	repo types.ICashbackRepository
}

func InitService(cRep types.ICashbackRepository) *CashbackService {
	return &CashbackService{
		repo: cRep,
	}
}

const alifID = 4

func (srv *CashbackService) GetBonusInfo(userid int) (int, error) {
	var (
		err error
		a   types.AccountStruct
	)

	a, err = srv.repo.GetAccounts(userid, "user", "bonus")

	return a.Balance, err
}

var NotEnoughAmountErr = errors.New("Not enough cash at your account")

func (srv *CashbackService) PurchaseByCash(userId int, serviceId int) error {
	var transactionId int

	userCash, err := srv.repo.GetAccounts(userId, "cash", "user")
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR 1: ", err)
		return err
	}

	service, err := srv.repo.GetService(serviceId)
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR 2: ", err)
		return err
	}

	userBonus, err := srv.repo.GetAccounts(userId, "bonus", "user")
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR 3: ", err)
		return err
	}
	SPCash, err := srv.repo.GetAccounts(service.Service_provider_id, "cash", "service_provider")
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR 4: ", err)
		return err
	}
	serviceProvider, err := srv.repo.GetServiceProvider(service.Service_provider_id)
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR 5: ", err)
		return err
	}

	alifBonus, err := srv.repo.GetAccounts(alifID, "cash", "service_provider")

	if userCash.Balance < service.Price {
		err = NotEnoughAmountErr
		logrus.Error(err)
		fmt.Println("ERROR 6: ", err)

		return err
	}

	cashback := (service.Price * serviceProvider.Cashback_percent) / 100
	userCash.Balance -= service.Price
	SPCash.Balance += service.Price - cashback
	userBonus.Balance += (cashback * (serviceProvider.Users_cashback_share / 100) / 10)
	alifBonus.Balance += (cashback * (serviceProvider.Alifs_cashback_share / 100) / 10)

	err = srv.repo.SaveTransactionToDB(userCash, SPCash, userBonus, alifBonus)
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of saveTransactionToDB: ", err)
		return err
	}

	err = srv.repo.SaveTransactionInfoIntoDBTable(userId, service, "purchase")
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of SaveTransactionInfoIntoDBTable: ", err)
		return err
	}

	transactionId, err = srv.repo.GetTransactionId()
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of getTransactionId: ", err)
		return err
	}

	err = srv.repo.SaveCashbackInfo(transactionId, cashback)
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of SaveCashbackInfo: ", err)
		return err
	}

	return err
}

func (srv *CashbackService) PurchaseByBonus(userId int, serviceId int) error {
	var (
		user    types.AccountStruct
		sp      types.AccountStruct
		service types.Services
		err     error
	)

	user, err = srv.repo.GetAccounts(userId, "bonus", "user")
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of get user account: ", err)
		return err
	}

	service, err = srv.repo.GetService(serviceId)
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of get service: ", err)
		return err
	}

	sp, err = srv.repo.GetAccounts(service.Service_provider_id, "cash", "service_provider")
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of get service provider: ", err)
		return err
	}

	if user.Balance < service.Price {
		err = NotEnoughAmountErr
		logrus.Error(err)
		fmt.Println("ERROR: ", err)

		return err
	}

	user.Balance -= service.Price
	sp.Balance += service.Price

	err = srv.repo.SaveTransactionByBonusToDB(sp, user)
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of saveTransactionByBonusToDB: ", err)
		return err
	}

	err = srv.repo.SaveTransactionInfoIntoDBTable(user.Owner_id, service, "purchase_bonus")
	if err != nil {
		logrus.Error(err)
		fmt.Println("ERROR of SaveTransactionInfoIntoDBTable: ", err)
		return err
	}

	return err
}

func (srv *CashbackService) DoExchange(userId int, Bonus int) (us types.UserAccount, err error) {

	userCash, err := srv.repo.GetAccounts(userId, "cash", "user")
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR: ", err)
		return us, err
	}

	userBonus, err := srv.repo.GetAccounts(userId, "bonus", "user")
	if err != nil {
		logrus.Warn(err)
		fmt.Println("ERROR: ", err)
		return us, err
	}

	userBonus.Balance -= Bonus

	if userBonus.Balance < 0 {
		return us, errors.New("Not enough bonus in ur account")

	}

	userCash.Balance += Bonus * BonusCoef
	srv.repo.SaveExchange(userId, userCash.Balance, userBonus.Balance)
	us = types.UserAccount{
		userId,
		userCash.Balance,
		userBonus.Balance,
	}
	return us, nil
}

// Происходить проверка счета
// и запускаеться функ перевода
func (srv *CashbackService) DoTransfer(transfer types.BalanceTransfer) (error, types.BalanceTransfer) {
	var err error
	if transfer.SenderAccount.Balance-transfer.Balance < 0 {
		return errors.New("Balance not enough!"), transfer
	}
	if transfer.SenderAccount.Account_type != transfer.ReceiverAccount.Account_type {
		return errors.New("Wrong with Accounts type"), transfer
	}
	err = srv.repo.SaveTransferTransaction(transfer)
	transfer = srv.repo.GetUserAccount(transfer)
	return err, transfer
}
