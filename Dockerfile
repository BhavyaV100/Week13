# Use an official Golang runtime as a base image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required third-party dependencies into the container
RUN go get -u github.com/gin-gonic/gin

# Build the Go application
RUN go build -o main .

# Expose port 9091 to the outside world
EXPOSE 9091

# Command to run the executable
CMD ["./main"]
