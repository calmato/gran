FROM alpine:3.11.3

ENV TERRAFORM_VERSION 0.12.21

RUN apk add --no-cache \
      wget \
      unzip

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip

RUN mv terraform /usr/local/bin
