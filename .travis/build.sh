#!/bin/bash
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	. ./.travis/mac-env
	cd "$PROJECT"
fi
go vet
golint
bin/build-test