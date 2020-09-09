FROM golang:alpine

RUN apk add --update git \
	&& addgroup -g 1000 -S go \
	&& adduser -h /home/go -g "" -s /bin/sh -G go -S -D -u 1000 go