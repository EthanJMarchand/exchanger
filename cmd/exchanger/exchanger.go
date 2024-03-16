package main

import (
	"fmt"
	"net/http"

	"github.com/ethanjmarchand/exchanger/internal/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controller.Static)
	r.Get("/exchange/{have}/{want}", controller.Render)
	fmt.Println("Server starting on port :3000...")
	http.ListenAndServe(":3000", r)
}
