package ambee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message  string
	Stations Stations
}

type Stations []Station

type Station struct {
	Id      string `json:"_id"`
	PlaceId string
	CO      float64 `json:"CO"`
	NO2     float64 `json:"NO2"`
}

func AirQuality(country string) (Stations, error) {

	url := fmt.Sprintf("https://api.ambeedata.com/latest/by-country-code?countryCode=%v", country)

	headers := http.Header{
		"x-api-key":    []string{"f736d8ce64ef5b3f4d264eb186f4196d89f84e8834b8bfdb9472636389cad671"},
		"Content-Type": []string{"application/json"},
	}

	stations, err := Request(url, "GET", headers)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseData Response
	err = json.Unmarshal(stations, &responseData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return responseData.Stations, nil

}
