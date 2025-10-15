package transactionservice

import (
	"laundry-backend/services"
)

type TransactionService struct {
	service services.UsecaseService
}

func ApiTransactionService(service services.UsecaseService) TransactionService {
	return TransactionService{
		service: service,
	}
}
