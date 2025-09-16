package usecases

import (
	"laundry-backend/internal/entities"
)

type AuthUsecase interface {
	Login(request entities.LoginRequest) (*entities.LoginResponse, error)
}

type BrandUsecase interface {
	CreateBrand(request entities.RegisterBrandRequest) error
	GetBrandByID(id int) (*entities.Brand, error)
	GetAllBrands() ([]entities.Brand, error)
	GetAllBrandsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateBrand(id int, request entities.RegisterBrandRequest) error
	DeleteBrand(id int) error
}

type CabangUsecase interface {
	CreateCabang(request entities.RegisterCabangRequest) error
	GetCabangByID(id int) (*entities.Cabang, error)
	GetCabangsByBrandID(brandID int) ([]entities.Cabang, error)
	GetAllCabangs() ([]entities.Cabang, error)
	GetAllCabangsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateCabang(id int, request entities.RegisterCabangRequest) error
	DeleteCabang(id int) error
}

type OutletUsecase interface {
	CreateOutlet(request entities.RegisterOutletRequest) error
	GetOutletByID(id int) (*entities.Outlet, error)
	GetOutletsByCabangID(cabangID int) ([]entities.Outlet, error)
	GetAllOutlets() ([]entities.Outlet, error)
	GetAllOutletsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateOutlet(id int, request entities.RegisterOutletRequest) error
	DeleteOutlet(id int) error
}

type InquiryUsecase interface {
	ProcessInquiry(request entities.InquiryRequest) (*entities.InquiryResponse, error)
}

type EmployeeUsecase interface {
	CreateEmployee(request entities.RegisterEmployeeRequest) error
	GetEmployeeByID(id int) (*entities.Employee, error)
	GetAllEmployees() ([]entities.Employee, error)
	GetAllEmployeesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateEmployee(id int, request entities.RegisterEmployeeRequest) error
	DeleteEmployee(id int) error
}

type CustomerUsecase interface {
	CreateCustomer(request entities.RegisterCustomerRequest) error
	GetCustomerByID(id int) (*entities.Customer, error)
	GetCustomersByOutletID(outletID int) ([]entities.Customer, error)
	GetAllCustomers() ([]entities.Customer, error)
	GetAllCustomersDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateCustomer(id int, request entities.RegisterCustomerRequest) error
	DeleteCustomer(id int) error
}

type ServiceUsecase interface {
	CreateService(request entities.CreateServiceRequest) error
	GetServiceByID(id int) (*entities.Service, error)
	GetAllServices() ([]entities.Service, error)
	GetAllServicesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateService(id int, request entities.UpdateServiceRequest) error
	DeleteService(id int) error
	GetServicesByCategoryID(categoryID int) ([]entities.Service, error)
}

type ServiceCategoryUsecase interface {
	CreateServiceCategory(request entities.CreateServiceCategoryRequest) error
	GetServiceCategoryByID(id int) (*entities.ServiceCategory, error)
	GetAllServiceCategories() ([]entities.ServiceCategory, error)
	GetAllServiceCategoriesDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateServiceCategory(id int, request entities.UpdateServiceCategoryRequest) error
	DeleteServiceCategory(id int) error
}

type EmployeeAccessUsecase interface {
	CreateEmployeeAccess(request entities.CreateEmployeeAccessRequest) error
	GetEmployeeAccessByID(id int) (*entities.EmployeeAccess, error)
	GetAllEmployeeAccess() ([]entities.EmployeeAccess, error)
	GetAllEmployeeAccessDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	UpdateEmployeeAccess(id int, request entities.UpdateEmployeeAccessRequest) error
	UpdateEmployeePassword(id int, request entities.UpdateEmployeePasswordRequest) error
	DeleteEmployeeAccess(id int) error
	GetEmployeeAccessByOutletID(outletID int) ([]entities.EmployeeAccess, error)
	AuthenticateEmployee(request entities.EmployeeLoginRequest) (*entities.EmployeeLoginResponse, error)
}

type TransactionUsecase interface {
	GetAllTransactions() ([]entities.Transaction, error)
	GetAllTransactionsDataTables(request entities.DataTablesRequest) (*entities.DataTablesResponse, error)
	GetTransactionByID(id int) (*entities.Transaction, error)
	GetTransactionsByOutletID(outletID int) ([]entities.Transaction, error)
	GetTransactionDetails(transactionID int) ([]entities.TransactionDetail, error)
	UpdateTransactionStatus(id int, request entities.UpdateTransactionStatusRequest) error
	UpdatePaymentStatus(id int, request entities.UpdatePaymentStatusRequest) error
	ProcessPaymentCallback(request entities.PaymentCallbackRequest) error
}
