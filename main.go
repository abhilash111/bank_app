package main

import (
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
	app.Start()
}
