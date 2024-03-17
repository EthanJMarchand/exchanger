package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethanjmarchand/exchanger/internal/controller"
	"github.com/ethanjmarchand/exchanger/internal/currency"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func loadEnv() (currency.ConverterService, error) {
	var cfg currency.ConverterService
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}
	cfg.APIKey = os.Getenv("CCKEY")
	return cfg, nil
}

func main() {
	cfg, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	currencyService := &currency.ConverterService{
		APIKey: cfg.APIKey,
	}
	conv := controller.Converter{
		CS: currencyService,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controller.Static)
	r.Get("/exchange/{have}/{want}", conv.Render)
	fmt.Println("Server starting on port :3000...")
	http.ListenAndServe(":3000", r)
}
