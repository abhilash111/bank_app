package service

import (
	"github.com/abhilash111/bank_app/domain"
	"github.com/abhilash111/bank_app/dto"
	"github.com/abhilash111/bank_app/errors"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

type CustomerRepoInjector struct {
	repo domain.CustomerRepository
}

func (c CustomerRepoInjector) GetAllCustomers() ([]domain.Customer, *errors.AppError) {
	return c.repo.FindAll()
}

func (s CustomerRepoInjector) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) CustomerRepoInjector {
	return CustomerRepoInjector{repo: repository}
}
