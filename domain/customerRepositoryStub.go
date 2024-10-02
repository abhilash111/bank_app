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
			"Abhilash", "Bidar", "560990", "23/04/2999",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
