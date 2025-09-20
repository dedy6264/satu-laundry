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
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Brand, int, int, error)
	Update(brand *entities.Brand) error
	Delete(id int) error
}

type CabangRepository interface {
	Create(cabang *entities.Cabang) error
	FindByID(id int) (*entities.Cabang, error)
	FindByBrandID(brandID int) ([]entities.Cabang, error)
	FindAll() ([]entities.Cabang, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Cabang, int, int, error)
	Update(cabang *entities.Cabang) error
	Delete(id int) error
}

type OutletRepository interface {
	Create(outlet *entities.Outlet) error
	FindByID(id int) (*entities.Outlet, error)
	FindByCabangID(cabangID int) ([]entities.Outlet, error)
	FindAll(request entities.Outlet) ([]entities.Outlet, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Outlet, int, int, error)
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
	InsertTransactionWithTx(tx *sql.Tx, transaction *entities.Transaction) (int, error)
	InsertTransactionDetailWithTx(tx *sql.Tx, detail *entities.TransactionDetail) error
	InsertPaymentWithTx(tx *sql.Tx, payment *entities.Payment) error
	InsertHistoryStatusTransactionWithTx(tx *sql.Tx, history *entities.HistoryStatusTransaction) error
}

type EmployeeRepository interface {
	Create(employee *entities.Employee) error
	FindByID(id int) (*entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Employee, int, int, error)
	Update(employee *entities.Employee) error
	Delete(id int) error
}

type CustomerRepository interface {
	Create(customer *entities.Customer) error
	FindByID(id int) (*entities.Customer, error)
	FindByOutletID(outletID int) ([]entities.Customer, error)
	FindAll() ([]entities.Customer, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Customer, int, int, error)
	Update(customer *entities.Customer) error
	Delete(id int) error
}

type ServiceRepository interface {
	Create(service *entities.Service) error
	FindByID(id int) (*entities.Service, error)
	FindAll() ([]entities.Service, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Service, int, error)
	Update(service *entities.Service) error
	Delete(id int) error
	FindByCategoryID(categoryID int) ([]entities.Service, error)
}

type ServiceCategoryRepository interface {
	Create(category *entities.ServiceCategory) error
	FindByID(id int) (*entities.ServiceCategory, error)
	FindAll() ([]entities.ServiceCategory, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.ServiceCategory, int, int, error)
	Update(category *entities.ServiceCategory) error
	Delete(id int) error
}

type UserAccessRepository interface {
	Create(access *entities.UserAccess) error
	FindByID(id int) (*entities.UserAccess, error)
	FindByUsername(username string) (*entities.UserAccess, error)
	FindAll() ([]entities.UserAccess, error)
	FindAllWithPagination(limit, offset int) ([]entities.UserAccess, int, error)
	Update(access *entities.UserAccess) error
	UpdatePassword(id int, password string) error
	UpdateLastLogin(id int) error
	Delete(id int) error
	AuthenticateUser(username, password string) (*entities.UserAccess, error)
}

type TransactionRepository interface {
	FindAll() ([]entities.Transaction, error)
	FindAllWithPagination(limit, offset int, search string, orderBy string, orderDir string) ([]entities.Transaction, int, error)
	FindByID(id int) (*entities.Transaction, error)
	FindByOutletID(outletID int) ([]entities.Transaction, error)
	FindDetailsByTransactionID(transactionID int) ([]entities.TransactionDetail, error)
	UpdateTransactionStatus(id int, status string) error
	UpdatePaymentStatus(id int, status string) error
	UpdatePaymentCallback(transactionID int, request entities.PaymentCallbackRequest) error
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
