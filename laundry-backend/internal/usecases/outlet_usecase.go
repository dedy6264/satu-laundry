package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type outletUsecase struct {
	outletRepo repositories.OutletRepository
}

func NewOutletUsecase(outletRepo repositories.OutletRepository) OutletUsecase {
	return &outletUsecase{
		outletRepo: outletRepo,
	}
}

func (u *outletUsecase) CreateOutlet(request entities.RegisterOutletRequest) error {
	outlet := &entities.Outlet{
		CabangID:    request.CabangID,
		Name:        request.Name,
		Address:     request.Address,
		City:        request.City,
		Province:    request.Province,
		PostalCode:  request.PostalCode,
		Phone:       request.Phone,
		Email:       request.Email,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		OpenTime:    request.OpenTime,
		CloseTime:   request.CloseTime,
		PICName:     request.PICName,
		PICEmail:    request.PICEmail,
		PICTelepon:  request.PICTelepon,
	}

	return u.outletRepo.Create(outlet)
}

func (u *outletUsecase) GetOutletByID(id int) (*entities.Outlet, error) {
	return u.outletRepo.FindByID(id)
}

func (u *outletUsecase) GetOutletsByCabangID(cabangID int) ([]entities.Outlet, error) {
	return u.outletRepo.FindByCabangID(cabangID)
}

func (u *outletUsecase) GetAllOutlets() ([]entities.Outlet, error) {
	return u.outletRepo.FindAll()
}

func (u *outletUsecase) GetAllOutletsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Default ordering
	orderBy := "id"
	orderDir := "asc"
	
	// If ordering is specified
	if len(request.Order) > 0 && len(request.Columns) > request.Order[0].Column {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}
	
	// Get data with pagination
	outlets, recordsTotal, recordsFiltered, err := u.outletRepo.FindAllWithPagination(
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
		Data:            outlets,
	}
	
	return response, nil
}

func (u *outletUsecase) UpdateOutlet(id int, request entities.RegisterOutletRequest) error {
	outlet, err := u.outletRepo.FindByID(id)
	if err != nil {
		return err
	}

	if outlet == nil {
		return nil // Outlet tidak ditemukan
	}

	outlet.CabangID = request.CabangID
	outlet.Name = request.Name
	outlet.Address = request.Address
	outlet.City = request.City
	outlet.Province = request.Province
	outlet.PostalCode = request.PostalCode
	outlet.Phone = request.Phone
	outlet.Email = request.Email
	outlet.Latitude = request.Latitude
	outlet.Longitude = request.Longitude
	outlet.OpenTime = request.OpenTime
	outlet.CloseTime = request.CloseTime
	outlet.PICName = request.PICName
	outlet.PICEmail = request.PICEmail
	outlet.PICTelepon = request.PICTelepon

	return u.outletRepo.Update(outlet)
}

func (u *outletUsecase) DeleteOutlet(id int) error {
	return u.outletRepo.Delete(id)
}