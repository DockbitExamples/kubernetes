FROM alpine:latest
MAINTAINER The Dockbit Team "team@dockbit.com"

ARG track=stable
COPY source/$track/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
