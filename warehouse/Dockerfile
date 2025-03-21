FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./bin/install_dependency.sh
RUN make clean init && \
    go build -buildvcs=false -o edot-warehouse-be ./app
EXPOSE 9400
ENTRYPOINT ["./edot-warehouse-be", "rest"]
