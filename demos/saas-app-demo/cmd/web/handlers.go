package main

import (
	"fmt"
	"net/http"
)

// VirtualTerminal displays the virtual terminal page
func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key

	// Render the template
	if err := app.renderTemplate(w, r, "terminal", &templateData{
		StringMap: stringMap,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// (2) Payment Success
func (app *application) PaymentSuccess(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// Read Post data
	cardHolder := r.Form.Get("cardholder-name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment-intent")
	paymentMethod := r.Form.Get("payment-method")
	paymentAmount := r.Form.Get("payment-amount")
	paymentCurrency := r.Form.Get("payment-currency")

	data := make(map[string]interface{})

	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err := app.renderTemplate(w, r, "success", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}

// Displays the page to buy one widget
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "buy-once", nil); err != nil {
		app.errorLog.Println(err)
	}
}
