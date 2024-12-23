FROM golang:1.23.4-alpine

WORKDIR /app

COPY . .

RUN go mod download && go build -o /app/main

EXPOSE 8080

CMD ["/app/main"]