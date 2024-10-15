package app

import (
	"log"
	"net/http"

	"github.com/abhilash111/bank_app/domain"
	"github.com/abhilash111/bank_app/logger"
	"github.com/abhilash111/bank_app/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	logger.Info("Bank App Started Successfully at localhost:8000...")
	router := mux.NewRouter()
	dbClient := getDbClient()
	//wiring
	ch := CustomerHandler{Service: service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AccountHandler{Service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	connStr := "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error opening database: %v\n" + err.Error())
	}
	err = client.Ping()
	if err != nil {
		logger.Error("Unable to connect to the database: %v\n" + err.Error())
	}
	logger.Info("Connected to PostgreSQL!")
	return client
}
