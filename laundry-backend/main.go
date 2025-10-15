package main

import (
	"context"
	"database/sql"
	"fmt"
	"laundry-backend/apps"
	"laundry-backend/configs"
	"laundry-backend/repositories"
	"laundry-backend/routes"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

func main() {
	env, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	// Initialize Echo instance
	e := echo.New()

	// Initialize database connection
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repositories.NewRepositories(db, ctx)
	// Initialize usecases
	services := apps.SetupApp(db, repo)

	// Routing API
	routes.RouteApi(e, services)

	// Start server
	e.Logger.Fatal(e.Start(":" + env.Server.Address)) // You can use config if needed
}

func initDB() (*sql.DB, error) {
	env, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	// Use configuration values
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Database.Host, env.Database.Port, env.Database.User,
		env.Database.Password, env.Database.Name)

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
