FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD oam-controller /bin/oam-controller
ADD config.yml /etc/oam-controller/config.yml

CMD ["oam-controller", "-config", "/etc/oam-controller/config.yml"]