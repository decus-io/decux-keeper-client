FROM centos:7
WORKDIR /app
COPY decux-keeper-client .
COPY config.yaml .
ENTRYPOINT ["/app/decux-keeper-client"]
