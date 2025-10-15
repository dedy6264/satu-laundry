package usecases

import (
	"laundry-backend/internal/entities"

	"github.com/golang-jwt/jwt"
)

type AuthUsecase interface {
	Login(request entities.LoginRequest) (*entities.LoginResponse, error)
}

type BrandUsecase interface {
	CreateBrand(request entities.RegisterBrandRequest) error
	GetBrandByID(id int) (*entities.Brand, error)
	GetAllBrands() ([]entities.Brand, error)
	GetAllBrandsDataTables(request entities.DTRequest[entities.Brand]) (*entities.DataTablesResponse, error)
	UpdateBrand(id int, request entities.RegisterBrandRequest) error
	DeleteBrand(id int) error
}

type CabangUsecase interface {
	CreateCabang(request entities.RegisterCabangRequest) error
	GetCabangByID(id int) (*entities.Cabang, error)
	GetCabangsByBrandID(brandID int) ([]entities.Cabang, error)
	GetAllCabangs() ([]entities.Cabang, error)
	GetAllCabangsDataTables(request entities.DTRequest[entities.Cabang]) (*entities.DataTablesResponse, error)
	UpdateCabang(id int, request entities.RegisterCabangRequest) error
	DeleteCabang(id int) error
}

type OutletUsecase interface {
	CreateOutlet(request entities.RegisterOutletRequest) error
	GetOutletByID(id int) (*entities.Outlet, error)
	GetOutletsByCabangID(cabangID int) ([]entities.Outlet, error)
	GetAllOutlets() ([]entities.Outlet, error)
	GetAllOutletsDataTables(request entities.DTRequest[entities.Outlet]) (*entities.DataTablesResponse, error)
	UpdateOutlet(id int, request entities.RegisterOutletRequest) error
	DeleteOutlet(id int) error
}

type InquiryUsecase interface {
	ProcessInquiry(request entities.InquiryRequest, claims jwt.MapClaims) (entities.InquiryResponse, error)
}

type EmployeeUsecase interface {
	CreateEmployee(request entities.RegisterEmployeeRequest) error
	GetEmployeeByID(id int) (*entities.Employee, error)
	GetAllEmployees() ([]entities.Employee, error)
	GetAllEmployeesDataTables(request entities.DTRequest[entities.Employee]) (*entities.DataTablesResponse, error)
	UpdateEmployee(id int, request entities.RegisterEmployeeRequest) error
	DeleteEmployee(id int) error
}

type CustomerUsecase interface {
	CreateCustomer(request entities.RegisterCustomerRequest) error
	GetCustomerByID(id int) (*entities.Customer, error)
	GetCustomersByOutletID(outletID int) ([]entities.Customer, error)
	GetAllCustomers() ([]entities.Customer, error)
	GetAllCustomersDataTables(request entities.DTRequest[entities.Customer]) (*entities.DataTablesResponse, error)
	UpdateCustomer(id int, request entities.RegisterCustomerRequest) error
	DeleteCustomer(id int) error
}

type ServiceUsecase interface {
	CreateService(request entities.CreateServiceRequest) error
	GetServiceByID(id int) (*entities.Service, error)
	GetAllServices() ([]entities.Service, error)
	GetAllServicesDataTables(request entities.DTRequest[entities.Service]) (*entities.DataTablesResponse, error)
	UpdateService(request entities.Service) error
	DeleteService(id int) error
	GetServicesByCategoryID(categoryID int) ([]entities.Service, error)
}

type ServiceCategoryUsecase interface {
	CreateServiceCategory(request entities.ServiceCategory) error
	GetServiceCategoryByID(id int) (*entities.ServiceCategory, error)
	GetAllServiceCategories() ([]entities.ServiceCategory, error)
	GetAllServiceCategoriesDataTables(request entities.DTRequest[entities.ServiceCategory]) (*entities.DataTablesResponse, error)
	UpdateServiceCategory(request entities.ServiceCategory) error
	DeleteServiceCategory(id int) error
}

type UserAccessUsecase interface {
	CreateUserAccess(request entities.CreateUserAccessRequest) error
	GetUserAccessByID(id int) (*entities.UserAccess, error)
	GetAllUserAccess() ([]entities.UserAccess, error)
	GetAllUserAccessDataTables(request entities.DTRequest[entities.UserAccess]) (*entities.DataTablesResponse, error)
	UpdateUserAccess(id int, request entities.UpdateUserAccessRequest) error
	UpdateUserPassword(id int, request entities.UpdateUserPasswordRequest) error
	DeleteUserAccess(id int) error
	AuthenticateUser(request entities.UserLoginRequest) (*entities.UserLoginResponse, error)
}

type TransactionUsecase interface {
	GetAllTransactions() ([]entities.Transaction, error)
	GetAllTransactionsDataTables(request entities.DTRequest[entities.Transaction]) (entities.DataTablesResponse, error)
	GetTransactionByID(id int) (entities.Transaction, error)
	GetTransactionsByOutletID(outletID int) ([]entities.Transaction, error)
	GetTransactionDetails(request entities.Transaction) (entities.TransactionComplete, error)
	UpdateTransactionStatus(id int, request entities.UpdateTransactionStatusRequest) error
	UpdatePaymentStatus(id int, request entities.UpdatePaymentStatusRequest) error
	ProcessPaymentCallback(request entities.PaymentCallbackRequest) error
}
type PaymentMethodUsecase interface {
	CreatePaymentMethod(request entities.CreatePaymentMethodRequest) error
	GetPaymentMethodByID(id int) (*entities.PaymentMethod, error)
	GetAllPaymentMethods() ([]entities.PaymentMethod, error)
	GetAllPaymentMethodsDataTables(request entities.DTRequest[entities.PaymentMethod]) (*entities.DataTablesResponse, error)
	UpdatePaymentMethod(id int, request entities.UpdatePaymentMethodRequest) error
	DeletePaymentMethod(id int) error
}
