package main

import (
    "log"
    "net/http"

    "github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("db/hosts.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	router := NewRouter()
	
	log.Fatal(http.ListenAndServe(":5000", router))
}