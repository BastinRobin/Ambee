package main

import (
	"fmt"
	"log"

	"github.com/BastinRobin/Ambee/ambee"
)

func main() {

	stations, err := ambee.AirQuality("IN")
	if err != nil {
		log.Println(err)
	}

	for _, station := range stations {
		fmt.Println(station.CO, station.Id, station.PlaceId, station.NO2)
	}
}
