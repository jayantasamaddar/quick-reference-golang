// Main entrypoint for the api package
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare application constants
const (
	VERSION string = "1.0.0"
)

// Terminal Colours
const (
	COLOR_RESET  = "\033[0m"
	COLOR_RED    = "\033[31m"
	COLOR_GREEN  = "\033[32m"
	COLOR_BLUE   = "\033[34m"
	COLOR_PURPLE = "\033[35m"
	COLOR_CYAN   = "\033[36m"
	COLOR_GRAY   = "\033[37m"
	COLOR_WHITE  = "\033[97m"
)

// Configuration information for the application
type config struct {
	// What port to expose the app on
	port int
	// Environment: `production` or `development`
	env string
	// Database information
	db struct {
		// Database URI
		dsm string
	}
	// Stripe Credentials
	stripe struct {
		key    string
		secret string
	}
}

// Receiver for various parts of the application
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	// Log the Starting HTTP Server to console:
	app.infoLog.Printf("Starting backend server in %s mode on %s:%d%s\n", app.config.env, COLOR_GREEN+"http://localhost", app.config.port, COLOR_RESET)

	return srv.ListenAndServe()
}

func main() {
	// Create a variable
	var cfg config

	// Define command line flags
	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development | production | maintenance}")

	flag.Parse()

	// Get environment variables and assign to the config
	// cfg.stripe.key = os.Getenv("STRIPE_KEY")
	// cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	cfg.stripe.key = "pk_test_UmFFBynPTZ3lwp7kR2YLzomE00y5SWtfFZ"
	cfg.stripe.secret = "sk_test_zRtqf4nLwKSJw5z6es21nkaz007Fr90xeS"

	// Setup Loggers
	infoLog := log.New(os.Stdout, COLOR_BLUE+"INFO:\t"+COLOR_RESET, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, COLOR_RED+"ERROR:\t"+COLOR_RESET, log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  VERSION,
	}

	err := app.serve()
	if err != nil {
		errorLog.Printf("Backend server at PORT %d failed to start!", app.config.port)
		log.Fatal(err)
	}
}
