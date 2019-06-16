package service

import "net/http"

//Route represents the structure of each route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a slice of type Route struct
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetAccount",            // Name
		"GET",                   // HTTP method
		"/accounts/{accountId}", // Route pattern
		getAccountHandler,
	},
}
