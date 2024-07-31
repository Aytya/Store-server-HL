FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o e-commerce ./cmd/main.go

CMD ["./e-commerce"]