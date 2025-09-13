package usecases

import (
	"fmt"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"laundry-backend/internal/utils"
	"time"
)

type employeeAccessUsecase struct {
	employeeAccessRepo repositories.EmployeeAccessRepository
	jwtSecret          string
	tokenExpiry        time.Duration
}

func NewEmployeeAccessUsecase(
	employeeAccessRepo repositories.EmployeeAccessRepository,
	jwtSecret string,
	tokenExpiry time.Duration,
) EmployeeAccessUsecase {
	return &employeeAccessUsecase{
		employeeAccessRepo: employeeAccessRepo,
		jwtSecret:          jwtSecret,
		tokenExpiry:        tokenExpiry,
	}
}

func (u *employeeAccessUsecase) CreateEmployeeAccess(request entities.CreateEmployeeAccessRequest) error {
	access := &entities.EmployeeAccess{
		EmployeeID: request.EmployeeID,
		Username:   request.Username,
		Password:   request.Password,
		Role:       request.Role,
		IsActive:   request.IsActive,
	}
	return u.employeeAccessRepo.Create(access)
}

func (u *employeeAccessUsecase) GetEmployeeAccessByID(id int) (*entities.EmployeeAccess, error) {
	return u.employeeAccessRepo.FindByID(id)
}

func (u *employeeAccessUsecase) GetAllEmployeeAccess() ([]entities.EmployeeAccess, error) {
	return u.employeeAccessRepo.FindAll()
}

func (u *employeeAccessUsecase) GetAllEmployeeAccessDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// For simplicity, we'll use default pagination values
	// In a real implementation, you would parse these from the request
	limit := request.Length
	if limit <= 0 {
		limit = 10
	}

	offset := request.Start
	if offset < 0 {
		offset = 0
	}

	accesses, totalCount, err := u.employeeAccessRepo.FindAllWithPagination(limit, offset)
	if err != nil {
		return nil, err
	}

	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            accesses,
	}

	return response, nil
}

func (u *employeeAccessUsecase) UpdateEmployeeAccess(id int, request entities.UpdateEmployeeAccessRequest) error {
	access, err := u.employeeAccessRepo.FindByID(id)
	if err != nil {
		return err
	}

	if access == nil {
		return nil // Employee access not found
	}

	access.Username = request.Username
	access.Role = request.Role
	access.IsActive = request.IsActive

	return u.employeeAccessRepo.Update(access)
}

func (u *employeeAccessUsecase) UpdateEmployeePassword(id int, request entities.UpdateEmployeePasswordRequest) error {
	// First, get the employee access to verify current password
	access, err := u.employeeAccessRepo.FindByID(id)
	if err != nil {
		return err
	}

	if access == nil {
		return nil // Employee access not found
	}

	// Authenticate with current password
	authAccess, err := u.employeeAccessRepo.AuthenticateEmployee(access.Username, request.CurrentPassword)
	if err != nil || authAccess == nil {
		return err // Current password is incorrect
	}

	// Update with new password
	return u.employeeAccessRepo.UpdatePassword(id, request.NewPassword)
}

func (u *employeeAccessUsecase) DeleteEmployeeAccess(id int) error {
	return u.employeeAccessRepo.Delete(id)
}

func (u *employeeAccessUsecase) GetEmployeeAccessByOutletID(outletID int) ([]entities.EmployeeAccess, error) {
	return u.employeeAccessRepo.FindByOutletID(outletID)
}

func (u *employeeAccessUsecase) AuthenticateEmployee(request entities.EmployeeLoginRequest) (*entities.EmployeeLoginResponse, error) {
	fmt.Println("::::::", request.Password)
	// Authenticate the employee
	employeeAccess, err := u.employeeAccessRepo.AuthenticateEmployee(request.Username, request.Password)
	if err != nil || employeeAccess == nil {
		return nil, err // Authentication failed
	}
	// Generate JWT token
	tokenString, err := utils.GenerateJWT(employeeAccess.ID, employeeAccess.EmployeeID, employeeAccess.Username, employeeAccess.Role)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":       employeeAccess.ID,
	// 	"username": employeeAccess.Username,
	// 	"role":     employeeAccess.Role,
	// 	"exp":      float64(time.Now().Add(u.tokenExpiry).Unix()),
	// })

	// tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}

	response := &entities.EmployeeLoginResponse{
		AccessToken: tokenString,
		User:        *employeeAccess,
	}

	return response, nil
}
