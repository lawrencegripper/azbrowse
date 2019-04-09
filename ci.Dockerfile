FROM golang:1.12-stretch

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install xorg openbox xserver-xorg-video-dummy xterm -y

COPY . /go/src/github.com/lawrencegripper/azbrowse
WORKDIR /go/src/github.com/lawrencegripper/azbrowse

ENTRYPOINT [ "/bin/bash" ]