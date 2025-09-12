package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type serviceCategoryUsecase struct {
	serviceCategoryRepo repositories.ServiceCategoryRepository
}

func NewServiceCategoryUsecase(serviceCategoryRepo repositories.ServiceCategoryRepository) ServiceCategoryUsecase {
	return &serviceCategoryUsecase{
		serviceCategoryRepo: serviceCategoryRepo,
	}
}

func (u *serviceCategoryUsecase) CreateServiceCategory(request entities.CreateServiceCategoryRequest) error {
	category := &entities.ServiceCategory{
		Name:        request.Name,
		Description: request.Description,
	}

	return u.serviceCategoryRepo.Create(category)
}

func (u *serviceCategoryUsecase) GetServiceCategoryByID(id int) (*entities.ServiceCategory, error) {
	return u.serviceCategoryRepo.FindByID(id)
}

func (u *serviceCategoryUsecase) GetAllServiceCategories() ([]entities.ServiceCategory, error) {
	return u.serviceCategoryRepo.FindAll()
}

func (u *serviceCategoryUsecase) GetAllServiceCategoriesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Default ordering
	orderBy := "kategori_id"
	orderDir := "asc"

	// If ordering is specified
	if len(request.Order) > 0 && len(request.Columns) > request.Order[0].Column {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}

	// Get data with pagination
	categories, recordsTotal, recordsFiltered, err := u.serviceCategoryRepo.FindAllWithPagination(
		request.Length,
		request.Start,
		request.Search.Value,
		orderBy,
		orderDir,
	)

	if err != nil {
		return nil, err
	}

	// Create response
	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsFiltered,
		Data:            categories,
	}

	return response, nil
}

func (u *serviceCategoryUsecase) UpdateServiceCategory(id int, request entities.UpdateServiceCategoryRequest) error {
	category, err := u.serviceCategoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return nil // Category not found
	}

	category.Name = request.Name
	category.Description = request.Description

	return u.serviceCategoryRepo.Update(category)
}

func (u *serviceCategoryUsecase) DeleteServiceCategory(id int) error {
	return u.serviceCategoryRepo.Delete(id)
}
