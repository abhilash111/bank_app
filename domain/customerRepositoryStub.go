package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			"201", "Abhilash", "Bidar", "23/04/2999", "560990", "1",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
