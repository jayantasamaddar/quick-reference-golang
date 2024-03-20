package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	// Build a request to use with AddDefaultData
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
	}

	// Add session context to the request
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "Flash")
	testApp.Session.Put(ctx, "warning", "Warning")
	testApp.Session.Put(ctx, "error", "Error")

	td := testApp.AddDefaultData(&TemplateData{}, r)

	if td.Flash != "Flash" {
		t.Error("Failed to get flash data!")
	}

	if td.Warning != "Warning" {
		t.Error("Failed to get warning data!")
	}

	if td.Error != "Error" {
		t.Error("Failed to get error data!")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	// Build a request to use with IsAuthenticated
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
	}

	// Add session context to the request
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	// Case 1: Test when not authenticated
	auth := testApp.IsAuthenticated(r)

	// if auth true then test failed as there should be nothing
	if auth {
		t.Error("Returns true for authenticated, when it should be false.")
	}

	// Case 2: Test when authenticated
	testApp.Session.Put(ctx, "userID", 1)
	auth = testApp.IsAuthenticated(r)

	// if auth true then test failed as there should be nothing
	if !auth {
		t.Error("Returns false for authenticated, when it should be true.")
	}

}

func TestConfig_render(t *testing.T) {
	pathToTemplates = "./templates"

	// Both a request and a response is needed
	rr := httptest.NewRecorder()

	// Build a request to use with render
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
	}

	// Add session context to the request
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	// Case 1: Test an existing template
	testApp.render(rr, r, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Error("Failed to render page!")
	}

	// Case 2: Test a non-existing template
	// testApp.render(rr, r, "non-existing-page.page.gohtml", &TemplateData{})
	// if rr.Code == 200 {
	// 	t.Error("Page shouldn't have existed, but exists!")
	// }
}
