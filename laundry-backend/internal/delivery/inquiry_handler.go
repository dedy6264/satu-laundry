package delivery

import (
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type InquiryHandler struct {
	inquiryUsecase usecases.InquiryUsecase
}

func NewInquiryHandler(inquiryUsecase usecases.InquiryUsecase) *InquiryHandler {
	return &InquiryHandler{
		inquiryUsecase: inquiryUsecase,
	}
}

func (h *InquiryHandler) ProcessInquiry(c echo.Context) error {

	var (
		request entities.InquiryRequest
		svcName = "ProcessInquiry"
	)
	if err := c.Bind(&request); err != nil {
		utils.LoggMsg(svcName, "Failed to bind request", err)
		return ErrorResponse(c, http.StatusBadRequest, "Invalid request format", err.Error())
	}
	// Ambil token dari context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	request.EmployeeID = int(claims["employee_id"].(float64)) // JSON number → float64 → int
	// role := claims["role"].(string)
	// if role != "employee" && role != "kasir" {
	// 	utils.LoggMsg(svcName, fmt.Sprintf("Unauthorized role: %s", role), nil)
	// 	return ErrorResponse(c, http.StatusForbidden, "Unauthorized role", "Only employees can process inquiries")
	// }
	response, err := h.inquiryUsecase.ProcessInquiry(request)
	if err != nil {
		fmt.Printf("Failed to process inquiry: %v\n", err)
		return ErrorResponse(c, http.StatusBadRequest, "Failed to process inquiry", err.Error())
	}

	fmt.Printf("=== PROCESS INQUIRY HANDLER END ===\n")
	return SuccessResponse(c, http.StatusOK, "Inquiry processed successfully", response)
}
