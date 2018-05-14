FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD oamd /bin/oamd
ADD config.yml /etc/oamd/config.yml

CMD ["oamd", "-config", "/etc/oamd/config.yml"]