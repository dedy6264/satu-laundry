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
	// Get order column
	var orderBy string
	var orderDir string
	if len(request.Order) > 0 && request.Order[0].Column < len(request.Columns) {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}

	// Map column names to database column names
	columnMap := map[string]string{
		"id":          "id_kategori",
		"nama_kategori": "nama_kategori",
		"deskripsi":   "deskripsi",
		"created_at":  "created_at",
	}

	if dbColumn, exists := columnMap[orderBy]; exists {
		orderBy = dbColumn
	} else {
		orderBy = "id_kategori"
	}

	categories, totalCount, _, err := u.serviceCategoryRepo.FindAllWithPagination(
		request.Length,
		request.Start,
		request.Search.Value,
		orderBy,
		orderDir,
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

func (u *serviceCategoryUsecase) UpdateServiceCategory(id int, request entities.UpdateServiceCategoryRequest) error {
	// First get the existing category
	existingCategory, err := u.serviceCategoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	category := &entities.ServiceCategory{
		ID:          id,
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   existingCategory.CreatedAt,
	}

	return u.serviceCategoryRepo.Update(category)
}

func (u *serviceCategoryUsecase) DeleteServiceCategory(id int) error {
	return u.serviceCategoryRepo.Delete(id)
}