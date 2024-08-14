FROM golang:1.22-alpine

RUN mkdir /udemy
WORKDIR /udemy

COPY . /udemy