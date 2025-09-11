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

	// Initialize usecases
	authUsecase := usecases.NewAuthUsecase(userRepo)
	brandUsecase := usecases.NewBrandUsecase(brandRepo)
	cabangUsecase := usecases.NewCabangUsecase(cabangRepo)
	outletUsecase := usecases.NewOutletUsecase(outletRepo)
	inquiryUsecase := usecases.NewInquiryUsecase(inquiryRepo)
	employeeUsecase := usecases.NewEmployeeUsecase(employeeRepo)
	customerUsecase := usecases.NewCustomerUsecase(customerRepo)

	// Initialize handlers
	authHandler := delivery.NewAuthHandler(authUsecase)
	brandHandler := delivery.NewBrandHandler(brandUsecase)
	cabangHandler := delivery.NewCabangHandler(cabangUsecase)
	outletHandler := delivery.NewOutletHandler(outletUsecase)
	inquiryHandler := delivery.NewInquiryHandler(inquiryUsecase)
	employeeHandler := delivery.NewEmployeeHandler(employeeUsecase)
	customerHandler := delivery.NewCustomerHandler(customerUsecase)

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
	e.POST("/api/v1/pegawai/login", employeeHandler.Login)

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
