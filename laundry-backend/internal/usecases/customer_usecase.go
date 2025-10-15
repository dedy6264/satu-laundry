package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type customerUsecase struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerUsecase(customerRepo repositories.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (u *customerUsecase) CreateCustomer(request entities.RegisterCustomerRequest) error {
	customer := &entities.Customer{
		OutletID: request.OutletID,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	return u.customerRepo.Create(customer)
}

func (u *customerUsecase) GetCustomerByID(id int) (*entities.Customer, error) {
	return u.customerRepo.FindByID(id)
}

func (u *customerUsecase) GetCustomersByOutletID(outletID int) ([]entities.Customer, error) {
	return u.customerRepo.FindByOutletID(outletID)
}

func (u *customerUsecase) GetAllCustomers() ([]entities.Customer, error) {
	return u.customerRepo.FindAll()
}

func (u *customerUsecase) GetAllCustomersDataTables(request entities.DTRequest[entities.Customer]) (*entities.DataTablesResponse, error) {

	// Get data with pagination
	customers, recordsTotal, err := u.customerRepo.FindAllWithPagination(
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
		Data:            customers,
	}

	return response, nil
}

func (u *customerUsecase) UpdateCustomer(id int, request entities.RegisterCustomerRequest) error {
	customer, err := u.customerRepo.FindByID(id)
	if err != nil {
		return err
	}

	if customer == nil {
		return nil // Customer not found
	}

	customer.OutletID = request.OutletID
	customer.Name = request.Name
	customer.Email = request.Email
	customer.Phone = request.Phone
	customer.Address = request.Address

	return u.customerRepo.Update(customer)
}

func (u *customerUsecase) DeleteCustomer(id int) error {
	return u.customerRepo.Delete(id)
}
