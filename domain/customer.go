package domain

import "github.com/abhilash111/bank_app/errors"

type Customer struct {
	Id          string
	Name        string
	City        string
	DateofBirth string
	Zipcode     string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	ById(string) (*Customer, *errors.AppError)
}
