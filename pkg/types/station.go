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
	Location  string
	TimeZone  string
	Forecast  string
	Note      string
}

func StationToString(s *Station) {
	log.Printf("Station ID: %s, Owner: %s, Type: %s, Name: %s, Location: %s, Timezone: %s\n",
		s.StationId,
		s.Owner,
		s.Type,
		s.Name,
		s.Location,
		s.TimeZone)
}
