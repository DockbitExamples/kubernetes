FROM alpine:latest
MAINTAINER The Dockbit Team "team@dockbit.com"

ARG version=1.0
COPY source/1.0/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
