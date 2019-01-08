package domainsdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SearchForDomain checks the lists of registered domains for names
// containing particular words/phrases/numbers or symbols
func SearchForDomain(query string) (Result, error) {
	var result Result

	resp, err := http.Get(endpoint)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}
