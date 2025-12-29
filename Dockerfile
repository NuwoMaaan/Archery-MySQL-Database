FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . . 
RUN go mod download
RUN go build -o ./app ./cmd/app
 
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/app ./app

CMD ["/app/app"]

