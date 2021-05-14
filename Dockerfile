FROM centos:7
WORKDIR /app
COPY decus-keeper-client .
COPY config.yaml .
ENTRYPOINT ["/app/decus-keeper-client"]
