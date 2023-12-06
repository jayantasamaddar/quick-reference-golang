package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// Create a multiplexer
	mux := chi.NewRouter()

	mux.Get("/virtualterminal", app.VirtualTerminal)
	mux.Post("/payment-success", app.PaymentSuccess)

	mux.Get("/charge-once", app.ChargeOnce)

	// Create FileServer for serving static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
