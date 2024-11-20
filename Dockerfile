FROM golang:1:13.3-bookworm AS builder
COPY . /bedfusion
WORKDIR /bedfusion
RUN go build .

FROM debian:bookworm
COPY --from=builder /bedfusion/bedfusion /usr/local/bin/
