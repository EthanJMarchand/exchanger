package currency

import (
	"errors"
	"fmt"
	"time"

	currconv "github.com/kitloong/go-currency-converter-api/v2"
)

type ConverterService struct {
	APIKey string
}

func NewService(APIKey string) (ConverterService, error) {
	if APIKey == "" {
		return ConverterService{}, errors.New("APIKey cannot be empty")
	}
	return ConverterService{
		APIKey: APIKey,
	}, nil
}

func (cs *ConverterService) Compare(have, want string) (*currconv.ConvertHistorical, error) {
	//  line 25 - 30 would all happen in NewService
	api := currconv.NewAPI(currconv.Config{
		// Change base URL to be an env. variable
		BaseURL: "https://free.currconv.com",
		Version: "v7",
		APIKey:  cs.APIKey,
	})
	conver, err := api.ConvertHistorical(currconv.ConvertHistoricalRequest{
		Q:    []string{have + "_" + want},
		Date: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-7, 0, 0, 0, 0, time.UTC),
		// Create a data and subtract a date from it.
		EndDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}
	return conver, err
}
