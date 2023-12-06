# Payment and Subscriptions App

# Approach

1. We will be building two binaries from the same codebase.

- `cmd/web`: for the frontend
- `cmd/api`: for the backend api server

2. Add the CORS Middleware. We need to enable CORS to allow only specific IP Addresses or Domains to access. Use [go-chi/cors](https://github.com/go-chi/cors).
3. Set up route handlers.
4. Connect Frontend application to Backend
   - Remove hardcoded keys and api endpoints and replace with something we can get using code

---

## Project Structure

```bash
├── README.md
├── cmd
│   ├── api
│   └── web
│       └── main.go
├── go.mod
└── internal
```

- cmd
  - api
  - web: Frontend of the application

---

# Stripe

## PaymentIntent

A PaymentIntent transitions through multiple statuses throughout its lifetime as it interfaces with Stripe.js to perform authentication flows and ultimately creates at most one successful transaction.

---

## Response

```json
{
  "paymentIntent": {
    "id": "pi_1Nz6i9F3UonGYRD66SiTWvVQ",
    "object": "payment_intent",
    "amount": 10000,
    "amount_details": {
      "tip": {}
    },
    "automatic_payment_methods": {
      "allow_redirects": "always",
      "enabled": true
    },
    "canceled_at": null,
    "cancellation_reason": null,
    "capture_method": "automatic",
    "client_secret": "pi_1Nz6i9F3UonGYRD66SiTWvVQ_secret_vhlR9pdvkUuWA8l1Qe0biZTYB",
    "confirmation_method": "automatic",
    "created": 1696808917,
    "currency": "inr",
    "description": null,
    "last_payment_error": null,
    "livemode": false,
    "next_action": null,
    "payment_method": "pm_1Nz6iAF3UonGYRD62Zsws6r3",
    "payment_method_configuration_details": null,
    "payment_method_types": ["card"],
    "processing": null,
    "receipt_email": null,
    "setup_future_usage": null,
    "shipping": null,
    "source": null,
    "status": "succeeded"
  }
}
```

---

## Usage

- Define a `.env` file with environment variables.
  - `STRIPE_KEY`
  - `STRIPE_SECRET`
- Launch the frontend application by running `go run ./cmd/web -p=[PORT] -e=["development"|"production"] -api="http://localhost:4001"`

---

# Selling a Product Online

## Database Setup

1. Create Database

```sql
CREATE DATABASE widgets
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = 'libc'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
```

---

## Create a Product Page

---

# Referenes

- [Stripe Testing Reference](https://stripe.com/docs/testing)
- [Stripe Go](https://github.com/stripe/stripe-go)
