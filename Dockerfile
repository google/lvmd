FROM alpine:latest

RUN set -ex; \
  apk add --no-cache --no-progress --virtual .build-deps git gcc musl-dev openssh bash go; \
  env GOPATH=/go go get -v github.com/google/lvmd; \
  install -t /bin /go/bin/lvmd; \
  rm -rf /go; \
  apk --no-progress del .build-deps
