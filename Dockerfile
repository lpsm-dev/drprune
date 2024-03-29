FROM golang:1.20.4-alpine3.16 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o drprune ./cmd/drprune/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/drprune ./
ENTRYPOINT [ "/app/drprune" ]
