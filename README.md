NDBC TXT Data gRPC Service
A Go-based gRPC service that provides programmatic access to National Data Buoy Center (NDBC) text data files. This service parses and serves structured data from NDBC's text files available at https://www.ndbc.noaa.gov/data/.
Features

Access to NDBC station data from standardized text files
Supported data types:

Standard meteorological data
Detailed wave data
Prediction data
Historical data


Automatic parsing of NDBC's fixed-width text format
Structured data responses using Protocol Buffers
Built with Go and modern cloud-native practices

Prerequisites

Go 1.21 or higher
Protocol Buffers compiler (protoc)
Go gRPC tools (go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)
Docker (optional, for containerized deployment)  
