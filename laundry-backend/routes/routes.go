package routes

import (
	"laundry-backend/internal/middleware"
	"laundry-backend/services"
	customerservice "laundry-backend/services/customerService"
	inquiryservice "laundry-backend/services/inquiryService"
	paymentmethodservice "laundry-backend/services/paymentMethodService"
	serviceservice "laundry-backend/services/serviceService"
	transactionservice "laundry-backend/services/transactionService"
	useraccessservice "laundry-backend/services/userAccessService"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func RouteApi(e *echo.Echo, service services.UsecaseService) {
	// Initialize service handlers
	useraccessSvc := useraccessservice.ApiUserAccessService(service)
	serviceSvc := serviceservice.ApiServiceService(service)
	customerSvc := customerservice.ApiCustomerService(service)
	paymentMethodSvc := paymentmethodservice.ApiPaymentMethodService(service)
	inquirySvc := inquiryservice.ApiInquiryService(service)
	transactionSvc := transactionservice.ApiTransactionService(service)
	// Middleware
	// e.Use(echoMiddleware.LoggerWithConfig(
	// 	echoMiddleware.LoggerConfig{
	// 		Format: "\033[32m[${time_rfc3339}] ${remote_ip} ${method} ${uri} ${status} ${latency_human}\nRequest Headers: ${headers}\nRequest Body: ${request}\nResponse: ${response}\033[0m\n",
	// 		Output: nil, // nil means default to os.Stdout
	// 	},
	// ))
	e.Use(echoMiddleware.Recover())

	// CORS
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Auth routes
	// e.POST("/api/v1/login", authSvc.Login)
	e.Use(middleware.NewLoggingMiddleware().LogRequestResponse)
	e.POST("/api/v1/employee/login", useraccessSvc.UserLogin)
	// Administrator routes
	// adminGroup := e.Group("")
	{
		// Brand routes
		// adminGroup.POST("/brands", brandSvc.CreateBrand)
		// adminGroup.GET("/brands/:id", brandSvc.GetBrandByID)
		// adminGroup.GET("/brands", brandSvc.GetAllBrands)
		// adminGroup.PUT("/brands/:id", brandSvc.UpdateBrand)
		// adminGroup.DELETE("/brands/:id", brandSvc.DeleteBrand)

		// Cabang routes
		// adminGroup.POST("/cabangs", cabangSvc.CreateCabang)
		// adminGroup.GET("/cabangs/:id", cabangSvc.GetCabangByID)
		// adminGroup.GET("/cabangs/brand/:brand_id", cabangSvc.GetCabangsByBrandID)
		// adminGroup.GET("/cabangs", cabangSvc.GetAllCabangs)
		// adminGroup.PUT("/cabangs/:id", cabangSvc.UpdateCabang)
		// adminGroup.DELETE("/cabangs/:id", cabangSvc.DeleteCabang)

		// User Access routes
		// adminGroup.POST("/api/v1/user-access", userAccessSvc.CreateUserAccess)
		// adminGroup.GET("/api/v1/user-access/:id", userAccessSvc.GetUserAccessByID)
		// adminGroup.GET("/api/v1/user-access", userAccessSvc.GetAllUserAccessDataTables)
		// adminGroup.PUT("/api/v1/user-access/:id", userAccessSvc.UpdateUserAccess)
		// adminGroup.PUT("/api/v1/user-access/:id/password", userAccessSvc.UpdateUserPassword)
		// adminGroup.DELETE("/api/v1/user-access/:id", userAccessSvc.DeleteUserAccess)
	}

	// Protected routes
	api := e.Group("/api/v1")
	api.Use(echoMiddleware.JWT([]byte("laundry-secret-key")))
	{
		// Outlet routes
		// api.GET("/outlets/:id", outletSvc.GetOutletByID)
		// api.GET("/outlets", outletSvc.GetAllOutlets)

		// Service (Package) routes
		// api.POST("/services", serviceSvc.CreateService)

		// Employee routes
		// api.POST("/pegawai", employeeSvc.CreateEmployee)
		// api.GET("/pegawai/:id", employeeSvc.GetEmployeeByID)
		// api.GET("/pegawai", employeeSvc.GetAllEmployees)
		// api.PUT("/pegawai/:id", employeeSvc.UpdateEmployee)
		// api.DELETE("/pegawai/:id", employeeSvc.DeleteEmployee)

		// Customer routes
		// api.GET("/pelanggan/:id", customerSvc.GetCustomerByID)
		// api.GET("/pelanggan/outlet/:outlet_id", customerSvc.GetCustomersByOutletID)
		// api.PUT("/pelanggan/:id", customerSvc.UpdateCustomer)
		// api.DELETE("/pelanggan/:id", customerSvc.DeleteCustomer)

		// Transaction routes
		api.GET("/transactions/all", transactionSvc.GetAllTransactions)
		// api.GET("/transactions/:id", transactionSvc.GetTransactionByID)
		// api.GET("/transactions/outlet/:outlet_id", transactionSvc.GetTransactionsByOutletID)
		// api.POST("/transactions/detail", transactionSvc.GetTransactionDetails)
		// api.PUT("/transactions/:id/status", transactionSvc.UpdateTransactionStatus)
		// api.PUT("/transactions/:id/payment-status", transactionSvc.UpdatePaymentStatus)
		// api.POST("/transactions/payment-callback", transactionSvc.ProcessPaymentCallback)

		// Additional routes
		api.POST("/services/all", serviceSvc.GetAllServices)
		// api.POST("/service-categories/all", serviceCategorySvc.GetAllServiceCategories)
		api.POST("/pelanggan/all", customerSvc.GetAllCustomers)
		api.POST("/pelanggan/add", customerSvc.CreateCustomer)
		api.POST("/inquiry", inquirySvc.ProcessInquiry)

		// Payment Method routes
		api.POST("/payment-methods/all", paymentMethodSvc.GetAllPaymentMethods)
	}
}
