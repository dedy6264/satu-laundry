package inquiryservice

import (
	"database/sql"
	"fmt"
	"laundry-backend/entities"
	"laundry-backend/utils"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h InquiryService) ProcessInquiry(c echo.Context) error {

	var (
		outletId, brandId int
		grandTotal        float64
		request           entities.InquiryRequest
		svcName           = "ProcessInquiry"
		t                 = time.Now()
		detail            entities.TransactionDetail
		details           []entities.TransactionDetail
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	// Ambil token dari context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	request.UserID = int(claims["user_id"].(float64)) // JSON number → float64 → int
	if request.PaymentMethodID == 0 {
		utils.LoggMsg(svcName, "Invalid Payment Method", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Payment Method", "")
	}

	// 1. validasi user access
	userAccess, err := h.service.RepoUserAccess.FindByID(request.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.LoggMsg(svcName, "invalid userAccess", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid userAccess", "")
		}
		utils.LoggMsg(svcName, "Invalid:: FindById", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid:: FindById", "")
	}
	// 2. get outlet
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "pegawai":
			employee, err := h.service.RepoEmployee.FindByID(userAccess.ReferenceID)
			if err != nil {
				utils.LoggMsg(svcName, "Invalid:: FindById Pegawai", nil)
				return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid:: FindById Pegawai", "")
			}
			outletId = employee.OutletID
			brandId = employee.BrandID
		case "outlet":
			outlet, err := h.service.RepoOutlet.FindByID(userAccess.ReferenceID)
			if err != nil {
				utils.LoggMsg(svcName, "Invalid:: FindById outlet", nil)
				return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid:: FindById outlet", "")
			}
			outletId = outlet.ID
			brandId = outlet.BrandID
		default:
			utils.LoggMsg(svcName, "Invalid Reference Level", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Reference Level", "")
		}
	} else {
		if request.OutletID == 0 {
			utils.LoggMsg(svcName, "Outlet ID CAnnot be Null", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "Outlet ID CAnnot be Null", "")
		}
		//validasi outlet
		outletArr, err := h.service.RepoOutlet.FindAll(entities.Outlet{
			CabangID: userAccess.ReferenceID,
			ID:       request.OutletID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				utils.LoggMsg(svcName, "invalid Package", nil)
				return utils.ErrorResponse(c, http.StatusBadRequest, "invalid Package", "")
			}
			utils.LoggMsg(svcName, "Invalid ", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ", "")
		}
		if len(outletArr) == 0 {
			utils.LoggMsg(svcName, "invalid OutletID", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "invalid OutletID", "")
		}
		//get brands
		outlet, err := h.service.RepoOutlet.FindByID(request.OutletID)
		if err != nil {
			utils.LoggMsg(svcName, "Invalid ", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ", "")
		}
		outletId = outlet.ID

		brandId = outlet.BrandID
	}
	// 3. Validasi paket layanan
	servicePackage, err := h.service.RepoService.FindAll(entities.Service{
		BrandID: brandId,
		ID:      request.ServicePackageID,
	}) //find apakah paket itu tersedia untuk brand kita?

	if err != nil {
		if err == sql.ErrNoRows {
			utils.LoggMsg(svcName, "Invalid Package", nil)
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Package", "")
		}
		utils.LoggMsg(svcName, "Invalid Package", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid Package", "")
	}
	for _, data := range request.Product {
		for _, v := range servicePackage {
			if data.ServicePackageID == v.ID {
				grandTotal = grandTotal + (v.Price * data.Quantity)
			}
		}
	}
	// 4. Validate customer
	valid, err := h.service.RepoInquiry.ValidateCustomer(request.CustomerID)
	if err != nil {
		utils.LoggMsg(svcName, "Invalid :: Validate Customer", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid :: Validate Customer", "")
	}
	if !valid {
		utils.LoggMsg(svcName, "invalid customer", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid customer", "")
	}
	// 5. validasi payment method
	paymentMethod, err := h.service.RepoPaymentMethod.FindByID(request.PaymentMethodID)
	if err != nil {
		utils.LoggMsg(svcName, "Invalid :: Find By ID Payment", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid :: Find By ID Payment", "")
	}

	// Begin database transaction
	tx, err := h.service.RepoInquiry.BeginTransaction()
	if err != nil {
		utils.LoggMsg(svcName, "failed to begin transaction", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to begin transaction", "")
	}

	// Create transaction entity
	if request.OutletID == 0 {
		request.OutletID = outletId
	}
	transaction := entities.Transaction{
		CustomerID:    request.CustomerID,
		OutletID:      request.OutletID,
		InvoiceNumber: utils.GenerateInvoiceNumber(),
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
	id, err := h.service.RepoInquiry.InsertTransactionWithTx(tx, transaction)
	if err != nil {
		tx.Rollback()
		utils.LoggMsg(svcName, "failed to insert transaction", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to insert transaction", "")
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
		utils.LoggMsg(svcName, "invalid Transaction", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "invalid Transaction", "")
	}

	//2. Insert transaction detail with transaction
	err = h.service.RepoInquiry.InsertTransactionDetailWithTx(tx, details)
	if err != nil {
		tx.Rollback()
		utils.LoggMsg(svcName, "failed to insert transaction detail", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to insert transaction detail", "")
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
	err = h.service.RepoInquiry.InsertPaymentWithTx(tx, payment)
	if err != nil {
		tx.Rollback()
		utils.LoggMsg(svcName, "failed to insert payment", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to insert payment", "")
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
	err = h.service.RepoInquiry.InsertHistoryStatusTransactionWithTx(tx, history)
	if err != nil {
		tx.Rollback()
		utils.LoggMsg(svcName, "failed to insert history status transaction", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to insert history status transaction", "")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		utils.LoggMsg(svcName, "failed to commit transaction", nil)
		return utils.ErrorResponse(c, http.StatusBadRequest, "failed to commit transaction", "")
	}

	// Prepare the response
	response := entities.InquiryResponse{
		Transaction:        transaction,
		TransactionDetails: details, // Use slice of pointers, not dereferenced
		Payment:            payment,
		History:            history,
	}

	fmt.Printf("=== PROCESS INQUIRY HANDLER END ===\n")
	return utils.SuccessResponse(c, http.StatusOK, "Inquiry processed successfully", response)
}
