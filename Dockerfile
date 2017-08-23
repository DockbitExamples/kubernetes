FROM alpine:latest
MAINTAINER The Dockbit Team "team@dockbit.com"

COPY app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
