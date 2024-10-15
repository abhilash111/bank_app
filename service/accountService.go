package service

import (
	"time"

	"github.com/abhilash111/bank_app/domain"
	"github.com/abhilash111/bank_app/dto"
	"github.com/abhilash111/bank_app/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.AccountRequest) (*domain.Account, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.AccountRequest) (*domain.Account, *errs.AppError) {
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo,
	}
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewUnExpectedError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil

}
