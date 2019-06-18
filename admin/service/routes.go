package service

import "net/http"

//Route represents the structure of each route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  middleware
}

//Routes is a slice of type Route struct
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"SignIn",  // Name
		"POST",    // HTTP method
		"/submit", // Route pattern
		SubmitHandler,
		nil,
	},
	Route{
		"SignUp",
		"POST",
		"/delete",
		DeleteHandler,
		nil,
	},
	Route{
		"GetQuestions",
		"GET",
		"/get-questions",
		nil,
		nil,
	},
}
