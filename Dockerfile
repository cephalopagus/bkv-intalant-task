FROM golang:1.25.7-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/app/main.go

FROM scratch
COPY --from=builder /app/main .
CMD ["/main"]
