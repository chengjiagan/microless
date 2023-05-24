FROM golang:1.20 as build

# using go modules proxy in China
ARG GOPROXY="https://goproxy.cn"

WORKDIR /go/src/loadbalancer
COPY . .
RUN go mod download
