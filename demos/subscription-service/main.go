package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
	"github.com/jayantasamaddar/quick-reference-golang/subscription-service/data"
)

const (
	PORT = 80
)

func main() {
	// Connect to the database
	db := initDB()

	// Create sessions
	session := initSession()

	// Create channels

	// Create waitgroup
	wg := sync.WaitGroup{}

	// Set up the application config
	app := Config{
		Session:       session,
		DB:            db,
		InfoLog:       log.New(os.Stdout, std.PrintC(std.Blue, "INFO:\t"), log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, std.PrintC(std.Red, "ERROR:\t"), log.Ldate|log.Ltime|log.Lshortfile),
		Wait:          &wg,
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorDoneChan: make(chan bool),
	}

	// Set up mailer
	app.Mailer = app.createMail()
	go app.listenForEmail()

	// Listen for signals
	go app.listenForShutdown()

	// Listen for errors
	go app.listenForErrors()

	// Listen for web connections
	app.serve()
}

func (app *Config) listenForShutdown() {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quitCh

	// Gracefully shutdown any running processes
	app.shutdown()

	// Exit
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("Running cleanup tasks...")

	// Block until waitgroup is empty
	app.Wait.Wait()

	app.Mailer.DoneChan <- true
	app.ErrorDoneChan <- true

	app.InfoLog.Println("Closing channels and shutting down processes...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	close(app.ErrorChan)
	close(app.ErrorDoneChan)
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Cannot connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		if connection, err := openDB(dsn); err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	if db, err := sql.Open("pgx", dsn); err != nil {
		return nil, err
	} else {
		err := db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}

func initSession() *scs.SessionManager {
	// We can't put the entire user struct in the session as is. We need to register it as we won't know the type.
	gob.Register(data.User{})
	// set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return redisPool
}

func (app *Config) serve() {
	// start Http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: app.routes(),
	}
	app.InfoLog.Printf("Starting server at Port %d...\n", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func (app *Config) createMail() Mail {
	return Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "Info",
		FromAddress: "info@mycompany.com",
		Wait:        app.Wait,
		ErrorChan:   make(chan error),
		MailerChan:  make(chan Message, 100),
		DoneChan:    make(chan bool),
	}
}

func (app *Config) listenForErrors() {
	for {
		select {
		case <-app.ErrorDoneChan:
			return
		case err := <-app.ErrorChan:
			// Handle errors (In real life: notify a slack channel, send multiple notifications etc.)
			app.ErrorLog.Println(err)
		}
	}
}
