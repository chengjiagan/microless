FROM golang:1.20 as build

WORKDIR /go/src/loadbalancer
COPY . .
RUN go mod download
