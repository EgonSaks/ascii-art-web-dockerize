# This Dockerfile builds an image for a Go application
FROM golang:alpine as builder

# Install zsh and bash
RUN apk add --no-cache zsh
RUN apk add --no-cache bash

# Set metadata for the image
LABEL version="1.0" maintainer="Egon Saks (egonsaks@gmail.com)"

# Create a new directory for the app and copy the code into it
RUN mkdir /app
ADD . /app
WORKDIR /app

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Run the main binary when the container starts
CMD ["/app/main"]
