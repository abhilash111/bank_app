package domain

import (
	"fmt"
	"strconv"

	"github.com/abhilash111/bank_app/errs"
	"github.com/abhilash111/bank_app/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `
	INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING account_id
	`
	var id int64
	err := d.client.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)

	if err != nil {
		fmt.Println("Err", err)
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database error")
	}
	if err != nil {
		logger.Error("Error while Fetching Lastinserted Id from new account" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Database error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	// inserting bank account transaction and returning the transaction_id
	var transactionId int64
	err = tx.QueryRow(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											VALUES ($1, $2, $3, $4) RETURNING transaction_id`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionId)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 WHERE account_id = $2`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 WHERE account_id = $2`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	// updating the transaction struct with the latest balance and transaction id
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
