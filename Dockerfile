FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/YDshop

ENV key=value

RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get github.com/labstack/echo

CMD go run main.go
