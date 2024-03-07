FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod downliad

COPY . .

RUN go build -o main .

CMD ["./main"]
