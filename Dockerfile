FROM golang:1.24 AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/bin/app ./cmd
RUN chmod +x /build/bin/app

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /build/bin/app ./app
CMD ["/app/app"]

