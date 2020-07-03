## What

This code can be used in a environment where you want to self-host behind an dynamic public IPv4 address.

The dynamic-do-dns program will (on every run) check what is the current public IPv4 address (using the icanhazip.com by default) and then update the given list of the existing A records on the Digital Ocean DNS, if needed.

## Building
    # Locally
    go build

    # Docker
    for arch in amd64 arm64; do docker build -t "dynamic-do-dns:<VER>-$arch" --build-arg ARCH=$arch .; done

## Usage
    MYDOTOKEN="<YOUR-DO-TOKEN"> MYDOMAINS="<COMMA-SEPARATED-LIST-OF-A-RECORDS" MYRESOLVER="https://icanhazip.com" ./dynamic-do-dns
