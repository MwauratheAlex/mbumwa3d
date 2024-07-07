# Use an official Golang image as the base image
FROM golang:1.19 AS build

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Set the working directory
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Install dependencies
RUN --mount=type=cache,id=go-mod,target=/root/.cache/go-mod go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN --mount=type=cache,id=go-build,target=/root/.cache/go-build make build

# Use a minimal base image for the final container
FROM gcr.io/distroless/base-debian10

# Copy the built application binary from the build stage
COPY --from=build /app/bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)

# Set the entry point for the container
ENTRYPOINT ["/usr/local/bin/$(APP_NAME)"]

# Expose the application's port
EXPOSE 3000

# Set the default command
CMD ["$(APP_NAME)"]
