package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
	"github.com/jayantasamaddar/quick-reference-golang/subscription-service/data"
)

var testApp Config

// Special function in Go that runs our tests for us
func TestMain(m *testing.M) {
	gob.Register(data.User{})

	tempPath = "./tmp"
	pathToManual = "./pdf"

	// set up session
	session := scs.New()
	// no need to test redis
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	// Set up the test application config specific to tests
	testApp = Config{
		Session:       session,
		DB:            nil,
		InfoLog:       log.New(os.Stdout, std.PrintC(std.Blue, "INFO:\t"), log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, std.PrintC(std.Red, "ERROR:\t"), log.Ldate|log.Ltime|log.Lshortfile),
		Models:        data.TestNew(nil), // Get from data/test-models.go
		Wait:          &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorDoneChan: make(chan bool),
	}

	// create a dummy mailer
	testApp.Mailer = Mail{
		Wait:       testApp.Wait,
		ErrorChan:  make(chan error),
		MailerChan: make(chan Message, 100),
		DoneChan:   make(chan bool),
	}

	go func() {
		for {
			select {
			case <-testApp.Mailer.MailerChan:
				testApp.Wait.Done()
			case <-testApp.Mailer.ErrorChan:
			case <-testApp.Mailer.DoneChan:
				return
			}
		}
	}()

	// Listen for errors
	go func() {
		for {
			select {
			case <-testApp.ErrorDoneChan:
				return
			case err := <-testApp.ErrorChan:
				// Handle errors (In real life: notify a slack channel, send multiple notifications etc.)
				testApp.ErrorLog.Println(err)
			}
		}
	}()

	// Run our tests
	os.Exit(m.Run())

}

// Get session information for tests
func getCtx(r *http.Request) context.Context {
	ctx, err := testApp.Session.Load(r.Context(), r.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx
}
