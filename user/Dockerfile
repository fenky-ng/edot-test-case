FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./bin/install_dependency.sh
RUN make clean init && \
    go build -buildvcs=false -o edot-user-be ./app
EXPOSE 9100
ENTRYPOINT ["./edot-user-be", "rest"]
