package myip

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/pion/stun/v2"
)

// DiscoverIP returns the public IP address discovered.
func DiscoverIP(server string, v6 bool) (net.IP, error) {
	log.Printf("parse STUN URI: %s", server)
	uri, err := stun.ParseURI(server)
	if err != nil {
		return nil, fmt.Errorf("invalid URI '%s': %w", server, err)
	}

	log.Println("initialize client for STUN server")
	network := "udp4"
	address := net.JoinHostPort(uri.Host, strconv.Itoa(uri.Port))
	if v6 {
		network = "udp6"
	}
	client, err := stun.Dial(network, address)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	log.Println("building binding request with random transaction id")
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	log.Println("sending request to STUN server, waiting for response message")
	ipaddress := net.IP{}
	var responseError error
	if err = client.Do(message, responseCallback(&ipaddress, &responseError)); err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	if responseError != nil {
		return nil, fmt.Errorf("response error: %w", responseError)
	}
	if err := client.Close(); err != nil {
		return nil, fmt.Errorf("failed to close connection: %w", err)
	}
	return ipaddress, nil
}

func responseCallback(ipaddress *net.IP, err *error) func(res stun.Event) {
	return func(res stun.Event) {
		if res.Error != nil {
			*err = fmt.Errorf("failed STUN transaction: %w", res.Error)
			return
		}

		var xorAddr stun.XORMappedAddress
		if getError := xorAddr.GetFrom(res.Message); getError != nil {
			*err = fmt.Errorf("failed to get XOR-MAPPED-ADDRESS: %w", getError)
			return
		}

		log.Printf("got IP address: %s", xorAddr.IP)
		*ipaddress = xorAddr.IP
	}
}
