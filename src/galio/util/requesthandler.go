package util

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func GetJSON(url string, m map[string]string) string {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	for key, value := range m {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("HTTP request failed: %s\n", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	return string(data)
}