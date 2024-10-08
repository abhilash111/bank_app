package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abhilash111/bank_app/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := c.Service.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (c *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := c.Service.GetCustomer(id)
	if err != nil {
		fmt.Println("Err", err.Code)
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
