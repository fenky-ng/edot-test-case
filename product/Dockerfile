FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./bin/install_dependency.sh
RUN make clean init && \
    go build -buildvcs=false -o edot-product-be ./app
EXPOSE 9300
ENTRYPOINT ["./edot-product-be", "rest"]
