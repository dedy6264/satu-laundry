package main

import (
	"database/sql"
	"fmt"
	"log"

	"laundry-backend/internal/delivery"
	"laundry-backend/internal/middleware"
	"laundry-backend/internal/repositories"
	"laundry-backend/internal/usecases"
	"laundry-backend/internal/utils"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Debug: Print configuration
	fmt.Printf("DB Config: host=%s port=%s user=%s password=%s dbname=%s\n",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.Name)

	// Initialize database connection
	db, err := initDB(config)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	brandRepo := repositories.NewBrandRepository(db)
	cabangRepo := repositories.NewCabangRepository(db)
	outletRepo := repositories.NewOutletRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	inquiryRepo := repositories.NewInquiryRepository(db, employeeRepo)
	customerRepo := repositories.NewCustomerRepository(db)
	serviceRepo := repositories.NewServiceRepository(db)
	serviceCategoryRepo := repositories.NewServiceCategoryRepository(db)
	employeeAccessRepo := repositories.NewEmployeeAccessRepository(db)

	// Initialize usecases
	authUsecase := usecases.NewAuthUsecase(userRepo)
	brandUsecase := usecases.NewBrandUsecase(brandRepo)
	cabangUsecase := usecases.NewCabangUsecase(cabangRepo)
	outletUsecase := usecases.NewOutletUsecase(outletRepo)
	inquiryUsecase := usecases.NewInquiryUsecase(inquiryRepo)
	employeeUsecase := usecases.NewEmployeeUsecase(employeeRepo)
	customerUsecase := usecases.NewCustomerUsecase(customerRepo)
	serviceUsecase := usecases.NewServiceUsecase(serviceRepo)
	serviceCategoryUsecase := usecases.NewServiceCategoryUsecase(serviceCategoryRepo)
	employeeAccessUsecase := usecases.NewEmployeeAccessUsecase(employeeAccessRepo, "laundry-secret-key", 24*60*60) // 24 hours

	// Initialize handlers
	authHandler := delivery.NewAuthHandler(authUsecase)
	brandHandler := delivery.NewBrandHandler(brandUsecase)
	cabangHandler := delivery.NewCabangHandler(cabangUsecase)
	outletHandler := delivery.NewOutletHandler(outletUsecase)
	inquiryHandler := delivery.NewInquiryHandler(inquiryUsecase)
	employeeHandler := delivery.NewEmployeeHandler(employeeUsecase)
	customerHandler := delivery.NewCustomerHandler(customerUsecase)
	serviceHandler := delivery.NewServiceHandler(serviceUsecase)
	serviceCategoryHandler := delivery.NewServiceCategoryHandler(serviceCategoryUsecase)
	employeeAccessHandler := delivery.NewEmployeeAccessHandler(employeeAccessUsecase, "laundry-secret-key")

	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Custom logging middleware
	loggingMiddleware := middleware.NewLoggingMiddleware()
	e.Use(loggingMiddleware.LogRequestResponse)

	// CORS
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	// Auth routes
	e.POST("/api/v1/login", authHandler.Login)
	e.POST("/api/v1/employee/login", employeeAccessHandler.EmployeeLogin)

	// Protected routes
	api := e.Group("/api/v1")
	{
		api.Use(echoMiddleware.JWT([]byte("laundry-secret-key")))

		// Brand routes
		api.POST("/brands", brandHandler.CreateBrand)
		api.GET("/brands/:id", brandHandler.GetBrandByID)
		api.GET("/brands", brandHandler.GetAllBrands)
		api.PUT("/brands/:id", brandHandler.UpdateBrand)
		api.DELETE("/brands/:id", brandHandler.DeleteBrand)

		// Cabang routes
		api.POST("/cabangs", cabangHandler.CreateCabang)
		api.GET("/cabangs/:id", cabangHandler.GetCabangByID)
		api.GET("/cabangs/brand/:brand_id", cabangHandler.GetCabangsByBrandID)
		api.GET("/cabangs", cabangHandler.GetAllCabangs)
		api.PUT("/cabangs/:id", cabangHandler.UpdateCabang)
		api.DELETE("/cabangs/:id", cabangHandler.DeleteCabang)

		// Outlet routes
		api.POST("/outlets", outletHandler.CreateOutlet)
		api.GET("/outlets/:id", outletHandler.GetOutletByID)
		api.GET("/outlets/cabang/:cabang_id", outletHandler.GetOutletsByCabangID)
		api.GET("/outlets", outletHandler.GetAllOutlets)
		api.PUT("/outlets/:id", outletHandler.UpdateOutlet)
		api.DELETE("/outlets/:id", outletHandler.DeleteOutlet)

		// Inquiry routes
		api.POST("/inquiry", inquiryHandler.ProcessInquiry)

		// Employee routes
		api.POST("/pegawai", employeeHandler.CreateEmployee)
		api.GET("/pegawai/:id", employeeHandler.GetEmployeeByID)
		api.GET("/pegawai", employeeHandler.GetAllEmployees)
		api.PUT("/pegawai/:id", employeeHandler.UpdateEmployee)
		api.DELETE("/pegawai/:id", employeeHandler.DeleteEmployee)

		// Customer routes
		api.POST("/pelanggan", customerHandler.CreateCustomer)
		api.GET("/pelanggan/:id", customerHandler.GetCustomerByID)
		api.GET("/pelanggan/outlet/:outlet_id", customerHandler.GetCustomersByOutletID)
		api.GET("/pelanggan", customerHandler.GetAllCustomers)
		api.PUT("/pelanggan/:id", customerHandler.UpdateCustomer)
		api.DELETE("/pelanggan/:id", customerHandler.DeleteCustomer)

		// Service routes
		api.POST("/services", serviceHandler.CreateService)
		api.GET("/services/:id", serviceHandler.GetServiceByID)
		api.GET("/services", serviceHandler.GetAllServices)
		api.PUT("/services/:id", serviceHandler.UpdateService)
		api.DELETE("/services/:id", serviceHandler.DeleteService)
		api.GET("/services/category/:category_id", serviceHandler.GetServicesByCategoryID)

		// Service Category routes
		api.POST("/service-categories", serviceCategoryHandler.CreateServiceCategory)
		api.GET("/service-categories/:id", serviceCategoryHandler.GetServiceCategoryByID)
		api.GET("/service-categories", serviceCategoryHandler.GetAllServiceCategories)
		api.PUT("/service-categories/:id", serviceCategoryHandler.UpdateServiceCategory)
		api.DELETE("/service-categories/:id", serviceCategoryHandler.DeleteServiceCategory)

		// Employee Access routes
		api.POST("/employee-access", employeeAccessHandler.CreateEmployeeAccess)
		api.GET("/employee-access/:id", employeeAccessHandler.GetEmployeeAccessByID)
		api.GET("/employee-access", employeeAccessHandler.GetAllEmployeeAccessDataTables)
		api.PUT("/employee-access/:id", employeeAccessHandler.UpdateEmployeeAccess)
		api.PUT("/employee-access/:id/password", employeeAccessHandler.UpdateEmployeePassword)
		api.DELETE("/employee-access/:id", employeeAccessHandler.DeleteEmployeeAccess)
		api.GET("/employee-access/outlet/:outlet_id", employeeAccessHandler.GetEmployeeAccessByOutletID)
	}
	// Start server
	e.Logger.Fatal(e.Start(config.Server.Address))
}

func initDB(config *utils.Config) (*sql.DB, error) {
	// Format PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60) // 5 minutes

	fmt.Println("Successfully connected to PostgreSQL database")
	return db, nil
}
