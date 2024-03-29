package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethanjmarchand/exchanger/internal/currency"
	"github.com/go-chi/chi"
)

func Static(w http.ResponseWriter, r *http.Request) {
	urlString := `
		<h1>Is today a great time to swap?</h1>
		<p>Visit the endpoint /exchange/{have}/{want} with the 3 letter currency code you are considering in the URL</p>
	`
	fmt.Fprint(w, urlString)
}

type Converter struct {
	CS *currency.ConverterService
}

func (c Converter) Render(w http.ResponseWriter, r *http.Request) {
	have := chi.URLParam(r, "have")
	want := chi.URLParam(r, "want")
	conver, err := c.CS.Compare(have, want)
	if err != nil {
		fmt.Println("render: ", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	bs, err := json.Marshal(conver)
	if err != nil {
		fmt.Println("Render: %w", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
