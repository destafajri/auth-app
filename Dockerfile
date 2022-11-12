# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Desta <destafajri@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/auth-app

# Download necessary Go modules
COPY . .
EXPOSE 9000
RUN go mod tidy
RUN go build
CMD go run main.go

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...

# Start a new stage from scratch
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
# COPY --from=builder $GOPATH/src/auth-app/main .
# COPY --from=builder $GOPATH/src/auth-app/.env .       

# # Expose port 9000 to the outside world
# EXPOSE 9000

# #Command to run the executable
# CMD ["./main"]