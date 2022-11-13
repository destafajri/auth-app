# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Desta <destafajri@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src

# Download necessary Go modules
COPY . .
EXPOSE 9000
RUN go mod tidy
RUN go build
CMD go run main.go

# # Stage 2
# FROM builder AS service_builder

# WORKDIR /usr/app

# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o bin

# # Stage 3
# FROM alpine:latest  

# RUN apk --no-cache add ca-certificates tzdata
# WORKDIR /root/

# RUN mkdir -p /root/fetch-app
# COPY --from=service_builder /usr/app/bin bin
# COPY --from=service_builder /usr/app/.env .env
# COPY --from=service_builder /usr/app/fetch-app /root/fetch-app

# ENTRYPOINT ["./bin"]