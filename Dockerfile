
FROM alpine:latest as certs

RUN apk --update add ca-certificates

COPY promql_exporter /promql_exporter

ENTRYPOINT ["/promql_exporter"]

EXPOSE 9312
