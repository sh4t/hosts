package main

import (
    "net/http"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route { "Index", "GET", "/", Root, },
	Route { "HostIndex", "GET", "/hosts", HostIndex, },
	Route { "HostShow", "GET", "/hosts/{hostId}", HostShow, },
	Route { "HostCreate", "POST", "/hosts", HostCreate, },
}