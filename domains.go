package main

import (
	"fmt"
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

// Converts given string from punnycode to string
func fromPunnycode(domain string) (string, error) {
	var decoded string

	decoded, err := idna.ToUnicode(domain)
	if err != nil {
		err = fmt.Errorf("cannot convert punnycode '%v' to domain: %v", decoded, err)
		return decoded, err
	}

	return decoded, nil
}

// Returns the bare domain from the subdomain+domain string
func getDomain(subdomain string) (string, error) {
	components := strings.Split(subdomain, ".")
	comLen := len(components)

	if comLen == 3 {
		return fmt.Sprintf("%v.%v", components[comLen-2], components[comLen-1]), nil
	}

	return "", fmt.Errorf("invalid subdomain, expected 'sub.domain.tld', got '%v'", subdomain)
}
