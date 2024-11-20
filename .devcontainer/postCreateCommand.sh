#!/bin/bash

echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list
sudo -E apt update && \
    apt -y install --no-install-recommends --allow-downgrades \
    fish \

go install -v github.com/davidrjenni/reftools/cmd/fillstruct@latest
go install -v github.com/cweill/gotests/gotests@latest
go install -v github.com/fatih/gomodifytags@latest
go install -v github.com/josharian/impl@latest
go install -v github.com/haya14busa/goplay/cmd/goplay@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install -v github.com/goreleaser/goreleaser/v2@latest
