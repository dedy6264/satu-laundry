package usecases

import (
	"laundry-backend/internal/entities"
	"laundry-backend/internal/repositories"
)

type employeeUsecase struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{
		employeeRepo: employeeRepo,
	}
}

func (u *employeeUsecase) CreateEmployee(request entities.RegisterEmployeeRequest) error {
	// Hash the password before storing

	employee := &entities.Employee{
		OutletID:  request.OutletID,
		NIK:       request.NIK,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Address:   request.Address,
		BirthDate: request.BirthDate,
		Gender:    request.Gender,
		Position:  request.Position,
		Salary:    request.Salary,
		JoinDate:  request.JoinDate,
		Status:    request.Status,
	}

	return u.employeeRepo.Create(employee)
}

func (u *employeeUsecase) GetEmployeeByID(id int) (*entities.Employee, error) {
	return u.employeeRepo.FindByID(id)
}

func (u *employeeUsecase) GetAllEmployees() ([]entities.Employee, error) {
	return u.employeeRepo.FindAll()
}

func (u *employeeUsecase) GetAllEmployeesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error) {
	// Default ordering
	orderBy := "id_pegawai"
	orderDir := "asc"

	// If ordering is specified
	if len(request.Order) > 0 && len(request.Columns) > request.Order[0].Column {
		orderBy = request.Columns[request.Order[0].Column].Data
		orderDir = request.Order[0].Dir
	}

	// Get data with pagination
	employees, recordsTotal, recordsFiltered, err := u.employeeRepo.FindAllWithPagination(
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
		Data:            employees,
	}

	return response, nil
}

func (u *employeeUsecase) UpdateEmployee(id int, request entities.RegisterEmployeeRequest) error {
	employee, err := u.employeeRepo.FindByID(id)
	if err != nil {
		return err
	}

	if employee == nil {
		return nil // Employee not found
	}

	employee.OutletID = request.OutletID
	employee.NIK = request.NIK
	employee.Name = request.Name
	employee.Email = request.Email
	employee.Phone = request.Phone
	employee.Address = request.Address
	employee.BirthDate = request.BirthDate
	employee.Gender = request.Gender
	employee.Position = request.Position
	employee.Salary = request.Salary
	employee.JoinDate = request.JoinDate
	employee.Status = request.Status

	return u.employeeRepo.Update(employee)
}

func (u *employeeUsecase) DeleteEmployee(id int) error {
	return u.employeeRepo.Delete(id)
}
