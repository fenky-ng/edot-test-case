FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./bin/install_dependency.sh
RUN make clean init && \
    go build -buildvcs=false -o edot-shop-be ./app
EXPOSE 9200
ENTRYPOINT ["./edot-shop-be", "rest"]
