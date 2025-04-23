FROM golang:1.24-alpine AS builder
WORKDIR /build
COPY . .
RUN apk add --no-cache git
RUN go mod tidy
RUN go build -ldflags "-s -w" -o chance-proxy .

FROM scratch
COPY --from=builder /build/chance-proxy /app/chance-proxy
ENTRYPOINT ["/app/chance-proxy"]
