package service

import (
	"github.com/abhilash111/bank_app/domain"
	"github.com/abhilash111/bank_app/errors"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errors.AppError)
}

type CustomerRepoInjector struct {
	repo domain.CustomerRepository
}

func (c CustomerRepoInjector) GetAllCustomers() ([]domain.Customer, error) {
	return c.repo.FindAll()
}

func (c CustomerRepoInjector) GetCustomer(id string) (*domain.Customer, *errors.AppError) {
	return c.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) CustomerRepoInjector {
	return CustomerRepoInjector{repo: repository}
}
