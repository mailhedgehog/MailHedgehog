#
# MailHedgehog Dockerfile
#
# docker build  -t yaroslawww/mailhedgehog:latest . --network host
# docker push  yaroslawww/mailhedgehog:latest

FROM golang:1.18-alpine as builder

# Install MailHedgehog:
RUN apk --no-cache add --virtual build-dependencies \
    git \
  && mkdir -p /root/gocode \
  && export GOPATH=/root/gocode \
  && go install github.com/mailhedgehog/MailHedgehog@v1.4.0

FROM alpine:latest
# Add mailhedgehog user/group with uid/gid 1000.
# This is a workaround for boot2docker issue #581, see
# https://github.com/boot2docker/boot2docker/issues/581
RUN adduser -D -u 1000 mailhedgehog

COPY --from=builder /root/gocode/bin/MailHedgehog /usr/local/bin/

USER mailhedgehog

WORKDIR /home/mailhedgehog

ENTRYPOINT ["MailHedgehog"]

# Expose the SMTP and HTTP ports:
EXPOSE 1025 8025
