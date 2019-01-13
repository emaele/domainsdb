# domainsdb
[![GoDoc](https://godoc.org/github.com/emaele/domainsdb?status.svg)](https://godoc.org/github.com/emaele/domainsdb)


Domainsdb API wrapper written in Go

## Usage
```go
package main

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
}
```
