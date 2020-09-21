FROM golang:1.15.2-alpine3.12 AS builder
WORKDIR /app/
COPY go.* main.go ./
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder app/app .
EXPOSE 8080
CMD ["./app"]
