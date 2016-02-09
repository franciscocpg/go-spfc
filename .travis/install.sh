#!/bin/bash
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
    brew install go
    brew install glide
else
	# Install glide
    bash <(curl -s https://gist.githubusercontent.com/franciscocpg/ab10b57898978009638f/raw/)
fi
glide up