FROM golang:1.11 as build

COPY ./ /go/src/github.com/innovate-technologies/yp-server
WORKDIR /go/src/github.com/innovate-technologies/yp-server

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/innovate-technologies/yp-server/yp-server /opt/yp-server/yp-server

RUN chmod +x /opt/yp-server/yp-server

WORKDIR /opt/yp-server/

CMD ["/opt/yp-server/yp-server"]