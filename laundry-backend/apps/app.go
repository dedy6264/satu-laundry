package apps

import (
	"database/sql"
	"laundry-backend/repositories"
	"laundry-backend/repositories/brandrepo"
	"laundry-backend/repositories/cabangrepo"
	"laundry-backend/repositories/customerrepo"
	"laundry-backend/repositories/employeerepo"
	"laundry-backend/repositories/inquiryrepo"
	"laundry-backend/repositories/outletrepo"
	"laundry-backend/repositories/paymentmethodrepo"
	"laundry-backend/repositories/servicecategoryrepo"
	"laundry-backend/repositories/servicerepo"
	"laundry-backend/repositories/transactionrepo"
	"laundry-backend/repositories/useraccessrepo"
	"laundry-backend/repositories/userrepo"
	"laundry-backend/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repositories) services.UsecaseService {
	// Initialize repositories
	userRepo := userrepo.NewUserRepo(repo)
	brandRepo := brandrepo.NewBrandRepo(repo)
	cabangRepo := cabangrepo.NewCabangRepo(repo)
	outletRepo := outletrepo.NewOutletRepo(repo)
	employeeRepo := employeerepo.NewEmployeeRepo(repo)
	inquiryRepo := inquiryrepo.NewInquiryRepo(repo)
	customerRepo := customerrepo.NewCustomerRepo(repo)
	serviceRepo := servicerepo.NewServiceRepo(repo)
	serviceCategoryRepo := servicecategoryrepo.NewServiceCategoryRepo(repo)
	userAccessRepo := useraccessrepo.NewUserAccessRepo(repo)
	transactionRepo := transactionrepo.NewTransactionRepo(repo)
	paymentMethodRepo := paymentmethodrepo.NewPaymentMethodRepo(repo)

	usecaseSvc := services.NewUsecaseService(
		DB,
		userRepo,
		brandRepo,
		cabangRepo,
		outletRepo,
		employeeRepo,
		inquiryRepo,
		customerRepo,
		serviceRepo,
		serviceCategoryRepo,
		userAccessRepo,
		transactionRepo,
		paymentMethodRepo,
	)

	return usecaseSvc
}
