FROM debian:bookworm
COPY bedfusion /bedfusion
ENTRYPOINT ["/bedfusion"]
