package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ethanjmarchand/exchanger/internal/currency"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func Static(w http.ResponseWriter, r *http.Request) {
	urlString := `
		<h1>Is today a great time to swap?</h1>
		<p>Visit the endpoint /exchange/{have}/{want} with the 3 letter currency code you are considering in the URL</p>
	`
	fmt.Fprint(w, urlString)
}

func Render(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}
	have := chi.URLParam(r, "have")
	want := chi.URLParam(r, "want")
	conver, err := currency.Compare(os.Getenv("CCKEY"), have, want)
	if err != nil {
		fmt.Println("render: %w", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, conver)
}
