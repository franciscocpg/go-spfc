#!/bin/bash
set -e
if [[ $TRAVIS_OS_NAME == 'osx' ]]; then
	. ./.travis/mac-env
	cd "$PROJECT"
fi
if [[ "$DOCKER" = "true" ]]; then
	PROJECT="/root/go/src/github.com/franciscocpg/go-spfc"
	BRANCH=$(git rev-parse --abbrev-ref HEAD)
	pwd
	echo "BRANCH: $BRANCH"
	docker exec go-systemd-test bash -l -c 'cd '$PROJECT'; bash <(curl -s https://codecov.io/bash) -f $(cat coverage-filename) -t '$CODECOV_TOKEN' -b '$BRANCH''
else
	bash <(curl -s https://codecov.io/bash) -f $(cat coverage-filename)
fi
