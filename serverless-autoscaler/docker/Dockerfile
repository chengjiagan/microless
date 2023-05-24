FROM golang:1.20 as build

# using go modules proxy in China
ARG GOPROXY="https://goproxy.cn"

WORKDIR /go/src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/sa main.go

FROM gcr.io/distroless/static-debian11
ENV SERVERLESS_AUTOSCALER_CONFIG="/sa.json"

COPY --from=build /go/bin/sa /sa

CMD [ "/sa" ]
