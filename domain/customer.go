package domain

type Customer struct {
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
