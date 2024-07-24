FROM golang:1.22-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -ldflags "-X main.Environment=docker-prod" -o ./bin/mbumwa3d ./cmd/main.go

EXPOSE 3000

CMD ["./bin/mbumwa3d"]
