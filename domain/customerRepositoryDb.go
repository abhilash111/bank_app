package domain

import (
	"database/sql"

	"github.com/abhilash111/bank_app/errors"
	"github.com/abhilash111/bank_app/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {
	findAllQuery := "SELECT * FROM customers"
	customers := make([]Customer, 0)
	err := c.client.Select(&customers, findAllQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customers found" + err.Error())
			return nil, errors.NewNotFoundError("No Customers Found")
		} else {
			logger.Error("Un Expected Database error" + err.Error())
			return nil, errors.NewUnExpectedError("Un Expected Database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	custmerQuery := "select * from customers where customer_id = $1 "
	var c Customer
	err := d.client.Get(&c, custmerQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customer found" + err.Error())
			return nil, errors.NewNotFoundError(`Customer Not Found`)
		} else {
			logger.Error("Error While scanning customer" + err.Error())
			return nil, errors.NewUnExpectedError(`Unexpected Database error`)
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	connStr := "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error opening database: %v\n" + err.Error())
	}
	err = client.Ping()
	if err != nil {
		logger.Error("Unable to connect to the database: %v\n" + err.Error())
	}
	logger.Info("Connected to PostgreSQL!")
	return CustomerRepositoryDb{client: client}
}
