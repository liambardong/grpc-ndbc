package service

import (
	"context"
	"fmt"
	"net/http"

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

func (s *StationService) ListStation(ctx context.Context, _ *emptypb.Empty) (*pb.ListStationsResponse, error) {
	url := fmt.Sprintf("%sstations/station_table.txt", config.NDBC_URL)

	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	stationsList, err := parsers.ParseStationTableResponse(resp)

	for _, station := range stationsList {

	}
}
