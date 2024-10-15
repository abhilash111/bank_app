package app

import (
	"encoding/json"
	"net/http"

	"github.com/abhilash111/bank_app/dto"
	"github.com/abhilash111/bank_app/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service service.AccountService
}

func (a *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var requestDto dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&requestDto)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		requestDto.CustomerId = customerId
		account, err := a.Service.NewAccount(requestDto)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.Service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
