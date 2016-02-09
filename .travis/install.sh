#!/bin/bash
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	export GOPATH="$HOME/go"
	brew unlink go
	brew update > /dev/null
    brew install go
    mkdir -p "$GOPATH/src/github.com"
    cp -R ../../franciscocpg "$GOPATH"
    cd "$GOPATH/src/github.com/franciscocpg/go-spfc"
    brew install glide
    go env
    echo $(pwd)
else
	# Install glide
    bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
fi
glide up