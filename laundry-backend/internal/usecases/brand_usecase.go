package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type brandUsecase struct {
	brandRepo repositories.BrandRepository
}

func NewBrandUsecase(brandRepo repositories.BrandRepository) BrandUsecase {
	return &brandUsecase{
		brandRepo: brandRepo,
	}
}

func (u *brandUsecase) CreateBrand(request entities.RegisterBrandRequest) error {
	brand := &entities.Brand{
		Name:        request.Name,
		Description: request.Description,
		PICName:     request.PICName,
		PICEmail:    request.PICEmail,
		PICTelepon:  request.PICTelepon,
		LogoURL:     request.LogoURL,
	}

	return u.brandRepo.Create(brand)
}

func (u *brandUsecase) GetBrandByID(id int) (*entities.Brand, error) {
	return u.brandRepo.FindByID(id)
}

func (u *brandUsecase) GetAllBrands() ([]entities.Brand, error) {
	return u.brandRepo.FindAll()
}

func (u *brandUsecase) GetAllBrandsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Default ordering
	orderBy := "id"
	orderDir := "asc"
	
	// If ordering is specified
	if len(request.Order) > 0 && len(request.Columns) > request.Order[0].Column {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}
	
	// Get data with pagination
	brands, recordsTotal, recordsFiltered, err := u.brandRepo.FindAllWithPagination(
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
		Data:            brands,
	}
	
	return response, nil
}

func (u *brandUsecase) UpdateBrand(id int, request entities.RegisterBrandRequest) error {
	brand, err := u.brandRepo.FindByID(id)
	if err != nil {
		return err
	}

	if brand == nil {
		return nil // Brand tidak ditemukan
	}

	brand.Name = request.Name
	brand.Description = request.Description
	brand.PICName = request.PICName
	brand.PICEmail = request.PICEmail
	brand.PICTelepon = request.PICTelepon
	brand.LogoURL = request.LogoURL

	return u.brandRepo.Update(brand)
}

func (u *brandUsecase) DeleteBrand(id int) error {
	return u.brandRepo.Delete(id)
}