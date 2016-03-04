#!/bin/bash
set -e
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	. ./.travis/mac-env
	brew unlink go
    echo "Updating brew..."
	brew update > /dev/null
    brew install go
    mkdir -p "$GITHUB"
    cp -R ../../franciscocpg "$GITHUB"
    cd "$PROJECT"
    brew install glide
    echo $(pwd)
else
    if [[ "$DOCKER" = "true" ]]; then
        PROJECT="/root/go/src/github.com/franciscocpg/go-spfc"
        docker build -t go-systemd .
        docker run -itd --name go-systemd-test -v /sys/fs/cgroup:/sys/fs/cgroup:ro -v /tmp/$(mktemp -d):/run go-systemd
        docker ps
        docker exec go-systemd-test systemctl --version
        docker exec go-systemd-test bash -l -c 'go get -u github.com/golang/lint/golint'
        docker exec go-systemd-test bash -l -c 'cd '$PROJECT'; glide up'
    else
        # Install glide
        bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
    fi
fi
if [[ -z "$DOCKER" ]]; then
    go get -u github.com/golang/lint/golint
    glide up
fi
