# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19.3-alpine3.16

WORKDIR /app

# Download necessary Go modules
COPY . .
EXPOSE 9000
RUN go mod tidy
CMD go run main.go

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...