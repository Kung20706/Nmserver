FROM golang:1.14-alpine AS builder



ENV GOBIN=/go/bin

WORKDIR /go/src/BearApp

COPY ./config /BearApp/config

COPY . /go/src/BearApp
COPY ./internal /BearApp/internal
RUN ls
COPY ./main.go /BearApp

# COPY ./go.mod /BearApp

# COPY ./go.sum /BearApp

# COPY ./docs /BearApp/docs
RUN go mod init BearApp

RUN go install
# RUN go get "github.com/liuzl/gocc"

RUN go get github.com/swaggo/swag/cmd/swag 

RUN swag init

RUN sed "s/^func init/func Init/" docs/docs.go > docs/document.go 

RUN rm -rf  docs/docs.go

RUN go build -o app

ENTRYPOINT PROJECT_ENV=docker ./app server
