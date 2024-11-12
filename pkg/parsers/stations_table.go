package parsers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/liambardong/grpc-ndbc/pkg/types"
)

func ParseStationTableResponse(response *http.Response) ([]types.Station, error) {
	stationList := []types.Station{}
	log.Println("Station Table Parsing: Starting")
	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), `"`, `'`)

		if strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line by "|" and trim spaces from each part
		columns := strings.Split(line, "|")
		for i := range columns {
			columns[i] = strings.TrimSpace(columns[i])
		}

		// Skip malformed lines or lines with missing columns
		if len(columns) < 10 {
			log.Printf("Skipping malformed line: %s", line)
			continue
		}

		coord, err := extractLatitudeLongitude(columns[6])
		if err != nil {
			return nil, fmt.Errorf("error extracting coordinates %v", err)
		}
		// Add the station to the list
		stationList = append(stationList, types.Station{
			StationId: columns[0],
			Owner:     columns[1],
			Type:      columns[2],
			Hull:      columns[3],
			Name:      columns[4],
			Payload:   columns[5],
			Latitude:  coord[0],
			Longitude: coord[1],
			TimeZone:  columns[7],
			Forecast:  columns[8],
			Note:      columns[9],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}
	log.Println("Station Table Parsing: Completed")
	return stationList, nil
}

func extractLatitudeLongitude(input string) ([]float64, error) {
	// Regular expression to match latitude and longitude
	re := regexp.MustCompile(`(\d+\.\d+)\s+([NS])\s+(\d+\.\d+)\s+([EW])`)

	// Find matches
	match := re.FindStringSubmatch(input)
	if len(match) == 5 {
		// Parse latitude and longitude values
		lat, _ := strconv.ParseFloat(match[1], 64)
		lon, _ := strconv.ParseFloat(match[3], 64)

		// Adjust based on hemisphere
		if match[2] == "S" {
			lat = -lat
		}
		if match[4] == "W" {
			lon = -lon
		}

		return []float64{lat, lon}, nil
	} else {
		return nil, fmt.Errorf("No match found from regex")
	}
}
