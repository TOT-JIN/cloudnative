FROM golang:1.17.2-alpine3.14 as builder

ADD main.go /opt
WORKDIR /opt
RUN go build main.go

FROM alpine:3.14
MAINTAINER zhanjie.jin i@zhanjie.me

RUN apk --no-cache add ca-certificates
COPY --from=builder /opt/main /opt
WORKDIR /opt
VOLUME /opt
EXPOSE 80
CMD ["/opt/main"]
