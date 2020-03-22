FROM golang:1.13.5-alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apk add --no-cache \
      git \
      gcc \
      libc-dev \
      make

RUN go get -u github.com/oxequa/realize
