// // this is a completely fake database for now
// this is going to become our boltdb wrapper
package main

// import (
// 	"fmt"
// 	"os"
// 	"github.com/boltdb/bolt"
// )
// func (self *Store) Open() (err error) {
// 	db, err := bolt.Open("db/hosts.db", 0644, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// }

// // mock our Struct for testing
// 	// host := &Host{
// 	// 	Name: "box1.sjc1.domain.tld",
// 	// 	Active: false,
// 	// 	Location: "San Jose, California, USA",
// 	// 	Provider: "Some Company",
// 	// 	PrimaryIp: "1.2.3.4",
// 	// 	IataCode: "SJC",
// 	// 	Created: time.Now(),
// 	// }	


// 	// create a bucket if it does not exist
// 	db.Update(func(tx *bolt.Tx) error {
//     	bucket, err := tx.CreateBucketIfNotExists([]byte("hosts"))
//     	if err != nil {
//         	return fmt.Errorf("create bucket: %s", err)
//     	}

// 		// generate next available ID
// 		id, _ := bucket.NextSequence()
// 		temp := int(id)
// 		s := strconv.Itoa(temp)
// 		host.Id = temp
// 		buffer, err := json.Marshal(host)
// 		if err != nil {
// 			return err
// 		}

// 		// binaryId := make([]byte, 8)
// 	 	// binary.BigEndian.PutUint64(binaryId, uint64(host.Id))

// 		// persist, id as key
// 		return bucket.Put([]byte(s), buffer)

// 		// persist with date as key
// 		// return bucket.Put([]byte(host.Created.Format(time.RFC3339)), buffer)
// 	})

// 	db.View(func(tx *bolt.Tx) error {
// 	    bucket := tx.Bucket([]byte("hosts"))
// 	    bucket.ForEach(func(k, v []byte) error {
// 	        fmt.Printf("key=%s, value=%s\n", k, v)
// 	        return nil
// 	    })
// 	    return nil
// 	})