FROM golang:1.26.5

WORKDIR /watchdog

COPY . .

RUN go build -o watchdog ./src/server/test-server.go

CMD ["./watchdog"]
