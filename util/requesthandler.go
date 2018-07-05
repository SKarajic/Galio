package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetJSON is a simple implementation to get a JSON
//
// parameters:
// - url (string)                  : the JSON
// - headerMap (map[string]string) : a map of Headers
//
// returns: (string) the requested JSON
func GetJSON(url string, headerMap map[string]string) string {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("HTTP request failed: %s\n", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	return string(data)
}
