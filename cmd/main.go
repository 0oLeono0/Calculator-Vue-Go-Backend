package main

import (
	"calculator/internal/calculationService"
	"calculator/internal/db"
	"calculator/internal/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHendlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHendlers.GetCalculations)
	e.POST("/calculations", calcHendlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHendlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHendlers.DeleteCalculations)

	serverAddr := getEnv("SERVER_ADDR", ":8080")
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
