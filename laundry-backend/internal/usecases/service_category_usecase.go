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

func (u *serviceCategoryUsecase) CreateServiceCategory(request entities.ServiceCategory) error {
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

func (u *serviceCategoryUsecase) GetAllServiceCategoriesDataTables(request entities.DTRequest[entities.ServiceCategory]) (*entities.DataTablesResponse, error) {

	categories, totalCount, err := u.serviceCategoryRepo.FindAllWithPagination(
		request,
	)
	if err != nil {
		return nil, err
	}

	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            categories,
	}

	return response, nil
}

func (u *serviceCategoryUsecase) UpdateServiceCategory(request entities.ServiceCategory) error {
	// First get the existing category
	existingCategory, err := u.serviceCategoryRepo.FindByID(request.ID)
	if err != nil {
		return err
	}

	category := &entities.ServiceCategory{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   existingCategory.CreatedAt,
	}

	return u.serviceCategoryRepo.Update(category)
}

func (u *serviceCategoryUsecase) DeleteServiceCategory(id int) error {
	return u.serviceCategoryRepo.Delete(id)
}
