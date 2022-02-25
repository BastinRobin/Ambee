package ambee

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Module for calling any api passing `url`, `headers`, `method`
func Request(url string, method string, headers http.Header) ([]byte, error) {

	// Create a client to call the request
	client := http.Client{}

	// Create NewRequest structure
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create request headers
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
