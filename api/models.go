package main
	
import (
//	"time"
)

// defining our host model/struct
// TODO: add sensors and secondary IPv4 and IPv6
type Host struct {
	Id 		int		`json:"id"`
	Name	string	`json:"hostname"`
	Active	bool	`json:"active"`
	Location	string	`json:"location"`
	Provider	string	`json:"provider"`
	PrimaryIp	string	`json:"primary_ip"`
	IataCode	string	`json:"iata_code"`

}

type Hosts []Host

// type Sensor struct {
// Name		string
// 	
//}