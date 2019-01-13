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

	url := endpoint + query
	resp, err := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}
