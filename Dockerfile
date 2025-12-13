# Build stage
 FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# Run stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server /server
COPY --from=builder /app/results.json /app/results.json

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/server"]
