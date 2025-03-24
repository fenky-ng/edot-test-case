#!/usr/bin/env bash

GO_INSTALL_OPTS=""
install_go_dependencies() {

    echo "Installing dependencies..."
    go install $GO_INSTALL_OPTS go.uber.org/mock/mockgen@v0.3.0
    go install $GO_INSTALL_OPTS github.com/google/wire/cmd/wire@v0.6.0
}

check_version() {
    echo "Checking version..."
    go version
    mockgen --version
}

install_go_dependencies
check_version
