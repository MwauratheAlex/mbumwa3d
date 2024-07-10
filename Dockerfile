FROM golang:1.22

# Install make, bash, curl, and other necessary tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    make \
    bash \
    curl \
    gcc \
    libc-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

CMD ["./bin/mbumwa3d"]

EXPOSE 8080
EXPOSE 5432
