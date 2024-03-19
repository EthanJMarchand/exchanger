package main

import (
	"errors"
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

// TODO: Test this funtion.
// Mixing concerns. This should only load an Env variable. Should be returning a string, or config struct.
func loadEnvKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	CCKey := os.Getenv("CCKEY")
	// Check to see if CCKEY is empty string.
	if CCKey == "" {
		return "", errors.New("CCKey cannot be an empty string.")
	}
	return CCKey, nil
}

func run(CCKey string) error {
	currencyService, err := currency.NewService(CCKey)
	if err != nil {
		return err
	}
	conv := controller.Converter{
		CS: &currencyService,
	}
	// TODO: Get rid of chi.
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controller.Static)
	r.Get("/exchange/{have}/{want}", conv.Render)
	fmt.Println("Server starting on port :3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	CCKey, err := loadEnvKey()
	if err != nil {
		log.Fatal(err)
	}
	err = run(CCKey)
	if err != nil {
		log.Fatal(err)
	}
}
