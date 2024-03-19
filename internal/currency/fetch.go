package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Config is our config file storing out API URL, and our API Key.
type Config struct {
	CCKey string
	URL   string
}

// Type ConverterService has all of the handler functions as reciever's
type ConverterService struct {
	APIKey string
	URL    string
}

// New service takes a config, and returns is the ConverterService
func NewService(config Config) (ConverterService, error) {
	if config.CCKey == "" {
		return ConverterService{}, errors.New("cckey cannot be empty")
	}
	return ConverterService{
		APIKey: config.CCKey,
		URL:    config.URL,
	}, nil
}

type ConvertHistorical struct {
	Query struct {
		Count int `json:"count"`
	} `json:"query"`
	Date    string `json:"date"`
	EndDate string `json:"endDate,omitempty"`
	Results map[string]struct {
		ID  string             `json:"id"`
		To  string             `json:"to"`
		Fr  string             `json:"fr"`
		Val map[string]float32 `json:"val"`
	} `json:"results"`
}

// TODO: Compare needs to be entirely re-written to use the standard library.
func (cs *ConverterService) Compare(have, want string) (*ConvertHistorical, error) {
	var conver = ConvertHistorical{}
	v := url.Values{}
	v.Add("q", have+"_"+want)
	pastdate := time.Now().AddDate(0, 0, -7)
	date := fmt.Sprintf("%d-%d-%d", pastdate.Year(), pastdate.Month(), pastdate.Day())
	endDate := fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	v.Add("date", date)
	v.Add("endDate", endDate)
	v.Add("apiKey", cs.APIKey)
	url := cs.URL + v.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("could not connect to external api")
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bs, &conver)
	if err != nil {
		return nil, err

	}
	return &conver, nil
}
