package currency

import (
	"fmt"
	"time"

	currconv "github.com/kitloong/go-currency-converter-api/v2"
)

type ConverterService struct {
	APIKey string
}

func (cs *ConverterService) Compare(have, want string) (*currconv.ConvertHistorical, error) {
	api := currconv.NewAPI(currconv.Config{
		BaseURL: "https://free.currconv.com",
		Version: "v7",
		APIKey:  cs.APIKey,
	})
	conver, err := api.ConvertHistorical(currconv.ConvertHistoricalRequest{
		Q:       []string{have + "_" + want},
		Date:    time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-7, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}
	return conver, err
}
