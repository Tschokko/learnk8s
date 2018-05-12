FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD hellok8s /bin/hellok8s
#ADD config.yml.dist /etc/news/config.yml

CMD ["hellok8s"] 
#, "-config", "/etc/news/config.yml"