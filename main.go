package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhilash111/bank_app/app"
	_ "github.com/lib/pq"
)

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt string
}

func main() {
	connStr := "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	id := 1
	rows, err := db.Query("SELECT * FROM accounts WHERE id = $1", id)
	if err != nil {
		fmt.Errorf("Failed to retrieve accounts: %v", err)
	}
	defer rows.Close()

	var accounts []Account
	// Iterate over the rows and scan data into the Account struct
	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency, &account.CreatedAt)
		if err != nil {
			fmt.Errorf("Failed to scan account: %v", err)
		}
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		fmt.Errorf("Error in row iteration: %v", err)
	}
	for _, account := range accounts {
		fmt.Printf("ID: %d, Owner: %s, Balance: %d, Currency: %s, CreatedAt: %s\n",
			account.ID, account.Owner, account.Balance, account.Currency, account.CreatedAt)
	}
	app.Start()
}
