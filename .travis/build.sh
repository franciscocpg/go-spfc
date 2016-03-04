#!/bin/bash
set -e
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	. ./.travis/mac-env
	cd "$PROJECT"
fi
if [[ "$DOCKER" = "true" ]]; then
	PROJECT="/root/go/src/github.com/franciscocpg/go-spfc"
	docker exec go-systemd-test bash -l -c 'cd '$PROJECT'; go vet'
	docker exec go-systemd-test bash -l -c 'cd '$PROJECT'; golint'
	docker exec go-systemd-test bash -l -c 'cd '$PROJECT'; bin/build-test'
else
	go vet
	golint
	bin/build-test	
fi
