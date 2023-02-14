FROM golang:1.20.1-alpine3.16 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o drprune ./cmd/drprune/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/drprune ./
ENTRYPOINT [ "/app/drprune" ]
