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
	var (
		t = time.Now()
		// tdb = t.Local().Format(time.RFC3339)
	)
	// Validate service package
	valid, err := u.inquiryRepo.ValidateServicePackage(request.ServicePackageID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid service package")
	}

	// Validate employee and get employee data

	employee, err := u.inquiryRepo.ValidateEmployee(request.EmployeeID)
	if err != nil {
		return err
	}
	if employee == nil {
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

	// Begin database transaction
	tx, err := u.inquiryRepo.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create transaction entity
	// Use employee's outlet ID instead of request.OutletID
	transaction := &entities.Transaction{
		CustomerID:    request.CustomerID,
		OutletID:      employee.OutletID,
		InvoiceNumber: generateInvoiceNumber(),
		EntryDate:     t,
		Status:        "diterima", // Default status
		Note:          request.Note,

		CreatedAt:              t,
		UpdatedAt:              t,
		CreatedBy:              employee.Name,
		UpdatedBy:              employee.Name,
		EmployeeID:             employee.ID,
		TotalPrice:             subtotal,
		PaymentStatus:          "belum lunas",
		PaymentMethod:          "tunai",
		StatusCode:             "009",
		StatusMessage:          "INQUIRY SUCCESS",
		PaymentReferenceNumber: "",
	}

	// Insert transaction with transaction
	id, err := u.inquiryRepo.InsertTransactionWithTx(tx, transaction)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	// Create transaction detail
	detail := &entities.TransactionDetail{
		TransactionID: id,
		ServiceID:     request.ServicePackageID,
		Quantity:      request.Quantity,
		Price:         price,
		Subtotal:      subtotal,
		CreatedAt:     t,
		UpdatedAt:     t,
		CreatedBy:     employee.Name,
		UpdatedBy:     employee.Name,
	}

	// Insert transaction detail with transaction
	err = u.inquiryRepo.InsertTransactionDetailWithTx(tx, detail)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert transaction detail: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
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
