package usecases

import (
	"database/sql"
	"errors"
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
	"log"
)

type serviceUsecase struct {
	serviceRepo    repositories.ServiceRepository
	userAccessRepo repositories.UserAccessRepository
	outletRepo     repositories.OutletRepository
	employeeRepo   repositories.EmployeeRepository
}

func NewServiceUsecase(serviceRepo repositories.ServiceRepository,
	userAccessRepo repositories.UserAccessRepository,
	outletRepo repositories.OutletRepository,
	employeeRepo repositories.EmployeeRepository,
) ServiceUsecase {
	return &serviceUsecase{
		serviceRepo:    serviceRepo,
		userAccessRepo: userAccessRepo,
		outletRepo:     outletRepo,
		employeeRepo:   employeeRepo,
	}
}

func (u *serviceUsecase) CreateService(request entities.CreateServiceRequest) error {
	service := &entities.Service{
		BrandID:     request.BrandID,
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Unit:        request.Unit,
		Estimation:  request.Estimation,
	}

	return u.serviceRepo.Create(service)
}

func (u *serviceUsecase) GetServiceByID(id int) (*entities.Service, error) {
	return u.serviceRepo.FindByID(id)
}

func (u *serviceUsecase) GetAllServices() ([]entities.Service, error) {
	return u.serviceRepo.FindAll(entities.Service{})
}

func (u *serviceUsecase) GetAllServicesDataTables(request entities.DTRequest[entities.Service]) (*entities.DataTablesResponse, error) {
	var brandID int
	// validasi user
	userAccess, err := u.userAccessRepo.FindByID(request.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid userAccess")
		}
		return nil, err
	}

	// // 2. get outlet
	if userAccess.ReferenceLevel != "cabang" {
		switch userAccess.ReferenceLevel {
		case "pegawai":

			employee, err := u.employeeRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				log.Println(":: employeeRepo.FindByID")
				return nil, err
			}
			brandID = employee.BrandID
		case "outlet":
			outlet, err := u.outletRepo.FindByID(userAccess.ReferenceID)
			if err != nil {
				log.Println(":: outletRepo.FindByID")
				return nil, err
			}
			brandID = outlet.BrandID
		default:
			return nil, errors.New("invalid reference level")
		}
	}

	request.Data.BrandID = brandID
	services, totalCount, err := u.serviceRepo.FindAllWithPagination(request)
	if err != nil {
		return nil, err
	}

	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            services,
	}

	return response, nil
}

func (u *serviceUsecase) UpdateService(request entities.Service) error {
	// First get the existing service to preserve the BrandID
	existingService, err := u.serviceRepo.FindByID(request.ID)
	if err != nil {
		return err
	}

	service := &entities.Service{
		ID:          request.ID,
		BrandID:     existingService.BrandID, // Preserve the existing BrandID
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Unit:        request.Unit,
		Estimation:  request.Estimation,
	}

	return u.serviceRepo.Update(service)
}

func (u *serviceUsecase) DeleteService(id int) error {
	return u.serviceRepo.Delete(id)
}

func (u *serviceUsecase) GetServicesByCategoryID(categoryID int) ([]entities.Service, error) {
	return u.serviceRepo.FindByCategoryID(categoryID)
}
