FROM golang:1.19-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["/bin/sh" , "-c", "air -build.cmd \"go build -o tmp/main cmd/main.go\""]