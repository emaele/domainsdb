# domainsdb
Domainsdb API wrapper written in Go

## Usage
```package main

import (
	"fmt"
	"log"

	"github.com/emaele/domainsdb"
)

func main() {
	a, err := domainsdb.SearchForDomain("shitposting.io")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v", a.Domains[0])
}```
