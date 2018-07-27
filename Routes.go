package main

import "net/http"

//Route is a struct which have 4 fields
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a Route Slice
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
}
