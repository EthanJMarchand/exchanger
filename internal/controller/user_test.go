package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethanjmarchand/exchanger/internal/currency"
)

// dont worry about HTTP tests. They mare more challenging.
func TestRender(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	currencyService := &currency.ConverterService{
		APIKey: "FAKEAPIKEY", //This does not work here. Struggling to solve how to test reciver functions where you're passing ENV variables. I passed in the actual APIKey, and the test passed.
	}
	conv := Converter{
		CS: currencyService,
	}
	conv.Render(w, r)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Render() wanted status code 200, got %v", resp.StatusCode)
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Render() : expected the Content-Type to be application/json, but got %v", contentType)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Render() had an error reading the body.")
	}
}
