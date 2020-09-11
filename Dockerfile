FROM golang:latest

ENV LANG C.UTF-8

RUN apt-get update && \
  apt-get install -y strace && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

EXPOSE 3000
