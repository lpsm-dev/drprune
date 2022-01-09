FROM golang:alpine3.15

COPY . .

RUN go mod tidy
