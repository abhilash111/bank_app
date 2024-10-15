package domain

import (
	"database/sql"

	"github.com/abhilash111/bank_app/errs"
	"github.com/abhilash111/bank_app/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllQuery := "SELECT * FROM customers"
	customers := make([]Customer, 0)
	err := c.client.Select(&customers, findAllQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customers found" + err.Error())
			return nil, errs.NewNotFoundError("No Customers Found")
		} else {
			logger.Error("Un Expected Database error" + err.Error())
			return nil, errs.NewUnExpectedError("Un Expected Database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	custmerQuery := "select * from customers where customer_id = $1 "
	var c Customer
	err := d.client.Get(&c, custmerQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customer found" + err.Error())
			return nil, errs.NewNotFoundError(`Customer Not Found`)
		} else {
			logger.Error("Error While scanning customer" + err.Error())
			return nil, errs.NewUnExpectedError(`Unexpected Database error`)
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}
