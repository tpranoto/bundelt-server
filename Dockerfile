FROM golang:1.15-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bundelt cmd/main.go


FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /app/bundelt .
CMD ["/bundelt"]
