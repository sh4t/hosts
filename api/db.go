// this is a completely fake database for now

package main

import (
	"fmt"
)

var currentId int
var hosts Hosts

// Seed Data
func init() {
	DbCreateHost(
		Host{ 
			Name: "host1.region.domain.tld",
			Active: true,
			Location: "Beauharnois, Quebec, Canada",
			Provider: "company.tld",
			PrimaryIp: "10.10.10.10",
			IataCode: "BHS",
		})
	DbCreateHost(
		Host{ 
			Name: "host1.region.domain.tld",
			Active: true,
			Location: "Hong Kong, China",
			Provider: "company.tld",
			PrimaryIp: "11.11.11.11",
			IataCode: "HKG",
		})
	DbCreateHost(
		Host{ 
			Name: "host1.region.domain.tld",
			Active: true,
			Location: "Ho Chi Minh City, Vietnam",
			Provider: "company.tld",
			PrimaryIp: "12.12.12.12",
			IataCode: "SGN",
		})
	DbCreateHost(
		Host{ 
			Name: "host1.region.domain.tld",
			Active: true,
			Location: "Sofia, Bulgaria",
			Provider: "company.tld",
			PrimaryIp: "13.14.15.16",
			IataCode: "SOF",	
		})
}

// find a single host
func DbFindHost(id int) Host {
	for _, h := range hosts {
		if h.Id == id {
			return h
		}
	}
	// empty
	return Host{}
}

func DbCreateHost(h Host) Host {
	currentId += 1
	h.Id = currentId
	hosts = append(hosts, h)
	return h
}

func DbDestroyHost(id int) error {
	for i, h := range hosts {
		if h.Id == id {
			hosts = append(hosts[:i], hosts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Unable to find Host with id of %d to delete", id)
}




// func (s *Store) CreateHost(h *Host) error {
// 	return s.db.Update(func(tx *bolt.Tx) error{
// 		// retrieve hosts bucket
// 		bucket := tx.Bucket([]byte("hosts"))

// 		// generate next available ID
// 		// ignore err check as it cannot happen in an Update() per doc
// 		id, _ = bucket.NextSequence()
// 		h.ID = int(id)

// 		buffer, err := json.Marshal(h)
// 		if err != nil {
// 			return err
// 		}

// 		// persist
// 		return bucket.Put(itob(h.ID), buffer)
// 	})
// }

// // itob returns an 8-byte big endian representation of v.
// func itob(v int) []byte {
//     b := make([]byte, 8)
//     binary.BigEndian.PutUint64(b, uint64(v))
//     return b
// }
