package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type serviceUsecase struct {
	serviceRepo repositories.ServiceRepository
}

func NewServiceUsecase(serviceRepo repositories.ServiceRepository) ServiceUsecase {
	return &serviceUsecase{
		serviceRepo: serviceRepo,
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
	return u.serviceRepo.FindAll()
}

func (u *serviceUsecase) GetAllServicesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Get order column
	var orderBy string
	var orderDir string
	if len(request.Order) > 0 && request.Order[0].Column < len(request.Columns) {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}

	// Map column names to database column names
	columnMap := map[string]string{
		"id":             "id_layanan",
		"kategori_id":    "kategori_id",
		"nama_layanan":   "nama_layanan",
		"harga":          "harga",
		"satuan":         "satuan",
		"estimasi_waktu": "estimasi_waktu",
		"created_at":     "created_at",
	}

	if dbColumn, exists := columnMap[orderBy]; exists {
		orderBy = dbColumn
	} else {
		orderBy = "id_layanan"
	}

	services, totalCount, err := u.serviceRepo.FindAllWithPagination(
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
		Data:            services,
	}

	return response, nil
}

func (u *serviceUsecase) UpdateService(id int, request entities.UpdateServiceRequest) error {
	// First get the existing service to preserve the BrandID
	existingService, err := u.serviceRepo.FindByID(id)
	if err != nil {
		return err
	}

	service := &entities.Service{
		ID:          id,
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
