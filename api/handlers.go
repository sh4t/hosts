package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
    _ "strconv"
    _ "github.com/gorilla/mux"
)

func Root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hosts api")
}

func HostIndex(w http.ResponseWriter, r *http.Request) {
	// list all the hosts
	// setting content-type because it is a pet peave of mine now..
	// thanks ASM!
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK) // return 200

	// if err := json.NewEncoder(w).Encode(hosts); 
	// err != nil { 
	// 	panic(err)
	// }
}

func HostShow(w http.ResponseWriter, r *http.Request) {
	// listing details about a specific host
	// vars := mux.Vars(r)
	// var hostId int
	// var err error

	// if hostId, err = strconv.Atoi(vars["hostId"]); err != nil {
	// 	panic(err)
	// }

	// host := DbFindHost(hostId)
	// if host.Id > 0 {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	if err := json.NewEncoder(w).Encode(host); err != nil {
	// 		panic(err)
	// 	}
	// 	return
	// }
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func HostCreate(w http.ResponseWriter, r *http.Request) {
	var host Host
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &host); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// h := DbCreateHost(host)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(h); err != nil {
	// 	panic(err)
	// }
}