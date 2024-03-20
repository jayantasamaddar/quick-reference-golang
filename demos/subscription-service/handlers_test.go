package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jayantasamaddar/quick-reference-golang/subscription-service/data"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "HomePage",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
		// expectedHTML:       "failed expected HTML",
	},
	{
		name:               "LoginPage",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "Logout",
		url:                "/logout",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.Logout,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Pages(t *testing.T) {
	pathToTemplates = "./templates"

	for _, entry := range pageTests {
		// Create a response recorder (for mimicking a http.ResponseWriter)
		rr := httptest.NewRecorder()

		// Create a request to mimick the http.Request
		r, err := http.NewRequest("GET", entry.url, nil)
		if err != nil {
			log.Println(err)
		}
		// Add session context to the request
		ctx := getCtx(r)
		r = r.WithContext(ctx)

		if len(entry.sessionData) > 0 {
			for key, val := range entry.sessionData {
				testApp.Session.Put(ctx, key, val)
			}
		}

		// Serve page
		entry.handler.ServeHTTP(rr, r)

		// Perform Test
		if rr.Code != entry.expectedStatusCode {
			t.Errorf("%s Test failed: Expected %d but got %d", entry.name, entry.expectedStatusCode, rr.Code)
		}

		// HTML expected to find in page
		if len(entry.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, entry.expectedHTML) {
				t.Errorf("%s Test failed: Failure to find '%s' as expected HTML", entry.name, entry.expectedHTML)
			}
		}
	}
}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates"

	postedData := url.Values{
		"email":    {"admin.example.com"},
		"password": {"password"},
	}

	// Create a response recorder (for mimicking a http.ResponseWriter)
	rr := httptest.NewRecorder()

	// Create a request to mimick the http.Request
	r, err := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	if err != nil {
		log.Println(err)
	}
	// Add session context to the request
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	// create Handler
	handler := http.HandlerFunc(testApp.PostLoginPage)

	// Serve HTTP
	handler.ServeHTTP(rr, r)

	// Tests

	// Check for the response
	if rr.Code != http.StatusSeeOther {
		t.Error("PostLoginPage Test failed: Wrong code returned")
	}

	// Check for session:
	// Case 1: Session doesn't exist
	if !testApp.Session.Exists(ctx, "userID") {
		t.Error(`PostLoginPage Test failed: Did not find "userID" in session`)
	}
}

func TestConfig_SubscribeToPlan(t *testing.T) {
	// Create a response recorder (for mimicking a http.ResponseWriter)
	rr := httptest.NewRecorder()

	// Create a request to mimick the http.Request
	r, err := http.NewRequest("GET", "/subscribe?id=1", nil)
	if err != nil {
		log.Println(err)
	}
	// Add session context to the request
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	// Mock add an user to the session
	testApp.Session.Put(ctx, "user", data.User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "AdminFirst",
		LastName:  "AdminLast",
		Active:    1,
	})

	// create Handler
	handler := http.HandlerFunc(testApp.SubscribeToPlan)

	// Serve HTTP
	handler.ServeHTTP(rr, r)

	// Test

	// Check for the response
	if rr.Code != http.StatusSeeOther {
		t.Error("PostLoginPage Test failed: Wrong code returned")
	}
	// At this point goroutines won't finish

	testApp.Wait.Wait()

}
