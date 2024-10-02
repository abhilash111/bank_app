package app

import (
	"encoding/json"
	"net/http"

	"github.com/abhilash111/bank_app/service"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := c.Service.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
