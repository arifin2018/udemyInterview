FROM golang:1.22-alpine

# RUN apt-get -y install git
RUN apk add --no-cache git

RUN mkdir /udemy
WORKDIR /udemy

COPY . /udemy