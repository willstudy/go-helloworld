FROM golang:1.9.2 

MAINTAINER Liu Chang(studywiller@gmail.com)

RUN mkdir /home/work
ADD . /home/work
WORKDIR /home/work
RUN go build -o hello-world main.go
EXPOSE 9000

ENTRYPOINT sh hello-world 
