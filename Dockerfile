# Use features from go version 1.16, using golang:1.16 is necessary
FROM golang:1.16

# Open up common web traffic (HTTP/S) ports for use
EXPOSE 80
EXPOSE 443

# Environment variable used by the server
ENV LOG_FILE_PATH="../server.log"
ENV CERTIFICATE_KEY_FILE="privkey.pem"
ENV CERTIFICATE_CHAIN_FILE="fullchain.pem"

# Copy all necessary website code and assets
WORKDIR /go/src
COPY /src .

# Move to the directory containing server code
WORKDIR /go/src/server

# Build the server code
RUN go build server.go

# Run the server
CMD ["/go/src/server/server"]
