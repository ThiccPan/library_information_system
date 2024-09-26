#syntax=docker/dockerfile:1

# stage 1
FROM golang:1.23 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go
EXPOSE 8080
CMD ["/app/main"]

# stage 2
FROM alpine:3
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]