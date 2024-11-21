FROM golang:1.21.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /service_one

EXPOSE 8080

# Run
CMD ["/service_one"]
