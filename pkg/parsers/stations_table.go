package parsers

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/liambardong/grpc-ndbc/pkg/types"
)

func ParseStationTableResponse(response *http.Response) ([]types.Station, error) {
	stationList := []types.Station{}
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue
		}

		fmt.Println(line)

		columns := strings.Split(line, "|")

		stationList = append(stationList, types.Station{
			StationId: columns[0],
			Owner:     columns[1],
			Type:      columns[2],
			Hull:      columns[3],
			Name:      columns[4],
			Payload:   columns[5],
			TimeZone:  columns[6],
			Forecast:  columns[7],
			Note:      columns[8],
		})

	}

	return stationList, nil
}
