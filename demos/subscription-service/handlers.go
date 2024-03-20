package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/subscription-service/data"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

var pathToManual = "./pdf"
var tempPath = "./tmp"

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	// Renew the token once the user logs in
	_ = app.Session.RenewToken(r.Context())

	// Parse form post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// Get the user from the database using email address
	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials")
		// Redirect user
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	if validPassword, err := app.Models.User.PasswordMatches(password); err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials")
		// Redirect user
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		if !validPassword {
			msg := Message{
				To:      email,
				Subject: "Failed logged in attempt",
				Data:    "Invalid login attempt",
			}
			app.sendEmail(msg)

			app.Session.Put(r.Context(), "error", "Invalid credentials")
			// Redirect user
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Log user in
		app.Session.Put(r.Context(), "userID", user.ID)
		app.Session.Put(r.Context(), "user", user)

		app.Session.Put(r.Context(), "flash", "Logged in successfully!")

		// redirect the user
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	// clean up session
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	// (1) Create an user
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// (2) Validate Data (not done here): Ensure user is not registered, all required fields are filled, etc.

	// (3) Create user
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0, // defaults to 0 anyway
		IsAdmin:   0, // defaults to 0 anyway
	}

	_, err = app.Models.User.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to create user")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// (4) Send activation email - can be expensive in terms of processing time and delay, need to be asynchronous (implemented with goroutines)
	url := fmt.Sprintf("http://localhost:%d/activate?email=%s", PORT, u.Email)
	signedURL := GenerateTokenFromString(url)
	app.InfoLog.Println(signedURL)

	// Create email message
	msg := Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedURL),
	}
	app.sendEmail(msg)
	app.Session.Put(r.Context(), "flash", "Confirmation e-mail sent. Check your e-mail.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Get Handler that activates registration (user clicks activation link)
func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// (1a) Validate URL
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost:%d%s", PORT, url)

	if !VerifyToken(testURL) {
		app.Session.Put(r.Context(), "error", "Invalid Token!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// (1b) Activate the account
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	err = app.Models.User.Update(*u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to activate user!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Account activated. You can now log in!")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	// Render template
	dataMap := make(map[string]any)
	dataMap["plans"] = plans
	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}

func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	// (1a) Get the ID of the plan as chosen
	q := r.URL.Query().Get("id")
	if len(q) == 0 {
		app.ErrorLog.Println("Failed to get Plan ID")
		http.Redirect(w, r, "/members/plans", http.StatusBadRequest)
		return
	}

	// (1b) Correctly typecast ID received as a valid integer
	id, err := strconv.Atoi(q)
	if err != nil {
		app.ErrorLog.Println("ID must be numeric")
		http.Redirect(w, r, "/members/plans", http.StatusBadRequest)
		return
	}

	// (2) Get the plan from the database and check if it exists
	plan, err := app.Models.Plan.GetOne(id)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to find Plan")
		app.ErrorLog.Printf("Plan with ID: %d not found!\n", id)
		http.Redirect(w, r, "/members/plans", http.StatusBadRequest)
		return
	}

	// (3) Get the user from the session
	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.Session.Put(r.Context(), "error", "Unable to find User")
		app.ErrorLog.Println("User not found!")
		http.Redirect(w, r, "/members/plans", http.StatusBadRequest)
		return
	}

	/*******************************************************************/
	// Concurrent operations
	/*******************************************************************/

	// (4) Generate an Invoice and email it
	app.Wait.Add(1)
	go func() {
		defer app.Wait.Done()

		invoice, err := app.getInvoice(user, plan)
		if err != nil {
			// send this to the ErrorChan channel
			app.ErrorChan <- err
		}

		// Send Email
		msg := Message{
			To:       user.Email,
			Subject:  "Your invoice",
			Data:     invoice,
			Template: "invoice",
		}

		app.sendEmail(msg)
	}()

	// (6) Generate a Manual (subscription collateral to be sent on subscription) and send an email
	// - Open an existing PDF
	// - Write some information to it
	// - Send the PDF to the user

	app.Wait.Add(1)
	go func() {
		defer app.Wait.Done()

		pdf := app.generateManual(user, plan)

		// Write to a temporary folder at the root level of the application
		err := pdf.OutputFileAndClose(fmt.Sprintf("%s/%d_manual.pdf", tempPath, user.ID))
		if err != nil {
			app.ErrorChan <- err
			return
		}

		// Send Email with Manual
		msg := Message{
			To:      user.Email,
			Subject: "Your manual",
			Data:    "Your user manual is attached.",
			AttachmentMap: map[string]string{
				"manual.pdf": fmt.Sprintf("%s/%d_manual.pdf", tempPath, user.ID),
			},
		}

		app.sendEmail(msg)
	}()

	// (7) Subscribe user to a Plan
	err = app.Models.Plan.SubscribeUserToPlan(user, *plan)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Error subscribing to plan!")
		app.ErrorLog.Println("Error subscribing to plan!")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	// (8) Update session user with current updated user from database
	u, err := app.Models.User.GetOne(user.ID)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Error getting user from database!")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "user", u)

	// (9) Redirect
	app.Session.Put(r.Context(), "flash", "Subscribed!")
	http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
}

// Dummy function
func (app *Config) getInvoice(u data.User, plan *data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}

// Generate a manual
func (app *Config) generateManual(u data.User, plan *data.Plan) *gofpdf.Fpdf {
	// orientationStr: "P" for portrait
	// unitStr: "mm" for millimetres (unit of measurement)
	// sizeStr: "Letter" (for size)
	// fontDirStr: "" for default
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)

	// Import a PDF
	importer := gofpdi.NewImporter()

	// Simulate the time needed to create a PDF (to demonstrate asynchronousity)
	time.Sleep(5 * time.Second)

	templateID := importer.ImportPage(pdf, fmt.Sprintf("%s/manual.pdf", pathToManual), 1, "/MediaBox")

	// Add a new page to the pdf
	pdf.AddPage()
	importer.UseImportedTemplate(pdf, templateID, 0, 0, 215.9, 0)

	// Set X and Y coordinates (You do this by measuring where you want something to appear on a page)
	pdf.SetX(75)
	pdf.SetY(150)

	pdf.SetFont("Arial", "", 12)

	// 0: Width
	// 4: Height: Spacing between lines
	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", u.FirstName, u.LastName), "", "C", false)

	// Apply line break
	pdf.Ln(5)

	pdf.MultiCell(0, 4, fmt.Sprintf("%s User Guide", plan.PlanName), "", "C", false)

	return pdf
}
