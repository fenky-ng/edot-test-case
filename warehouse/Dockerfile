FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./bin/install_dependency.sh
RUN make clean init && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o edot-warehouse-be ./app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/edot-warehouse-be .
RUN chmod +x edot-warehouse-be
