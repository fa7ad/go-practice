FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go build

CMD ["/go/src/app/go-practice"]
