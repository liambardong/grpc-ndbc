package service_test

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	pb "github.com/liambardong/grpc-ndbc/api/proto/v1"
	"github.com/liambardong/grpc-ndbc/internal/service" // Import the service package explicitly
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
}

type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func (m *mockHTTPClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return m.DoFunc(req)
}

func loadTestData(t *testing.T) string {
	projectRoot, err := filepath.Abs("../../")
	if err != nil {
		t.Fatalf("Failed to get project root: %v", err)
	}

	testDataPath := filepath.Join(projectRoot, "testdata", "fixtures", "station_table.txt")

	data, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	return string(data)
}

func setupGRPCServer(t *testing.T) (*grpc.Server, func()) {

	s := grpc.NewServer()
	pb.RegisterStationServiceServer(s, service.NewStationService())

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("Failed to serve: %v", err)
		}
	}()

	return s, func() {
		s.Stop()
		lis.Close()
	}
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestListStation(t *testing.T) {
	testData := loadTestData(t)

	// Server setup for future test cases
	server, cleanup := setupGRPCServer(t)
	defer cleanup()

	// Use the server variable here if needed
	// For example, you could log or check its status
	if server == nil {
		t.Fatal("Expected server to be initialized")
	}

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(testData)),
			}, nil
		},
	}

	http.DefaultTransport = mockClient

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create gRPC client: %v", err)
	}

	defer conn.Close()

	client := pb.NewStationServiceClient(conn)

	tests := []struct {
		name    string
		want    int
		wantErr bool
		check   func(*testing.T, *pb.Station)
	}{
		{
			name:    "Successfully fetch stations",
			want:    2,
			wantErr: false,
			check: func(t *testing.T, station *pb.Station) {
				if station.StationId == "0y2w3" {
					expectations := map[string]string{
						"StationId": "0y2w3",
						"Owner":     "CG",
						"Type":      "Weather Station",
						"Name":      "Sturgeon Bay CG Station, WI",
						"Timezone":  "C",
					}

					for field, expected := range expectations {
						got := ""
						switch field {
						case "StationId":
							got = station.StationId
						case "Owner":
							got = station.Owner
						case "Type":
							got = station.Type
						case "Name":
							got = station.Name
						case "Timezone":
							got = station.Timezone
						}
						if got != expected {
							t.Errorf("Station %s = %v, want %v", field, got, expected)
						}
					}

				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Starting test")
			response, err := client.ListStations(ctx, &emptypb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if response == nil {
				t.Error("ExpectedResponse, got nil")
				return
			}

			if got := len(response.Stations); got != tt.want {
				t.Errorf("ListStations() got %v station, expected %v", got, tt.want)
			}

			for _, station := range response.Stations {
				tt.check(t, station)
			}

		})
	}

}
