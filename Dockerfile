FROM golang:latest

ENV LANG C.UTF-8

RUN apt-get update && \
  apt-get install -y strace && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv && \
  go get -u github.com/motemen/gore/cmd/gore && \
  go get -u github.com/mdempsky/gocode && \
  go get -u github.com/k0kubun/pp && \
  go get golang.org/x/tools/cmd/godoc

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

EXPOSE 3000:3000
