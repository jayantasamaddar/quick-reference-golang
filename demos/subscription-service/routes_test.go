package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

// Test if all routes exist
func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(*chi.Mux)

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

// Test if a route exists
func routeExists(t *testing.T, routes chi.Router, route string) {
	var found bool

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("Did not find '%s' in registered routes", route)
	}
}
