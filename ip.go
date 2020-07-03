package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// Returns an public IPv4 address of the host
func publicIP(resolver string) (string, error) {
	resp, err := http.Get(resolver)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := strings.TrimSuffix(string(body), "\n")
	if !isValidIP(ip) {
		return "", fmt.Errorf("ip address from the resolver is invalid: '%v'", ip)
	}

	return ip, nil
}

// Checks if the given IP is valid
func isValidIP(host string) bool {
	return net.ParseIP(host) != nil
}
