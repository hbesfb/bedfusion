FROM bash:5.2.21@sha256:a422913be6a4b0ea5403d2af72eb73779b6d9ed84d0dcf85d6b4309e891a379e AS BASH
FROM scratch

# The following will add the minimum set of libraries to run bash
COPY --from=BASH /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1
COPY --from=BASH /lib/libc.musl-x86_64.so.1 /lib/libc.musl-x86_64.so.1
COPY --from=BASH /usr/lib/libncursesw.so.6.4  /usr/lib/libncursesw.so.6.4
COPY --from=BASH /usr/lib/libncursesw.so.6 /usr/lib/libncursesw.so.6

# Add bash to the expected location for nextflow
COPY --from=BASH /usr/local/bin/bash /bin/bash

COPY bedfusion /bedfusion
ENTRYPOINT ["/bedfusion"]
