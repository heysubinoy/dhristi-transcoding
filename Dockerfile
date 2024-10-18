# Stage 1: Build the Go application
FROM golang:latest AS build

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /server main.go

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Install required packages for ffmpeg, fuse, and s3fs
RUN apk update && apk add --no-cache \
    ffmpeg \
    fuse
RUN apk add git
RUN apk add build-base automake autoconf libxml2-dev fuse-dev curl-dev
RUN apk add s3fs-fuse
RUN mkdir -p /s3

# Copy the S3 credentials file
COPY .credentials .passwd-s3fs
RUN chmod 600 .passwd-s3fs

# Copy the built Go application from the build stage
COPY --from=build /server /server

# Expose the server port
EXPOSE 8080

# Command to run the server
CMD ["/server"]
