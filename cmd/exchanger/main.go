package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethanjmarchand/exchanger/internal/controller"
	"github.com/ethanjmarchand/exchanger/internal/currency"
	"github.com/joho/godotenv"
)

// loadEnv simple grabs our env variables from our .env file, and returns a config.
func loadEnv() (*currency.Config, error) {
	config := currency.Config{}
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config.CCKey = os.Getenv("CCKEY")
	config.URL = os.Getenv("APIURL")
	if config.CCKey == "" {
		return nil, errors.New("cckey cannot be an empty string")
	}
	return &config, nil
}

// run is our wrapper to check for errors and to hopefully better test.
func run(config currency.Config) error {
	currencyService, err := currency.NewService(config)
	if err != nil {
		return err
	}
	conv := controller.Converter{
		CS: &currencyService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Static)
	mux.HandleFunc("GET /exchange/{have}/{want}", conv.Render)
	fmt.Println("Server starting on port :3000...")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		return err
	}
	return nil
}

// main is simple the beginning of our application.
func main() {
	config, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = run(*config)
	if err != nil {
		log.Fatal(err)
	}
}
