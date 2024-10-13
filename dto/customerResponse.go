package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	DateofBirth string `json:"date_of_birth"`
	Zipcode     string `json:"zipcode"`
	Status      string `json:"status"`
}
