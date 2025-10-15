package services

import (
	"database/sql"
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
)

type UsecaseService struct {
	RepoDB *sql.DB

	// Repository fields
	RepoUser            userrepo.UserRepo
	RepoBrand           brandrepo.BrandRepo
	RepoCabang          cabangrepo.CabangRepo
	RepoOutlet          outletrepo.OutletRepo
	RepoEmployee        employeerepo.EmployeeRepo
	RepoInquiry         inquiryrepo.InquiryRepo
	RepoCustomer        customerrepo.CustomerRepo
	RepoService         servicerepo.ServiceRepo
	RepoServiceCategory servicecategoryrepo.ServiceCategoryRepo
	RepoUserAccess      useraccessrepo.UserAccessRepo
	RepoTransaction     transactionrepo.TransactionRepo
	RepoPaymentMethod   paymentmethodrepo.PaymentMethodRepo
}

func NewUsecaseService(
	repoDB *sql.DB,
	repoUser userrepo.UserRepo,
	repoBrand brandrepo.BrandRepo,
	repoCabang cabangrepo.CabangRepo,
	repoOutlet outletrepo.OutletRepo,
	repoEmployee employeerepo.EmployeeRepo,
	repoInquiry inquiryrepo.InquiryRepo,
	repoCustomer customerrepo.CustomerRepo,
	repoService servicerepo.ServiceRepo,
	repoServiceCategory servicecategoryrepo.ServiceCategoryRepo,
	repoUserAccess useraccessrepo.UserAccessRepo,
	repoTransaction transactionrepo.TransactionRepo,
	repoPaymentMethod paymentmethodrepo.PaymentMethodRepo,
) UsecaseService {
	return UsecaseService{
		RepoDB:              repoDB,
		RepoUser:            repoUser,
		RepoBrand:           repoBrand,
		RepoCabang:          repoCabang,
		RepoOutlet:          repoOutlet,
		RepoEmployee:        repoEmployee,
		RepoInquiry:         repoInquiry,
		RepoCustomer:        repoCustomer,
		RepoService:         repoService,
		RepoServiceCategory: repoServiceCategory,
		RepoUserAccess:      repoUserAccess,
		RepoTransaction:     repoTransaction,
		RepoPaymentMethod:   repoPaymentMethod,
	}
}
