package domain

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhilash111/bank_app/errors"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {
	rows, err := c.client.Query("SELECT * FROM customers")
	if err != nil {
		log.Fatal("Failed to LOAD Customers")
	}
	defer rows.Close()
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Zipcode, &c.Status)
		if err != nil {
			log.Fatal("Failed to SCAN Customer")
		}
		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		fmt.Errorf("Error in row iteration: %v", err)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	custmerQuery := "select * from customers where customer_id = $1 "
	row := d.client.QueryRow(custmerQuery, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Err", err)
			return nil, errors.NewNotFoundError(`Customer Not Found`)
		} else {
			log.Println("Error While scanning customer" + err.Error())
			return nil, errors.NewUnExpectedError(`Unexpected Database error`)
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	connStr := "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	err = client.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	return CustomerRepositoryDb{client: client}
}
