package main

import (
    "log"
    "net/http"
    "fmt"
    "time"
    "encoding/json"

    "github.com/boltdb/bolt"
)

func main() {
	// start := time.Now()

	// mock our Struct for testing
	host := &Host{
		Name: "box1.sjc1.domain.tld",
		Active: false,
		Location: "San Jose, California, USA",
		Provider: "Some Company",
		PrimaryIp: "1.2.3.4",
		IataCode: "SJC",
		Created: time.Now(),
	}	

	// open our boltdb file
	db, err := bolt.Open("db/hosts.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create a bucket if it does not exist
	db.Update(func(tx *bolt.Tx) error {
    	bucket, err := tx.CreateBucketIfNotExists([]byte("hosts"))
    	if err != nil {
        	return fmt.Errorf("create bucket: %s", err)
    	}

		// generate next available ID
		id, _ := bucket.NextSequence()
		host.Id = int(id)

		buffer, err := json.Marshal(host)
		if err != nil {
			return err
		}

		// persist
		// return bucket.Put(itob(host.ID), buffer)
		return bucket.Put([]byte(host.Created.Format(time.RFC3339)), buffer)
	})

	db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("hosts"))
    b.ForEach(func(k, v []byte) error {
        fmt.Printf("key=%s, value=%s\n", k, v)
        return nil
    })
    return nil
})

	router := NewRouter()
	
	log.Fatal(http.ListenAndServe(":5000", router))
}