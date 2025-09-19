package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"laundry-backend/internal/utils"
	"time"
)

type userAccessUsecase struct {
	userAccessRepo repositories.UserAccessRepository
	cabangRepo     repositories.CabangRepository
	outletRepo     repositories.OutletRepository
	employeeRepo   repositories.EmployeeRepository
	jwtSecret      string
	tokenExpiry    time.Duration
}

func NewUserAccessUsecase(
	userAccessRepo repositories.UserAccessRepository,
	cabangRepo repositories.CabangRepository,
	outletRepo repositories.OutletRepository,
	employeeRepo repositories.EmployeeRepository,
	jwtSecret string,
	tokenExpiry time.Duration,
) UserAccessUsecase {
	return &userAccessUsecase{
		userAccessRepo: userAccessRepo,
		cabangRepo:     cabangRepo,
		outletRepo:     outletRepo,
		employeeRepo:   employeeRepo,
		jwtSecret:      jwtSecret,
		tokenExpiry:    tokenExpiry,
	}
}

func (u *userAccessUsecase) CreateUserAccess(request entities.CreateUserAccessRequest) error {
	access := &entities.UserAccess{
		Username:       request.Username,
		Password:       request.Password,
		Role:           request.Role,
		IsActive:       request.IsActive,
		ReferenceLevel: request.ReferenceLevel,
		ReferenceID:    request.ReferenceID,
	}
	return u.userAccessRepo.Create(access)
}

func (u *userAccessUsecase) GetUserAccessByID(id int) (*entities.UserAccess, error) {
	return u.userAccessRepo.FindByID(id)
}

func (u *userAccessUsecase) GetAllUserAccess() ([]entities.UserAccess, error) {
	return u.userAccessRepo.FindAll()
}

func (u *userAccessUsecase) GetAllUserAccessDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
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

	accesses, totalCount, err := u.userAccessRepo.FindAllWithPagination(limit, offset)
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

func (u *userAccessUsecase) UpdateUserAccess(id int, request entities.UpdateUserAccessRequest) error {
	access, err := u.userAccessRepo.FindByID(id)
	if err != nil {
		return err
	}

	if access == nil {
		return nil // User access not found
	}

	access.Username = request.Username
	access.Role = request.Role
	access.IsActive = request.IsActive
	access.ReferenceLevel = request.ReferenceLevel
	access.ReferenceID = request.ReferenceID

	return u.userAccessRepo.Update(access)
}

func (u *userAccessUsecase) UpdateUserPassword(id int, request entities.UpdateUserPasswordRequest) error {
	// First, get the user access to verify current password
	access, err := u.userAccessRepo.FindByID(id)
	if err != nil {
		return err
	}

	if access == nil {
		return nil // User access not found
	}

	// Authenticate with current password
	authAccess, err := u.userAccessRepo.AuthenticateUser(access.Username, request.CurrentPassword)
	if err != nil || authAccess == nil {
		return err // Current password is incorrect
	}

	// Update with new password
	return u.userAccessRepo.UpdatePassword(id, request.NewPassword)
}

func (u *userAccessUsecase) DeleteUserAccess(id int) error {
	return u.userAccessRepo.Delete(id)
}

func (u *userAccessUsecase) AuthenticateUser(request entities.UserLoginRequest) (response *entities.UserLoginResponse, err error) {
	var (
		id int
	)
	// Authenticate the user
	userAccess, err := u.userAccessRepo.AuthenticateUser(request.Username, request.Password)
	if err != nil || userAccess == nil {
		return nil, err // Authentication failed
	}
	//get data hirarki by reference
	switch userAccess.ReferenceLevel {
	case "cabang":
		cabang, err := u.cabangRepo.FindByID(userAccess.ReferenceID)
		if err != nil {
			return response, err
		}
		id = cabang.ID
	case "outlet":
		outlet, err := u.outletRepo.FindByID(userAccess.ReferenceID)
		if err != nil {
			return response, err
		}
		id = outlet.ID
	default: //karyawan
		employee, err := u.employeeRepo.FindByID(userAccess.ReferenceID)
		if err != nil {
			return response, err
		}
		id = employee.ID
	}
	// Generate JWT token
	tokenString, err := utils.GenerateJWT(userAccess.ID, id, userAccess.Username, userAccess.Role, userAccess.ReferenceLevel)
	if err != nil {
		return nil, err
	}

	response = &entities.UserLoginResponse{
		AccessToken: tokenString,
		User:        *userAccess,
	}

	return response, nil
}
