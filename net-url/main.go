package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://arisahyper.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scheme: %s\n", u.Scheme)
	fmt.Printf("Opaque: %s\n", u.Opaque)
	fmt.Printf("User: %s\n", u.User)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)
	fmt.Printf("RawPath: %s\n", u.RawPath)
	fmt.Printf("ForceQuery: %s\n", u.ForceQuery)
	fmt.Printf("RawQuery: %s\n", u.RawQuery)
	fmt.Printf("Fragment: %s\n", u.Fragment)
}
