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

// Config is our config file.
type Config struct {
	CCKey  string
	URL    string
	Server string
}

// loadEnv simple grabs our env variables from our .env file, and returns a config.
func loadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		CCKey:  os.Getenv("CCKEY"),
		URL:    os.Getenv("URL"),
		Server: os.Getenv("SERVER_ADDRESS"),
	}
	if config.CCKey == "" {
		return nil, errors.New("cckey cannot be blank")
	}
	if config.URL == "" {
		return nil, errors.New("url cannot be blank")
	}
	if config.Server == "" {
		return nil, errors.New("server address cannot be blank")
	}
	return config, nil
}

// New service takes a config, and returns is the ConverterService
func NewService(config Config) (currency.ConverterService, error) {
	if config.CCKey == "" {
		return currency.ConverterService{}, errors.New("cckey cannot be empty")
	}
	return currency.ConverterService{
		APIKey: config.CCKey,
		URL:    config.URL,
	}, nil
}

// run is our wrapper to check for errors and to hopefully better test.
func run(config Config) error {
	currencyService, err := NewService(config)
	if err != nil {
		return err
	}
	conv := controller.Converter{
		CS: &currencyService,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Static)
	mux.HandleFunc("GET /exchange/{have}/{want}", conv.Render)
	fmt.Printf("Server starting on port %s..", config.Server)
	err = http.ListenAndServe(config.Server, mux)
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
