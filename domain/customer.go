package domain

import (
	"github.com/abhilash111/bank_app/dto"
	"github.com/abhilash111/bank_app/errors"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	DateofBirth string `db:"date_of_birth"`
	Zipcode     string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	ById(string) (*Customer, *errors.AppError)
}

func (c *Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.Status,
		Status:      c.Status,
	}
}
