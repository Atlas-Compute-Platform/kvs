# Atlas Dockerfile
FROM alpine:latest
RUN apk add go
COPY ./kvs /usr/local/bin/kvs
EXPOSE 8800/tcp
CMD ["/usr/local/bin/kvs"]
