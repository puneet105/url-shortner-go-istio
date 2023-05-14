FROM golang:1.19 AS builder

RUN cd /usr/local/go/src && mkdir -p github.com/puneet105/url-shortner-go/api

WORKDIR /usr/local/go/src/github.com/puneet105/url-shortner-go/api

ADD ./api ./

RUN go mod tidy && GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go


FROM alpine:latest

COPY --from=builder /usr/local/go/src/github.com/puneet105/url-shortner-go/api/main ./

EXPOSE 3001

