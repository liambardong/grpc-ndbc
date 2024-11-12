# NDBC gRPC Service

A Go-based gRPC service that provides programmatic access to National Data Buoy Center (NDBC) text data files. This service parses and serves structured data from NDBC's text files available at `https://www.ndbc.noaa.gov/data/`.

## Features

- Access to NDBC station data from standardized text files
- Supported data types:
  - Standard meteorological data
  - Detailed wave data
  - Prediction data
  - Historical data
- Automatic parsing of NDBC's fixed-width text format
- Structured data responses using Protocol Buffers
- Built with Go and modern cloud-native practices

## Prerequisites

- Go 1.21 or higher
- Protocol Buffers compiler (protoc)
- Go gRPC tools (`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`)
- Docker (optional, for containerized deployment)

## Installation

```bash
# Clone the repository
git clone https://github.com/liambardong/grpc-ndbc
cd grpc-ndbc

# Install dependencies
go mod download

# Generate Protocol Buffer code
make generate
```

## Usage

### Starting the Server

```bash
go run cmd/server/main.go
```

The server will start on port 50051 by default.

### Example Client Usage

```go
package main

import (
    "context"
    "log"
    pb "github.com/liambardong/grpc-ndbc/"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()
    
    log.Printf("Station data: %v", response)
}
```

## Data Types and Endpoints

The service supports the following NDBC text data types:



## API Reference

### Service Methods

- `StationList() returns (StationListResponse)`
  
See `api/proto/v1` for detailed message definitions.

## Configuration

The service can be configured using environment variables:

```env

```

## Docker Support

```bash
To Be Created
```

## Error Handling

The service handles common error cases:
- Invalid station IDs
- Missing or malformed text files
- Network connectivity issues with NDBC servers
- Malformed data in text files

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- National Data Buoy Center (NDBC) for providing open access to their data
- NDBC data is available at https://www.ndbc.noaa.gov/data/
- gRPC team for the excellent framework
- Go team for the amazing language and tools
