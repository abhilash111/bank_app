package app

import (
	"encoding/json"
	"net/http"

	"github.com/abhilash111/bank_app/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.Service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (c *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := c.Service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
