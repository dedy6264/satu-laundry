package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type cabangUsecase struct {
	cabangRepo repositories.CabangRepository
}

func NewCabangUsecase(cabangRepo repositories.CabangRepository) CabangUsecase {
	return &cabangUsecase{
		cabangRepo: cabangRepo,
	}
}

func (u *cabangUsecase) CreateCabang(request entities.RegisterCabangRequest) error {
	cabang := &entities.Cabang{
		BrandID:    request.BrandID,
		Name:       request.Name,
		Address:    request.Address,
		City:       request.City,
		Province:   request.Province,
		PostalCode: request.PostalCode,
		Phone:      request.Phone,
		Email:      request.Email,
		PICName:    request.PICName,
		PICEmail:   request.PICEmail,
		PICTelepon: request.PICTelepon,
	}

	return u.cabangRepo.Create(cabang)
}

func (u *cabangUsecase) GetCabangByID(id int) (*entities.Cabang, error) {
	return u.cabangRepo.FindByID(id)
}

func (u *cabangUsecase) GetCabangsByBrandID(brandID int) ([]entities.Cabang, error) {
	return u.cabangRepo.FindByBrandID(brandID)
}

func (u *cabangUsecase) GetAllCabangs() ([]entities.Cabang, error) {
	return u.cabangRepo.FindAll()
}

func (u *cabangUsecase) GetAllCabangsDataTables(request entities.DTRequest[entities.Cabang]) (*entities.DataTablesResponse, error) {

	// Get data with pagination
	cabangs, recordsTotal, err := u.cabangRepo.FindAllWithPagination(
		request,
	)

	if err != nil {
		return nil, err
	}

	// Create response
	response := &entities.DataTablesResponse{
		Draw:            request.Draw,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsTotal,
		Data:            cabangs,
	}

	return response, nil
}

func (u *cabangUsecase) UpdateCabang(id int, request entities.RegisterCabangRequest) error {
	cabang, err := u.cabangRepo.FindByID(id)
	if err != nil {
		return err
	}

	if cabang == nil {
		return nil // Cabang tidak ditemukan
	}

	cabang.BrandID = request.BrandID
	cabang.Name = request.Name
	cabang.Address = request.Address
	cabang.City = request.City
	cabang.Province = request.Province
	cabang.PostalCode = request.PostalCode
	cabang.Phone = request.Phone
	cabang.Email = request.Email
	cabang.PICName = request.PICName
	cabang.PICEmail = request.PICEmail
	cabang.PICTelepon = request.PICTelepon

	return u.cabangRepo.Update(cabang)
}

func (u *cabangUsecase) DeleteCabang(id int) error {
	return u.cabangRepo.Delete(id)
}
