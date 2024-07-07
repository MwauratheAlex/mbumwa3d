# Use an official Golang image as the base image
FROM golang:1.19 AS build

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install -y nodejs
