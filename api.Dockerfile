FROM golang:1.14 AS builder

WORKDIR /go/src/dht

COPY ./ ./

WORKDIR /go/src/dht/api

RUN go install -v

FROM ubuntu:latest

ARG PORT_ARG=5000

COPY --from=builder /go/bin /usr/local/bin

RUN echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc

WORKDIR /usr/local/bin

EXPOSE $PORT_ARG

CMD ["api"]