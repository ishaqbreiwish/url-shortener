# Use the official Go image as a base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the application port
EXPOSE 8080


CMD ["./main"]
