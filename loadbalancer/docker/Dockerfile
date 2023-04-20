FROM golang:1.18 as build
# using go modules proxy in China
ARG GOPROXY="https://goproxy.cn"

WORKDIR /go/src
COPY . .
RUN go mod download
# build load balancer, disable cgo
RUN CGO_ENABLED=0 go build -o /go/bin/loadbalancer cmd/main.go

FROM gcr.io/distroless/static-debian11
# path to the config file
ENV LOADBALANCER_CONFIG="/config.json"

COPY --from=build /go/bin/loadbalancer /server

EXPOSE 8080
EXPOSE 8081
CMD [ "/server" ]
