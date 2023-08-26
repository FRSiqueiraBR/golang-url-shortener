FROM golang:1.21 as builder

WORKDIR /app

# Copy Go module files
COPY . .
RUN go mod download

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -v -o ./golang-url-shortener

FROM alpine:3.14.10

EXPOSE 8080

# Copy files from builder stage
COPY --from=builder /app/golang-url-shortener .
COPY --from=builder /app/infra ./infra

RUN apk add build-base gcc musl-dev

# Increase GC percentage and limit the number of OS threads
ENV GOGC 1000
ENV GOMAXPROCS 3

# Run binary
CMD ["/golang-url-shortener"]