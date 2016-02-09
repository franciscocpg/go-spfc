#!/bin/bash
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	brew unlink go
    brew install go
    brew install glide
    go env
else
	# Install glide
    bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
fi
glide up