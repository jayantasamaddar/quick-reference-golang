package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jayantasamaddar/quick-reference-golang/demos/saas-app-demo/internal/cards"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
	}

	// Convert Amount to integer
	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
	}

	// Define Card
	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: payload.Currency,
	}

	okay := true

	paymentIntent, msg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}

	if okay {
		// Charge the credit card
		bytes, err := json.MarshalIndent(paymentIntent, "", "  ")
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	} else {
		j := jsonResponse{
			OK:      true,
			Message: msg,
			Content: "",
		}

		bytes, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLog.Println(err)
		}

		fmt.Println(bytes)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}
}
