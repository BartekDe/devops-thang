FROM golang:1.23.4-alpine@sha256:6c5c9590f169f77c8046e45c611d3b28fe477789acd8d3762d23d4744de69812

WORKDIR /app

COPY . .

RUN go mod download && go build -o /app/main

EXPOSE 8080

CMD ["/app/main"]