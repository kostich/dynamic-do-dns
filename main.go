package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

// Converts given string of domains to punnycode
func toPunnycode(domains string) ([]string, error) {
	var punnycode []string
	for _, d := range strings.Split(domains, ",") {
		domain, err := idna.ToASCII(d)
		if err != nil {
			err = fmt.Errorf("cannot convert domain '%v' to punnycode: %v", d, err)
			return punnycode, err
		}
		punnycode = append(punnycode, domain)
	}

	return punnycode, nil
}

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

func main() {
	// get the token and the domains from the environment
	token := os.Getenv("MYDOTOKEN")
	domains, err := toPunnycode(os.Getenv("MYDODOMAINS"))
	resolver := os.Getenv("MYRESOLVER")

	if err != nil {
		fmt.Printf("Invalid MYDOMAINS env variable: %v. Exiting.\n", err)
		os.Exit(1)
	} else if token == "" {
		fmt.Println("Missing MYDOTOKEN env variable. Exiting.")
		os.Exit(1)
	} else if domains[0] == "" {
		fmt.Println("Missing MYDODOMAINS env variable. Exiting.")
		os.Exit(1)
	} else if resolver == "" {
		resolver = "https://icanhazip.com"
		fmt.Printf("Missing MYRESOLVER env variable, using '%v' as the resolver.\n", resolver)
	}

	// determine the public ip
	ip, err := publicIP(resolver)
	if err != nil {
		fmt.Printf("Cannot determine the public IPv4 address: %v\n", err)
		os.Exit(1)
	}

	// get all domains
	// update the domains with the new public ip

	fmt.Printf("token: %v\ndomains: %v\n", token, domains)
	fmt.Printf("ip: %v\n", ip)
}
