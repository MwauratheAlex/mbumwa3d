# Start with a base image containing Go runtime
FROM golang:1.17-alpine as builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download dependencies (add debugging output)
RUN set -x \
    && go mod download

# Copy the entire source code from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -ldflags "-X main.Environment=production" -o ./bin/mbumwa3d ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest  

# Set necessary environment variables
ENV GIN_MODE=release

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/mbumwa3d /bin/mbumwa3d

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/bin/mbumwa3d"]
