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
        docker run -itd --name systemd -v /sys/fs/cgroup:/sys/fs/cgroup:ro -v /tmp/$(mktemp -d):/run milcom/centos7-systemd
        docker ps -a
    else
        # Install glide
        bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
    fi
fi
go get -u github.com/golang/lint/golint
glide up