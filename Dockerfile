FROM golang:latest

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD ./src /go/src/app