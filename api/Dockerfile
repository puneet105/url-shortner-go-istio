FROM golang:alpine as builder

RUN mkdir -p url-shortner-go/api

ADD . url-shortner-go/api

COPY .env url-shortner-go/api

RUN mv url-shortner-go /usr/local/go/src/ 

WORKDIR /usr/local/go/src/url-shortner-go/api

#RUN go build main.go

RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
#FROM alpine:latest

#COPY --from=builder /usr/local/go/src/url-shortner-go/api/main url-shortner-go/api

#COPY --from=builder /usr/local/go/src/url-shortner-go/api/.env url-shortner-go

#WORKDIR url-shortner-go

EXPOSE 3001

CMD ["./api"]
