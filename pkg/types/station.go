package types

import (
	"log"
)

type Station struct {
	StationId string
	Owner     string
	Type      string
	Hull      string
	Name      string
	Payload   string
	Latitude  float64
	Longitude float64
	TimeZone  string
	Forecast  string
	Note      string
}

func StationToString(s *Station) {
	log.Printf("Station ID: %s, Owner: %s, Type: %s, Name: %s, Latitude: %s, Longitude: %s, Timezone: %s\n",
		s.StationId,
		s.Owner,
		s.Type,
		s.Name,
		s.Latitude,
		s.Longitude,
		s.TimeZone)
}
