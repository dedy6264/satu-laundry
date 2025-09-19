package usecases

import (
	"errors"
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type inquiryUsecase struct {
	inquiryRepo    repositories.InquiryRepository
	userAccessRepo repositories.UserAccessRepository
	cabangRepo     repositories.CabangRepository
	outletRepo     repositories.OutletRepository
	employeeRepo   repositories.EmployeeRepository
}

func NewInquiryUsecase(inquiryRepo repositories.InquiryRepository, userAccessRepo repositories.UserAccessRepository,
	cabangRepo repositories.CabangRepository,
	outletRepo repositories.OutletRepository,
	employeeRepo repositories.EmployeeRepository) InquiryUsecase {
	return &inquiryUsecase{
		inquiryRepo:    inquiryRepo,
		userAccessRepo: userAccessRepo,
		cabangRepo:     cabangRepo,
		outletRepo:     outletRepo,
		employeeRepo:   employeeRepo,
	}
}

func (u *inquiryUsecase) ProcessInquiry(request entities.InquiryRequest, claims jwt.MapClaims) (response *entities.InquiryResponse, err error) {
	var (
		t  = time.Now()
		id int
		// tdb = t.Local().Format(time.RFC3339)
		// cabang   *entities.Cabang cabang ga bisa dipake buat trx karena harus ada outlet
		outlet   *entities.Outlet
		employee *entities.Employee
	)
	reference_id := int(claims["reference_id"].(float64))
	reference_level := claims["reference_level"].(string)
	// role := claims["role"].(string)
	switch reference_level {
	// case "cabang":
	// 	cabang, err = u.cabangRepo.FindByID(reference_id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	id = cabang.ID
	case "outlet":
		outlet, err = u.outletRepo.FindByID(reference_id)
		if err != nil {
			return nil, err
		}
		id = outlet.ID
	default: //karyawan
		employee, err = u.employeeRepo.FindByID(reference_id)
		if err != nil {
			return nil, err
		}
		id = employee.ID
	}
	// Validate service package
	valid, err := u.inquiryRepo.ValidateServicePackage(request.ServicePackageID)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("invalid service package")
	}

	userAccess, err := u.userAccessRepo.FindByID(request.UserID)
	if err != nil {
		return nil, err
	}
	if userAccess == nil {
		return nil, errors.New("invalid userAccess")
	}

	// Validate customer
	valid, err = u.inquiryRepo.ValidateCustomer(request.CustomerID)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("invalid customer")
	}

	// Get service package price
	price, err := u.inquiryRepo.GetServicePackagePrice(request.ServicePackageID)
	if err != nil {
		return nil, err
	}

	// Calculate subtotal
	subtotal := price * request.Quantity

	// Begin database transaction
	tx, err := u.inquiryRepo.BeginTransaction()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create transaction entity
	if request.OutletID == 0 {
		request.OutletID = employee.OutletID
	}
	transaction := &entities.Transaction{
		CustomerID:    request.CustomerID,
		OutletID:      request.OutletID,
		InvoiceNumber: generateInvoiceNumber(),
		EntryDate:     &t,
		Status:        "diterima", // Default status
		Note:          request.Note,
		CreatedAt:     t,
		UpdatedAt:     t,
		CreatedBy:     &userAccess.Username,
		UpdatedBy:     &userAccess.Username,
		UserID:        &userAccess.ID,
		TotalPrice:    subtotal,
	}

	// Insert transaction with transaction
	id, err = u.inquiryRepo.InsertTransactionWithTx(tx, transaction)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert transaction: %w", err)
	}

	// Create transaction detail
	detail := &entities.TransactionDetail{
		TransactionID: id,
		ServiceID:     request.ServicePackageID,
		Quantity:      &request.Quantity,
		Price:         &price,
		Subtotal:      &subtotal,
		CreatedAt:     t,
		UpdatedAt:     t,
		CreatedBy:     &employee.Name,
		UpdatedBy:     &employee.Name,
	}

	// Insert transaction detail with transaction
	err = u.inquiryRepo.InsertTransactionDetailWithTx(tx, detail)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert transaction detail: %w", err)
	}

	// Create initial payment record with default values
	payment := &entities.Payment{
		TransactionID: id,
		PaymentDate:   &t,
		Amount:        0,  // Default to 0 as no payment has been made yet
		Method:        "", // Default to empty as no payment method selected yet
		CreatedAt:     t,
		UpdatedAt:     t,
	}

	// Insert payment record
	err = u.inquiryRepo.InsertPaymentWithTx(tx, payment)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert payment: %w", err)
	}

	// Create initial history status transaction record
	history := &entities.HistoryStatusTransaction{
		TransactionID: id,
		OldStatus:     "diterima", // No old status as this is the initial status
		NewStatus:     "diterima",
		ChangeTime:    &t,
		Description:   "Transaksi baru dibuat",
		CreatedAt:     t,
		UpdatedAt:     t,
	}

	// Insert history status transaction record
	err = u.inquiryRepo.InsertHistoryStatusTransactionWithTx(tx, history)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert history status transaction: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Prepare the response
	response = &entities.InquiryResponse{
		Transaction:        *transaction,
		TransactionDetails: []entities.TransactionDetail{*detail},
		Payment:            *payment,
		History:            *history,
	}

	return response, nil
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

// stringPtr returns a pointer to the given string
func stringPtr(s string) *string {
	return &s
}

// intPtr returns a pointer to the given int
func intPtr(i int) *int {
	return &i
}

// timePtr returns a pointer to the given time
func timePtr(t time.Time) *time.Time {
	return &t
}
