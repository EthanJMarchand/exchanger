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

// Type ConverterService has all of the handler functions as reciever's
type ConverterService struct {
	APIKey string
	URL    string
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

// qstring is my small helper function that just formats our string for us.
func qstring(have, want string) string {
	return have + "_" + want
}

func (cs *ConverterService) Compare(have, want string) (*ConvertHistorical, error) {
	var conver = ConvertHistorical{}
	v := url.Values{}
	v.Add("q", qstring(have, want))
	pastdate := time.Now().AddDate(0, 0, -7)
	date := fmt.Sprintf("%d-%d-%d", pastdate.Year(), pastdate.Month(), pastdate.Day())
	endDate := fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	v.Add("date", date)
	v.Add("endDate", endDate)
	v.Add("apiKey", cs.APIKey)
	url := cs.URL + v.Encode()
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Compare(): %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("could not connect to external api")
	}
	bs, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Compare(): %w", err)
	}
	err = json.Unmarshal(bs, &conver)
	if err != nil {
		return nil, fmt.Errorf("Compare(): %w", err)

	}
	return &conver, nil
}
