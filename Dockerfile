FROM golang:1.23.3-bookworm AS builder

WORKDIR /bedfusion
ENV CGO_ENABLED=0
COPY . .
RUN go build ./...

FROM debian:bookworm
COPY --from=builder /bedfusion/bedfusion /usr/local/bin/
ENTRYPOINT ["bedfusion"]
