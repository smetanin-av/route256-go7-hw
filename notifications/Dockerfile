FROM ubuntu:22.04

RUN mkdir -p /usr/local/share/ca-certificates/extra
COPY telegram.crt /usr/local/share/ca-certificates/extra
RUN apt-get update && apt-get install ca-certificates -y && update-ca-certificates

ADD ./bin/app /app
ADD ./config.yaml /config.yaml

CMD ["/app"]
