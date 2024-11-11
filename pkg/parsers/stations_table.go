package parsers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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
		if len(columns) < 9 {
			log.Printf("Skipping malformed line: %s", line)
			continue
		}

		// Add the station to the list
		stationList = append(stationList, types.Station{
			StationId: columns[0],
			Owner:     columns[1],
			Type:      columns[2],
			Hull:      columns[3],
			Name:      columns[4],
			Payload:   columns[5],
			Location:  columns[6],
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
