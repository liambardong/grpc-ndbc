package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/liambardong/grpc-ndbc/api/proto/v1"
	"github.com/liambardong/grpc-ndbc/config"
	"github.com/liambardong/grpc-ndbc/pkg/parsers"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StationService struct {
	pb.UnimplementedStationServiceServer
}

func NewStationService() *StationService {
	return &StationService{}
}

func (s *StationService) ListStations(ctx context.Context, _ *emptypb.Empty) (*pb.ListStationsResponse, error) {
	url := fmt.Sprintf("%vstations/station_table.txt", config.NDBC_URL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	//DEBUG reasons
	//log.Println("Response status: ", resp.Status)

	stationsList, err := parsers.ParseStationTableResponse(resp)
	if err != nil {
		log.Println("Error parsing station table:", err)
		return nil, err
	}

	//DEBUG reasons
	//log.Printf("Parsed %d stations", len(stationsList))

	protoStationList := make([]*pb.Station, 0, len(stationsList))

	for _, station := range stationsList {
		protoStationList = append(protoStationList, &pb.Station{
			StationId: station.StationId,
			Owner:     station.Owner,
			Type:      station.Type,
			Hull:      station.Hull,
			Name:      station.Name,
			Payload:   station.Payload,
			Latitude:  station.Latitude,
			Longitude: station.Longitude,
			Timezone:  station.TimeZone,
			Forecast:  station.Forecast,
			Note:      station.Note,
		})
	}
	log.Println(protoStationList)
	return &pb.ListStationsResponse{
		Stations: protoStationList,
	}, nil
}
