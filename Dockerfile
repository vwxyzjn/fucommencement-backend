FROM golang:1.9.4-alpine3.7

WORKDIR /go/src/gitlab.com/vwxyzjn/fucommencement-backend
COPY . /go/src/gitlab.com/vwxyzjn/fucommencement-backend
ENV APP_ENVIRONMENT=development

RUN apk add --no-cache git
RUN wget https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
RUN mv dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN dep ensure

EXPOSE 8080 80 443