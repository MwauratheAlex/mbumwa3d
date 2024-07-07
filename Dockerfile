# Use an official Golang image as the base image
FROM golang:1.19 AS build

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Set the working directory
WORKDIR /app

# Copy the built application binary from the build stage
COPY --from=build /app/bin/mbumwa3d /usr/local/bin/mbumwa3d

# Set the entry point for the container
ENTRYPOINT ["/usr/local/bin/mbumwa3d"]

# Expose the application's port
EXPOSE 3000

