#!/bin/bash
set -e
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	. ./.travis/mac-env
	brew unlink go
	brew update > /dev/null
    brew install go
    mkdir -p "$GITHUB"
    cp -R ../../franciscocpg "$GITHUB"
    cd "$PROJECT"
    brew install glide
    echo $(pwd)
else
	# Install glide
    bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
fi
go get -u github.com/golang/lint/golint
glide up