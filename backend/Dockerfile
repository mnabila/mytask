# Build stage
FROM golang:1.20.4-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.RELEASE=YES" -o mytask

# Release stage
FROM  alpine:3.18
WORKDIR /app

COPY --from=builder /app/mytask .

EXPOSE 8080
CMD [ "/app/mytask" ]
