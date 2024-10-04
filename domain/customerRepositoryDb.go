package domain

import (
	"database/sql"
	"fmt"
	"log"
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

func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	custmerQuery := "select * from customers where customer_id = $1 "
	row := d.client.QueryRow(custmerQuery, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Zipcode, &c.Status)
	if err != nil {
		fmt.Println("Err", err)
		return nil, fmt.Errorf("Customer with id %s not Found", id)
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
