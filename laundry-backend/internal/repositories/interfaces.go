package repositories

import (
	"database/sql"
	"laundry-backend/internal/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	Create(user *entities.User) error
}

type BrandRepository interface {
	Create(brand *entities.Brand) error
	FindByID(id int) (*entities.Brand, error)
	FindAll() ([]entities.Brand, error)
	FindAllWithPagination(request entities.DTRequest[entities.Brand]) ([]entities.Brand, int, error)
	Update(brand *entities.Brand) error
	Delete(id int) error
}

type CabangRepository interface {
	Create(cabang *entities.Cabang) error
	FindByID(id int) (*entities.Cabang, error)
	FindByBrandID(brandID int) ([]entities.Cabang, error)
	FindAll() ([]entities.Cabang, error)
	FindAllWithPagination(request entities.DTRequest[entities.Cabang]) ([]entities.Cabang, int, error)
	Update(cabang *entities.Cabang) error
	Delete(id int) error
}

type OutletRepository interface {
	Create(outlet *entities.Outlet) error
	FindByID(id int) (*entities.Outlet, error)
	FindByCabangID(cabangID int) ([]entities.Outlet, error)
	FindAll(request entities.Outlet) ([]entities.Outlet, error)
	FindAllWithPagination(request entities.DTRequest[entities.Outlet]) ([]entities.Outlet, int, error)
	Update(outlet *entities.Outlet) error
	Delete(id int) error
}

type InquiryRepository interface {
	// ValidateServicePackage(id int) (bool, error)
	ValidateEmployee(id int) (*entities.Employee, error)
	ValidateCustomer(id int) (bool, error)
	// GetServicePackagePrice(id int) (float64, error)
	// Transaction methods
	BeginTransaction() (*sql.Tx, error)
	InsertTransactionWithTx(tx *sql.Tx, transaction entities.Transaction) (int, error)
	InsertTransactionDetailWithTx(tx *sql.Tx, detail []entities.TransactionDetail) (err error)
	InsertPaymentWithTx(tx *sql.Tx, payment entities.Payment) error
	InsertHistoryStatusTransactionWithTx(tx *sql.Tx, history entities.HistoryStatusTransaction) error
}

type EmployeeRepository interface {
	Create(employee *entities.Employee) error
	FindByID(id int) (*entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	FindAllWithPagination(request entities.DTRequest[entities.Employee]) ([]entities.Employee, int, error)
	Update(employee *entities.Employee) error
	Delete(id int) error
}

type CustomerRepository interface {
	Create(customer *entities.Customer) error
	FindByID(id int) (*entities.Customer, error)
	FindByOutletID(outletID int) ([]entities.Customer, error)
	FindAll() ([]entities.Customer, error)
	FindAllWithPagination(request entities.DTRequest[entities.Customer]) ([]entities.Customer, int, error)
	Update(customer *entities.Customer) error
	Delete(id int) error
}

type ServiceRepository interface {
	Create(service *entities.Service) error
	FindByID(id int) (*entities.Service, error)
	FindAll(request entities.Service) ([]entities.Service, error)
	FindAllWithPagination(request entities.DTRequest[entities.Service]) ([]entities.Service, int, error)
	Update(service *entities.Service) error
	Delete(id int) error
	FindByCategoryID(categoryID int) ([]entities.Service, error)
}

type ServiceCategoryRepository interface {
	Create(category *entities.ServiceCategory) error
	FindByID(id int) (*entities.ServiceCategory, error)
	FindAll() ([]entities.ServiceCategory, error)
	FindAllWithPagination(request entities.DTRequest[entities.ServiceCategory]) ([]entities.ServiceCategory, int, error)
	Update(category *entities.ServiceCategory) error
	Delete(id int) error
}

type UserAccessRepository interface {
	Create(access *entities.UserAccess) error
	FindByID(id int) (*entities.UserAccess, error)
	FindByUsername(username string) (*entities.UserAccess, error)
	FindAll() ([]entities.UserAccess, error)
	FindAllWithPagination(request entities.DTRequest[entities.UserAccess]) ([]entities.UserAccess, int, error)
	Update(access *entities.UserAccess) error
	UpdatePassword(id int, password string) error
	UpdateLastLogin(id int) error
	Delete(id int) error
	AuthenticateUser(username, password string) (*entities.UserAccess, error)
}

type TransactionRepository interface {
	FindAll() ([]entities.Transaction, error)
	FindAllWithPagination(request entities.DTRequest[entities.Transaction]) ([]entities.Transaction, int, error)
	FindByID(id int) (entities.Transaction, error)
	FindByOutletID(outletID int) ([]entities.Transaction, error)
	FindDetailsByTransactionID(transactionID int) ([]entities.TransactionDetail, error)
	UpdateTransactionStatus(id int, status string) error
	UpdatePaymentStatus(id int, status string) error
	UpdatePaymentCallback(transactionID int, request entities.PaymentCallbackRequest) error
}

type TransactionUsecase interface {
	GetAllTransactions() ([]entities.Transaction, error)
	GetAllTransactionsDataTables(request entities.DTRequest[entities.Transaction]) (*entities.DataTablesResponse, error)
	GetTransactionByID(id int) (*entities.Transaction, error)
	GetTransactionsByOutletID(outletID int) ([]entities.Transaction, error)
	GetTransactionDetails(transactionID int) ([]entities.TransactionDetail, error)
	UpdateTransactionStatus(id int, request entities.UpdateTransactionStatusRequest) error
	UpdatePaymentStatus(id int, request entities.UpdatePaymentStatusRequest) error
	ProcessPaymentCallback(request entities.PaymentCallbackRequest) error
}
type PaymentMethodRepository interface {
	Create(paymentMethod *entities.PaymentMethod) error
	FindByID(id int) (*entities.PaymentMethod, error)
	FindAll() ([]entities.PaymentMethod, error)
	FindAllWithPagination(request entities.DTRequest[entities.PaymentMethod]) ([]entities.PaymentMethod, int, error)
	Update(paymentMethod *entities.PaymentMethod) error
	Delete(id int) error
}
