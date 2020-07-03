package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/digitalocean/godo"
)

func main() {
	// get the token and the domains from the environment
	token := os.Getenv("MYDOTOKEN")
	domains, err := toPunnycode(os.Getenv("MYDODOMAINS"))
	resolver := os.Getenv("MYRESOLVER")

	if err != nil {
		log.Fatalf("ERROR: invalid MYDOMAINS env variable: %v. Exiting.\n", err)
	} else if token == "" {
		log.Fatalln("ERROR: missing the MYDOTOKEN env variable. Exiting.")
	} else if domains[0] == "" {
		log.Fatalln("ERROR: missing the MYDODOMAINS env variable. Exiting.")
	} else if resolver == "" {
		resolver = "https://icanhazip.com"
		log.Printf("INFO: missing the MYRESOLVER env variable, using '%v' as the resolver.\n", resolver)
	}

	// determine the public ip
	ip, err := publicIP(resolver)
	if err != nil {
		log.Fatalf("ERROR: cannot determine the public IPv4 address: %v.\n", err)
	}

	// update the A records with the new public ip
	client := godo.NewFromToken(token)

	for _, d := range domains {
		ctx := context.TODO()

		// determine the domain
		domain, err := getDomain(d)
		if err != nil {
			log.Fatalf("ERROR: cannot determine the domain: %v.\n", err)
		}

		// get an id for the subdomain
		options := &godo.ListOptions{Page: 1, PerPage: 5000}
		listRec, listResp, err := client.Domains.Records(ctx, domain, options)
		if err != nil {
			log.Fatalf("ERROR: cannot get the subdomain id: err: %v, resp: '%v'.\n", err, listResp)
		}

		id := 0
		subdomain := ""
		data := ""
		for _, item := range listRec {
			if fmt.Sprintf("%v.%v", item.Name, domain) == d {
				id = item.ID
				subdomain = item.Name
				data = item.Data
			}
		}

		// update the subdomain, if needed
		decoded, err := fromPunnycode(d)
		if err != nil {
			decoded = d
		}
		if ip != data {
			updateReq := &godo.DomainRecordEditRequest{
				Type: "A",
				Name: subdomain,
				Data: ip,
				TTL:  60,
			}

			_, resp, err := client.Domains.EditRecord(ctx, domain, id, updateReq)

			if err != nil {
				log.Fatalf("ERROR: cannot update the A record: err: %v, resp: '%v'.\n", err, resp)
			} else {
				log.Printf("INFO: updated the IP for the A record '%v'.\n", decoded)
			}
		} else {
			log.Printf("INFO: the A record '%v' is up to date. Skipping.\n", decoded)
		}
	}
}
