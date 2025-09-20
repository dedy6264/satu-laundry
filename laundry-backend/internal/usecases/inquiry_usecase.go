package usecases

import (
	"database/sql"
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
	paymentRepo    repositories.PaymentMethodRepository
	serviceRepo    repositories.ServiceRepository
}

func NewInquiryUsecase(inquiryRepo repositories.InquiryRepository, userAccessRepo repositories.UserAccessRepository,
	cabangRepo repositories.CabangRepository,
	outletRepo repositories.OutletRepository,
	employeeRepo repositories.EmployeeRepository,
	paymentRepo repositories.PaymentMethodRepository,
	serviceRepo repositories.ServiceRepository) InquiryUsecase {
	return &inquiryUsecase{
		inquiryRepo:    inquiryRepo,
		userAccessRepo: userAccessRepo,
		cabangRepo:     cabangRepo,
		outletRepo:     outletRepo,
		employeeRepo:   employeeRepo,
		paymentRepo:    paymentRepo,
		serviceRepo:    serviceRepo,
	}
}

func (u *inquiryUsecase) ProcessInquiry(request entities.InquiryRequest, claims jwt.MapClaims) (response *entities.InquiryResponse, err error) {
	var (
		t        = time.Now()
		outlerId int
	)
	// 1. validasi user access
	userAccess, err := u.userAccessRepo.FindByID(request.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid userAccess")
		}
		return nil, err
	}
	// 2. get outlet
	fmt.Println(":::", userAccess.ReferenceLevel)
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "karyawan":
			employee, err := u.employeeRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				return nil, err
			}
			outlerId = employee.OutletID
		case "outlet":
			outlet, err := u.outletRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				return nil, err
			}
			outlerId = outlet.ID
		default:
			return nil, errors.New("Invalid Reference Level")
		}
	} else {
		if request.OutletID == 0 {
			return nil, errors.New("Outlet ID CAnnot be Null")
		}
		//validasi outlet
		outletArr, err := u.outletRepo.FindAll(entities.Outlet{
			CabangID: userAccess.ReferenceID,
			ID:       request.OutletID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("invalid Package")
			}
			return nil, err
		}
		if len(outletArr) == 0 {
			return nil, errors.New("invalid OutletID")
		}
	}
	// 3. Validasi paket layanan
	servicePackage, err := u.serviceRepo.FindByID(request.ServicePackageID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid Package")
		}
		return nil, err
	}

	// 4. Validate customer
	valid, err := u.inquiryRepo.ValidateCustomer(request.CustomerID)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("invalid customer")
	}
	// 5. validasi payment method
	paymentMethod, err := u.paymentRepo.FindByID(request.PaymentMethodID)
	if err != nil {
		return nil, err
	}

	// Calculate subtotal
	subtotal := servicePackage.Price * request.Quantity

	// Begin database transaction
	tx, err := u.inquiryRepo.BeginTransaction()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create transaction entity
	if request.OutletID == 0 {
		request.OutletID = outlerId
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
	id, err := u.inquiryRepo.InsertTransactionWithTx(tx, transaction)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert transaction: %w", err)
	}

	// Create transaction detail
	detail := &entities.TransactionDetail{
		TransactionID: id,
		ServiceID:     request.ServicePackageID,
		Quantity:      &request.Quantity,
		Price:         &servicePackage.Price,
		Subtotal:      &subtotal,
		CreatedAt:     t,
		UpdatedAt:     t,
		CreatedBy:     &userAccess.Username,
		UpdatedBy:     &userAccess.Username,
	}

	// Insert transaction detail with transaction
	err = u.inquiryRepo.InsertTransactionDetailWithTx(tx, detail)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert transaction detail: %w", err)
	}

	// Create initial payment record with default values
	payment := &entities.Payment{
		TransactionID:   id,
		PaymentMethodID: paymentMethod.ID,
		PaymentDate:     &t,
		Amount:          0,                        // Default to 0 as no payment has been made yet
		Method:          paymentMethod.NamaMetode, // Default to empty as no payment method selected yet
		CreatedAt:       t,
		UpdatedAt:       t,
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
