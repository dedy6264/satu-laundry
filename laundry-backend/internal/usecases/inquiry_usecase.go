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

func NewInquiryUsecase(inquiryRepo repositories.InquiryRepository,
	userAccessRepo repositories.UserAccessRepository,
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

func (u *inquiryUsecase) ProcessInquiry(request entities.InquiryRequest, claims jwt.MapClaims) (response entities.InquiryResponse, err error) {
	var (
		t          = time.Now()
		outletId   int
		brandId    int
		detail     entities.TransactionDetail
		details    []entities.TransactionDetail
		grandTotal float64
	)

	// 1. validasi user access
	userAccess, err := u.userAccessRepo.FindByID(request.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, errors.New("invalid userAccess")
		}
		return response, err
	}
	// 2. get outlet
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "pegawai":
			employee, err := u.employeeRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				return response, err
			}
			outletId = employee.OutletID
			brandId = employee.BrandID
		case "outlet":
			outlet, err := u.outletRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				return response, err
			}
			outletId = outlet.ID
			brandId = outlet.BrandID
		default:
			return response, errors.New("Invalid Reference Level")
		}
	} else {
		if request.OutletID == 0 {
			return response, errors.New("Outlet ID CAnnot be Null")
		}
		//validasi outlet
		outletArr, err := u.outletRepo.FindAll(entities.Outlet{
			CabangID: userAccess.ReferenceID,
			ID:       request.OutletID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return response, errors.New("invalid Package")
			}
			return response, err
		}
		if len(outletArr) == 0 {
			return response, errors.New("invalid OutletID")
		}
		//get brands
		outlet, err := u.outletRepo.FindByID(request.OutletID)
		if err != nil {
			return response, err
		}
		outletId = outlet.ID

		brandId = outlet.BrandID
	}
	// 3. Validasi paket layanan
	servicePackage, err := u.serviceRepo.FindAll(entities.Service{
		BrandID: brandId,
		ID:      request.ServicePackageID,
	}) //find apakah paket itu tersedia untuk brand kita?

	if err != nil {
		if err == sql.ErrNoRows {
			return response, errors.New("invalid Package")
		}
		return response, err
	}
	for _, data := range request.Product {
		for _, v := range servicePackage {
			if data.ServicePackageID == v.ID {
				grandTotal = grandTotal + (v.Price * data.Quantity)
			}
		}
	}
	// 4. Validate customer
	valid, err := u.inquiryRepo.ValidateCustomer(request.CustomerID)
	if err != nil {
		return response, err
	}
	if !valid {
		return response, errors.New("invalid customer")
	}
	// 5. validasi payment method
	paymentMethod, err := u.paymentRepo.FindByID(request.PaymentMethodID)
	if err != nil {
		return response, err
	}

	// Begin database transaction
	tx, err := u.inquiryRepo.BeginTransaction()
	if err != nil {
		return response, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create transaction entity
	if request.OutletID == 0 {
		request.OutletID = outletId
	}
	transaction := entities.Transaction{
		CustomerID:    request.CustomerID,
		OutletID:      request.OutletID,
		InvoiceNumber: generateInvoiceNumber(),
		EntryDate:     t,
		Status:        "diterima", // Default status
		Note:          request.Note,
		CreatedAt:     t,
		UpdatedAt:     t,
		CreatedBy:     userAccess.Username,
		UpdatedBy:     userAccess.Username,
		UserID:        userAccess.ID,
		TotalPrice:    grandTotal,
	}

	//1. Insert transaction with transaction
	id, err := u.inquiryRepo.InsertTransactionWithTx(tx, transaction)
	if err != nil {
		tx.Rollback()
		return response, fmt.Errorf("failed to insert transaction: %w", err)
	}

	for _, data := range request.Product {
		for _, v := range servicePackage {
			if data.ServicePackageID == v.ID {
				detail = entities.TransactionDetail{
					TransactionID: id,
					ServiceID:     v.ID,
					Quantity:      data.Quantity,
					Price:         v.Price,
					Subtotal:      v.Price * data.Quantity,
					CreatedAt:     t,
					UpdatedAt:     t,
					CreatedBy:     userAccess.Username,
					UpdatedBy:     userAccess.Username,
				}
				details = append(details, detail)
				request.Quantity = request.Quantity + data.Quantity
			}
		}
	}
	if len(details) != len(request.Product) {
		tx.Rollback()
		return response, errors.New("invalid Transaction")
	}

	//2. Insert transaction detail with transaction
	err = u.inquiryRepo.InsertTransactionDetailWithTx(tx, details)
	if err != nil {
		tx.Rollback()
		return response, fmt.Errorf("failed to insert transaction detail: %w", err)
	}

	// Create initial payment record with default values
	payment := entities.Payment{
		TransactionID:   id,
		PaymentMethodID: paymentMethod.ID,
		PaymentDate:     t,
		Amount:          0,                        // Default to 0 as no payment has been made yet
		Method:          paymentMethod.NamaMetode, // Default to empty as no payment method selected yet
		CreatedAt:       t,
		UpdatedAt:       t,
	}

	// Insert payment record
	err = u.inquiryRepo.InsertPaymentWithTx(tx, payment)
	if err != nil {
		tx.Rollback()
		return response, fmt.Errorf("failed to insert payment: %w", err)
	}

	// Create initial history status transaction record
	history := entities.HistoryStatusTransaction{
		TransactionID: id,
		OldStatus:     "diterima", // No old status as this is the initial status
		NewStatus:     "diterima",
		ChangeTime:    t,
		Description:   "Transaksi baru dibuat",
		CreatedAt:     t,
		UpdatedAt:     t,
	}

	// Insert history status transaction record
	err = u.inquiryRepo.InsertHistoryStatusTransactionWithTx(tx, history)
	if err != nil {
		tx.Rollback()
		return response, fmt.Errorf("failed to insert history status transaction: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return response, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Prepare the response
	response = entities.InquiryResponse{
		Transaction:        transaction,
		TransactionDetails: details, // Use slice of pointers, not dereferenced
		Payment:            payment,
		History:            history,
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
