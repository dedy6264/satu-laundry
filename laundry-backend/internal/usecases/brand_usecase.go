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

func (u *brandUsecase) GetAllBrandsDataTables(request entities.DTRequest[entities.Brand]) (response *entities.DataTablesResponse, err error) {

	// Get data with pagination
	brands, totalCount, err := u.brandRepo.FindAllWithPagination(request)

	if err != nil {
		return response, err
	}

	// Create response
	response = &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalCount,
		Data:            brands,
	}

	return response, nil
}

func (u *brandUsecase) UpdateBrand(id int, request entities.RegisterBrandRequest) error {
	brand, err := u.brandRepo.FindByID(id)
	if err != nil {
		return err
	}

	// if brand == nil {
	// 	return nil // Brand tidak ditemukan
	// }

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
