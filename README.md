## What

This code can be used in a environment where you want to self-host behind an dynamic public IPv4 address.

The dynamic-do-dns program will (on every run) check what is the current public IPv4 address (using icanhazip.com) and then update the given list of the existing domains on the Digital Ocean DNS, if needed.
