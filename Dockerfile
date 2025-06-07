FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sureshort

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/sureshort .
COPY --from=builder /app/config .

# Expose port (optional)
EXPOSE 80

# Run the binary
CMD ["./sureshort"]
