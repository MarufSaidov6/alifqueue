package databaseinit

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/sirupsen/logrus"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

type CashbackRepository struct {
	DB *sqlx.DB
}

func (repo *CashbackRepository) GetAccounts(owner_id int, account_type string, owner_type string) (types.AccountStruct, error) {
	var (
		err     error
		query   = "SELECT * FROM accounts where owner_id = $1 and account_type = $2 and owner_type = $3;"
		account types.AccountStruct
	)
	err = repo.DB.QueryRowx(query, owner_id, account_type, owner_type).StructScan(&account)
	return account, err
}

func (repo *CashbackRepository) GetTransactionId() (int, error) {
	var (
		err           error
		query         = "SELECT id from transactions ORDER BY ID DESC LIMIT 1;"
		transactionId int
	)
	row := repo.DB.QueryRow(query)
	err = row.Scan(&transactionId)
	return transactionId, err
}

func (repo *CashbackRepository) GetTransactionListPakupki(pr *types.UserTransactionRequest) ([]types.TransactionPakupki, error) {
	var (
		query = `select  t.id, t.amount, t.transaction_type, t.date, u.name, u.surname, p.company_name, p.cashback_percent
	from transactions as t, users as u, service_provider as p where t.source_account_id = u.id and
	 t.target_account_id = p.id and t.transaction_type = 'purchase' `
		err error
		ps  []types.TransactionPakupki
	)
	if pr.Id == 0 {
		return nil, errors.New("Add the client id")
	} else if pr.Count != 0 && pr.DateFrom != "" && pr.DateTo != "" {
		err = repo.DB.Select(&ps, query+` and u.id = $1 and t.date between $2 and $3 order by  t.date desc limit $4`, pr.Id, pr.DateFrom, pr.DateTo, pr.Count)
		if err != nil {
			return nil, err
		}
		return ps, err
	} else if pr.Count == 0 && pr.DateFrom != "" && pr.DateTo != "" {
		err = repo.DB.Select(&ps, query+` and u.id=$1 and t.date between $2 and $3 order by  t.date desc;`, pr.Id, pr.DateFrom, pr.DateTo)
		if err != nil {
			return nil, err
		}
		return ps, err
	} else if pr.Count != 0 && pr.DateFrom == "" && pr.DateTo == "" {
		err = repo.DB.Select(&ps, query+` and u.id=$1 order by  t.date desc limit $2;`, pr.Id, pr.Count)
		if err != nil {
			return nil, err
		}
		return ps, err
	}
	return ps, err

}

func (repo *CashbackRepository) GetTransactionListPerevodi(pr *types.UserTransactionRequest) ([]types.TransactionPerevodi, error) {
	var (
		err  error
		ps   []types.TransactionPerevodi
		p    types.TransactionPerevodi
		rows *sql.Rows
	)
	query := `select t.id, t.amount, t.transaction_type, t.date, u.name, u.surname, u2.name, u2.surname
	from transactions as t, users as u, users as u2 where t.source_account_id = u.id and
	 t.target_account_id=u2.id and t.transaction_type='transfer'`
	if pr.Id == 0 {
		return nil, errors.New("Add the client id")
	} else if pr.Count != 0 && pr.DateFrom != "" && pr.DateTo != "" {
		rows, err = repo.DB.Query(query+` and u.id=$1 and t.date between $2 and $3 order by  t.date desc limit $4`, pr.Id, pr.DateFrom, pr.DateTo, pr.Count)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err = rows.Scan(&p.Id, &p.Amount, &p.Transaction_type, &p.Date, &p.Name, &p.Surname, &p.Name_u2, &p.Surname_u2)
			if err != nil {
				return nil, err
			}
			ps = append(ps, p)
		}
		return ps, err
	} else if pr.Count == 0 && pr.DateFrom != "" && pr.DateTo != "" {
		rows, err = repo.DB.Query(query+` and u.id=$1 and t.date between $2 and $3 order by  t.date desc;`, pr.Id, pr.DateFrom, pr.DateTo)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err = rows.Scan(&p.Id, &p.Date, &p.Transaction_type, &p.Amount, &p.Name, &p.Surname, &p.Name_u2, &p.Surname_u2)
			if err != nil {
				return nil, err
			}
			ps = append(ps, p)
		}
		return ps, err
	} else if pr.Count != 0 && pr.DateFrom == "" && pr.DateTo == "" {
		rows, err = repo.DB.Query(query+` and u.id=$1 order by  t.date desc limit $2;`, pr.Id, pr.Count)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err = rows.Scan(&p.Id, &p.Date, &p.Transaction_type, &p.Amount, &p.Name, &p.Surname, &p.Name_u2, &p.Surname_u2)
			if err != nil {
				return nil, err
			}
			ps = append(ps, p)
		}
		return ps, err
	}
	return ps, err
}

func (repo *CashbackRepository) GetService(id int) (types.Services, error) {
	var (
		err   error
		query = "SELECT * FROM services WHERE id = $1;"
		s     types.Services
	)
	err = repo.DB.QueryRowx(query, id).StructScan(&s)
	return s, err
}

func (repo *CashbackRepository) GetServiceProvider(id int) (types.Service_provider, error) {
	var (
		err   error
		query = "SELECT * FROM service_provider WHERE id = $1;"
		sp    types.Service_provider
	)
	err = repo.DB.QueryRowx(query, id).StructScan(&sp)
	return sp, err

}

func (repo *CashbackRepository) SaveCashbackInfo(transactionID int, amount int) error {
	var (
		qInsert = "INSERT INTO cashbacks (transaction_id, amount) VALUES ($1, $2);"
	)

	_, err := repo.DB.Exec(qInsert, transactionID, amount)

	return err
}

func (repo *CashbackRepository) SaveExchange(id, cashBalance, bonusBalance int) {
	tx := repo.DB.MustBegin()
	inq := `update accounts set balance=$2 where account_type='cash'
		and owner_type='user' and owner_id=$1`
	tx.MustExec(inq, id, cashBalance)
	inq = `update accounts set balance=$2 where account_type='bonus'
		and owner_type='user' and owner_id=$1`
	tx.MustExec(inq, id, bonusBalance)
	tx.Commit()
}

func (repo *CashbackRepository) SaveTransactionInfoIntoDBTable(userID int, s types.Services, tr_type string) error {
	var (
		qInsert = "INSERT INTO transactions (amount, transaction_type, date, source_account_id, target_account_id, service_id) VALUES ($1, $2, to_date($3,'yyyy:mm:dd'), $4, $5, $6);"
	)

	current_time := time.Now().Local()

	_, err := repo.DB.Exec(qInsert, s.Price, tr_type, current_time.Format("2006-01-02"), userID, s.Service_provider_id, s.Id)

	return err
}

func (repo *CashbackRepository) SaveTransactionToDB(userCash types.AccountStruct, SPCash types.AccountStruct, userBonus types.AccountStruct, alifBonus types.AccountStruct) error {
	var (
		//забираем деньги у пользователя
		subtructPriceFromUser = "UPDATE accounts SET balance = $1 WHERE owner_id=$2 AND account_type ='cash' AND owner_type = 'user';"
		//даем деньги провайдеру
		giveMoneyPriceToSP = "UPDATE accounts SET balance=$1 WHERE owner_id=$2 AND account_type = 'cash' AND owner_type = 'service_provider';"
		//даем пользователю кэшбэк
		giveBonusToUser = "UPDATE accounts SET balance=$1 WHERE owner_id=$2 AND account_type ='bonus' AND owner_type = 'user';"
		//даем алифу кэшбэк
		giveBonusToAlif = "UPDATE accounts SET balance=$1 WHERE owner_id=4 AND account_type = 'cash'  AND owner_type = 'service_provider';"
		err             error
	)
	tx, err := repo.DB.Begin()
	if err != nil {
		logrus.Warn(err)
	}

	_, err = tx.Exec(subtructPriceFromUser, userCash.Balance, userCash.Owner_id)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(giveMoneyPriceToSP, SPCash.Balance, SPCash.Owner_id)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(giveBonusToUser, userBonus.Balance, userBonus.Owner_id)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(giveBonusToAlif, alifBonus.Balance)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (repo *CashbackRepository) SaveTransactionByBonusToDB(sp types.AccountStruct, user types.AccountStruct) error {
	var (
		//забираем бонусы у пользователя
		subtructPriceFromUser = "UPDATE accounts SET balance = $1 WHERE owner_id=$2 AND account_type ='bonus' AND owner_type = 'user';"
		//даем деньги провайдеру
		giveMoneyPriceToSP = "UPDATE accounts SET balance=$1 WHERE owner_id=$2 AND account_type = 'cash' AND owner_type = 'service_provider';"
		err                error
	)

	tx, err := repo.DB.Begin()
	if err != nil {
		logrus.Warn(err)
	}

	_, err = tx.Exec(subtructPriceFromUser, user.Balance, user.Owner_id)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(giveMoneyPriceToSP, sp.Balance, sp.Owner_id)
	if err != nil {
		logrus.Warn(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (repo *CashbackRepository) TranInfoUserPakupki(pr *types.UserTranInfoRequest) ([]types.TranUserInfoPakupki, error) {
	var (
		rows *sql.Rows
		err  error
		ps   []types.TranUserInfoPakupki
		p    types.TranUserInfoPakupki
	)
	query := `select * from (select t.amount as tamount, c.amount as camount, t.transaction_type, t.date, u.name, u.surname, p.company_name, p.cashback_percent
		from transactions t, users  u, service_provider  p, cashbacks c where t.source_account_id = u.id and
		 t.target_account_id=p.id and t.transaction_type='purchase' and u.id=$1 and t.date   between $2 and $3 union
	select sum(t.amount), sum(c.amount), t.transaction_type, t.date, u.name, u.surname, p.company_name, p.cashback_percent
		from transactions t, users  u, service_provider  p, cashbacks c where t.source_account_id = u.id and
		 t.target_account_id=p.id and t.transaction_type='purchase' and u.id=$1 group by t.transaction_type, t.date, u.name, u.surname, p.company_name, p.cashback_percent)
	as main order by tamount;`
	rows, err = repo.DB.Query(query, pr.Id, pr.DateFrom, pr.DateTo)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.Amount, &p.Cashback, &p.Type, &p.Date, &p.Name, &p.Surname, &p.Company_name, &p.Cashback_percent)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, err
}

func (repo *CashbackRepository) TranInfoUserPerevodi(pr *types.UserTranInfoRequest) ([]types.TranUserInfoPerevodi, error) {
	var (
		rows *sql.Rows
		err  error
		ps   []types.TranUserInfoPerevodi
		p    types.TranUserInfoPerevodi
	)
	query := `select * from (select t.amount as tamount, c.amount as camount, t.transaction_type, t.date, u.name, u.surname, u2.name, u2.surname
		from transactions as t, users as u, users as u2, cashbacks as c where t.source_account_id = u.id and
		 t.target_account_id=u2.id and t.id=c.transaction_id and t.transaction_type='transfer' and u.id=$1 and t.date
	  between $2 and $3 union
	select sum(t.amount), sum(c.amount), t.transaction_type, t.date, u.name, u.surname, u2.name, u2.surname
		from transactions as t, users as u, users as u2, cashbacks as c where t.source_account_id = u.id and
		 t.target_account_id=u2.id and t.id=c.transaction_id and t.transaction_type='transfer' and u.id=$1 group by t.transaction_type, t.date, u.name, u.surname, u2.name, u2.surname)
	as main order by tamount;`
	rows, err = repo.DB.Query(query, pr.Id, pr.DateFrom, pr.DateTo)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.Amount, &p.Cashback, &p.Type, &p.Date, &p.Name, &p.Surname, &p.Name_u2, &p.Surname_u2)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, err
}

// Функ для транзакции перевода денег или бонуса
func (repo *CashbackRepository) SaveTransferTransaction(transfer types.BalanceTransfer) (err error) {
	tx := repo.DB.MustBegin()
	if transfer.SenderAccount.Account_type == "bonus" {
		// меняем счет юзера для преревода бонусов
		inq := `update accounts set balance=$1 where id=$2 and account_type=$3`
		_, err = tx.Exec(inq, transfer.SenderAccount.Balance-transfer.Balance, transfer.SenderAccount.Id,
			transfer.SenderAccount.Account_type)

		inq = `insert into transactions  values($1, $2, $3, $4, $5)`
		date := time.Time.Format(time.Now(), "01-02-2006")
		_, err = tx.Exec(inq, transfer.Balance, "transferbybonus", date, transfer.SenderAccount.Id,
			transfer.ReceiverAccount.Id, 1113)

		inq = `update accounts set balance=$1 where id=$2 and account_type=$3`
		_, err = tx.Exec(inq, transfer.ReceiverAccount.Balance+transfer.Balance, transfer.ReceiverAccount.Id,
			transfer.ReceiverAccount.Account_type)
		tx.Commit()
		return err
	} else if transfer.SenderAccount.Account_type == "cash" {
		inq := `update accounts set balance=$1 where id=$2 and account_type=$3`
		_, err = tx.Exec(inq, transfer.SenderAccount.Balance-transfer.Balance, transfer.SenderAccount.Id,
			transfer.SenderAccount.Account_type)

		inq = `insert into transactions  values($1, $2, $3, $4, $5)`
		date := time.Time.Format(time.Now(), "01-02-2006")
		_, err = tx.Exec(inq, transfer.Balance, "transferbycash", date, transfer.SenderAccount.Id,
			transfer.ReceiverAccount.Id, 1113)

		inq = `update accounts set balance=$1 where id=$2 and account_type=$3`
		_, err = tx.Exec(inq, transfer.ReceiverAccount.Balance+transfer.Balance, transfer.ReceiverAccount.Id,
			transfer.ReceiverAccount.Account_type)
		tx.Commit()
		return err
	} else {
		errors.New("undefined Transfer Type")
	}
	return
}
func (repo *CashbackRepository) GetUserAccount(transfer types.BalanceTransfer) types.BalanceTransfer {
	inq := `select * from accounts where id=$1`
	repo.DB.QueryRowx(inq, transfer.SenderAccount.Id).StructScan(&transfer.SenderAccount)
	fmt.Println(transfer.SenderAccount)
	repo.DB.QueryRowx(inq, transfer.ReceiverAccount.Id).StructScan(&transfer.ReceiverAccount)
	return transfer
}