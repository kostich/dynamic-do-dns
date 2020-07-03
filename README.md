## What

This code can be used in a environment where you want to self-host behind an dynamic public IPv4 address.

The dynamic-do-dns program will (on every run) check what is the current public IPv4 address (using the icanhazip.com by default) and then update the given list of the existing A records on the Digital Ocean DNS, if needed.

## Building
    go build

## Usage
    MYDOTOKEN="<YOUR-DO-TOKEN"> MYDOMAINS="<COMMA-SEPARATED-LIST-OF-A-RECORDS" MYRESOLVER="https://icanhazip.com" ./dynamic-do-dns
