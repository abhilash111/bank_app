package service

import "github.com/abhilash111/bank_app/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type CustomerRepoInjector struct {
	repo domain.CustomerRepository
}

func (c CustomerRepoInjector) GetAllCustomers() ([]domain.Customer, error) {
	return c.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) CustomerRepoInjector {
	return CustomerRepoInjector{repo: repository}
}
