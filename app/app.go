package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhilash111/bank_app/domain"
	"github.com/abhilash111/bank_app/service"
	"github.com/gorilla/mux"
)

func Start() {
	fmt.Printf("Bank App Started Successfully at localhost:8000...")
	router := mux.NewRouter()
	//wiring
	ch := CustomerHandler{Service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
