FROM alpine:latest
MAINTAINER The Dockbit Team "team@dockbit.com"

ARG version=1.0
COPY source/$version/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
