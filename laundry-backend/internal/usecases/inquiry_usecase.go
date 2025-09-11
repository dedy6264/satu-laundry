package usecases

import (
	"errors"
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"math/rand"
	"time"
)

type inquiryUsecase struct {
	inquiryRepo repositories.InquiryRepository
}

func NewInquiryUsecase(inquiryRepo repositories.InquiryRepository) InquiryUsecase {
	return &inquiryUsecase{
		inquiryRepo: inquiryRepo,
	}
}

func (u *inquiryUsecase) ProcessInquiry(request entities.InquiryRequest) error {
	// Validate service package
	valid, err := u.inquiryRepo.ValidateServicePackage(request.ServicePackageID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid service package")
	}

	// Validate employee and get employee data
	valid, err = u.inquiryRepo.ValidateEmployee(request.EmployeeID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid employee")
	}

	// Validate customer
	valid, err = u.inquiryRepo.ValidateCustomer(request.CustomerID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid customer")
	}

	// Get service package price
	price, err := u.inquiryRepo.GetServicePackagePrice(request.ServicePackageID)
	if err != nil {
		return err
	}

	// Calculate subtotal
	subtotal := price * request.Quantity

	// Create transaction entity
	transaction := &entities.Transaction{
		CustomerID:    request.CustomerID,
		OutletID:      request.OutletID,
		InvoiceNumber: generateInvoiceNumber(),
		EntryDate:     time.Now(),
		Status:        "baru", // Default status
		TotalCost:     subtotal,
		Note:          request.Note,
	}

	// Insert transaction
	err = u.inquiryRepo.InsertTransaction(transaction)
	if err != nil {
		return err
	}

	// Create transaction detail
	detail := &entities.TransactionDetail{
		TransactionID: transaction.ID,
		ServiceID:     request.ServicePackageID,
		Quantity:      request.Quantity,
		Price:         price,
		Subtotal:      subtotal,
	}

	// Insert transaction detail
	err = u.inquiryRepo.InsertTransactionDetail(detail)
	if err != nil {
		return err
	}

	return nil
}

// generateInvoiceNumber generates a unique invoice number
func generateInvoiceNumber() string {
	// Use current timestamp and random number to generate unique invoice number
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	// Generate a random 4-digit number
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(9000) + 1000

	return fmt.Sprintf("INV%d%02d%02d%02d%02d%02d%d", year, month, day, hour, minute, second, random)
}
