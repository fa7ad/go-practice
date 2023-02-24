FROM golang:1.20-alpine AS builder

RUN apk add --no-cache --update curl

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

ARG PORT=3000
ENV PORT $PORT

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/backend api/app.go

RUN touch $$DB_PATH

CMD ["/code/bin/backend"]

FROM builder as dev-env

RUN sh -c "apk add git; \
go install github.com/cosmtrek/air@latest"

CMD ["air"]

FROM builder as prod-env

RUN apk add --no-cache tini

COPY /code/bin/backend /usr/local/bin/backend

CMD ["/usr/local/bin/backend"]
ENTRYPOINT ["/sbin/tini", "-"]