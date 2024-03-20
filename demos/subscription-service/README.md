# Table of Contents

- [Table of Contents](#table-of-contents)
- [About](#about)
- [How to Use](#how-to-use)
- [Workflow](#workflow)
- [Tech Stack](#tech-stack)
  - [Packages Used](#packages-used)
  - [Other services](#other-services)
- [Implementation Details](#implementation-details)
  - [Graceful Shutdown](#graceful-shutdown)
  - [Implementing Asynchronous Email Sending](#implementing-asynchronous-email-sending)
  - [Testing](#testing)

---

# About

A Subscription service where an user can register for an account and then purchase from a list of subscriptions - Bronze, Silver, Gold.

**Features**:

- Notification Emails on User Registration with Activation Link.
- Session Management.
- Protected Plans route only visible to logged in user.
- On Subscription to a plan, receive an email invoice and a manual generated for the plan.
- PDF generation and emails, are sent concurrently and implemented via goroutines.
- Graceful shutdown that waits for all asynchronous processes (goroutines) to complete before shutting the main function.

---

# How to Use

1. Run the docker containers using `docker-compose up`
2. Use a PostgreSQL Client like PGAdmin to run the database migrations in `./db.sql`
3. Run `make start`
4. The Mailhog Email receiver client is available on port `8025`

---

# Workflow

- User Flow
  - User Registers.
  - Gets an Email with an Activation Link (Signed URL)
  - Can view the page where the available subscriptions are displayed

---

# Tech Stack

## Packages Used

| Package                                                                       | Use                                                                       |
| ----------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| [`pgx/v4`](https://github.com/jackc/pgx)                                      | PostgreSQL driver and toolkit                                             |
| [`scs`](https://github.com/alexedwards/scs)                                   | HTTP session management for Go                                            |
| [`scs/redisstore`](https://github.com/alexedwards/scs/tree/master/redisstore) | A Redis based session store for SCS                                       |
| [`go-chi/chi/v5`](https://github.com/go-chi/chi)                              | HTTP Router                                                               |
| [`go-premailer`](https://github.com/vanng822/go-premailer)                    | Inline styling for HTML mail in Go                                        |
| [`go-simple-mail](https://github.com/xhit/go-simple-mail)                     | Send emails in Go with SMTP Keep Alive and Timeout for Connect and Send   |
| [`go-alone`](https://github.com/bwmarrin/go-alone)                            | Generate and Verify signed text                                           |
| [`sweetalert2`](https://sweetalert2.github.io/)                               | JavaScript library for showing alerts to be used with the plans template. |
| [`gofpdf`](https://github.com/phpdave11/gofpdi)                               | Create a PDF / Import an existing PDF into a new PDF.                     |
| [`gofpdf/contrib/gofpdi`](https://github.com/phpdave11/gofpdi)                | Open an existing PDF and use it as a template.                            |

---

## Other services

| Service                                       | Use                                         |
| --------------------------------------------- | ------------------------------------------- |
| [mailhog](https://github.com/mailhog/MailHog) | Fake SMTP Server for Testing Emails locally |
| Postgres                                      | Database                                    |
| Redis                                         | Session Storage                             |

---

# Implementation Details

## Graceful Shutdown

The application will have a number of goroutines running in the background.
Some of those goroutines are just going to be listening to channels and some are actually going to be doing something like sending an email or generating invoice or something like that.

At some point, we may decide to stop this application, for whatever reason - maybe we need to do some work on the server, or you need to implement a hotfix or something like that.

If you just stop it, by typing, `make stop`, everything just stops. Any running goroutines just die without finishing.

That's not good because you might not send an email that needs to go out or generate an invoice or whatever the case may be and that process might still be running.

Hence, we need to implement graceful shutdown.

```go
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

	app.InfoLog.Println("Closing channels and shutting down processes...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
}
```

---

## Implementing Asynchronous Email Sending

Sending Email can slow the application down, because we are dependent on a mail service provider. Hence we need to run it using goroutines in the background.

**Using channels**

- Email sending is handled with two channels
  - Send information to a channel. Upon receiving, an email sending goroutine will be started in the background
  - The other channel will be listening for errors.
- Adding cleanup tasks on app shutdown is handled with one channel
  - The third channel is used to shut things down

---

## Testing

When we are testing web applications, we have to duplicate the environment, the various parts of our application run in. This is particularly true for handlers but also true for most parts of the application.

- `setup_tests.go`: Sets up the environment. Runs before the tests run.
- `routes_test.go`: Test if routes exist.
- `render_test.go`: Test the render functions - `AddDefaultData`, `IsAuthenticated` and the `render` function itself.
- `data/test-models.go`: Setup models to be used for testing instead of the database connection. Uses `data/interface.go`.
- `handlers_test.go`: Test the handlers.

**Test commands**:

```go
// Outputs a coverage.out file containing the current coverage
go test -coverprofile=coverage.out -v

// View on the browser, the exact coverage shown in green for covered, red for uncovered
go tool cover -html=coverage.out

// Test including race conditions and return verbose results
go test -v -race .
```

---
