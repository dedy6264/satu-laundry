package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type paymentMethodUsecase struct {
	paymentMethodRepo repositories.PaymentMethodRepository
}

func NewPaymentMethodUsecase(paymentMethodRepo repositories.PaymentMethodRepository) PaymentMethodUsecase {
	return &paymentMethodUsecase{
		paymentMethodRepo: paymentMethodRepo,
	}
}

func (u *paymentMethodUsecase) CreatePaymentMethod(request entities.CreatePaymentMethodRequest) error {
	paymentMethod := &entities.PaymentMethod{
		NamaMetode:  request.NamaMetode,
		URL:         request.URL,
		SKey:        request.SKey,
		MKey:        request.MKey,
		MerchantFee: request.MerchantFee,
		AdminFee:    request.AdminFee,
		Status:      request.Status,
		CreatedBy:   request.CreatedBy,
		UpdatedBy:   request.UpdatedBy,
	}

	return u.paymentMethodRepo.Create(paymentMethod)
}

func (u *paymentMethodUsecase) GetPaymentMethodByID(id int) (*entities.PaymentMethod, error) {
	return u.paymentMethodRepo.FindByID(id)
}

func (u *paymentMethodUsecase) GetAllPaymentMethods() ([]entities.PaymentMethod, error) {
	return u.paymentMethodRepo.FindAll()
}

func (u *paymentMethodUsecase) GetAllPaymentMethodsDataTables(request entities.DTRequest[entities.PaymentMethod]) (*entities.DataTablesResponse, error) {

	// Get data with pagination
	paymentMethods, recordsTotal, err := u.paymentMethodRepo.FindAllWithPagination(
		request,
	)

	if err != nil {
		return nil, err
	}

	// Create response
	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsTotal,
		Data:            paymentMethods,
	}

	return response, nil
}

func (u *paymentMethodUsecase) UpdatePaymentMethod(id int, request entities.UpdatePaymentMethodRequest) error {
	paymentMethod, err := u.paymentMethodRepo.FindByID(id)
	if err != nil {
		return err
	}

	if paymentMethod == nil {
		return nil // Payment method not found
	}

	paymentMethod.NamaMetode = request.NamaMetode
	paymentMethod.URL = request.URL
	paymentMethod.SKey = request.SKey
	paymentMethod.MKey = request.MKey
	paymentMethod.MerchantFee = request.MerchantFee
	paymentMethod.AdminFee = request.AdminFee
	paymentMethod.Status = request.Status
	paymentMethod.UpdatedBy = request.UpdatedBy

	return u.paymentMethodRepo.Update(paymentMethod)
}

func (u *paymentMethodUsecase) DeletePaymentMethod(id int) error {
	return u.paymentMethodRepo.Delete(id)
}
