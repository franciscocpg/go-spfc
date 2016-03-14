# go-spfc
[![GoDoc](https://godoc.org/github.com/franciscocpg/go-spfc?status.svg)](https://godoc.org/github.com/franciscocpg/go-spfc) [![Build Status](https://travis-ci.org/franciscocpg/go-spfc.svg?branch=master)](https://travis-ci.org/franciscocpg/go-spfc)

A lib for wrapping services in Linux ([Upstart](http://upstart.ubuntu.com/) or [systemd](https://www.freedesktop.org/wiki/Software/systemd/)) and MacOS ([launchd](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man8/launchd.8.html)).

## Installing

1. With go get: `go get github.com/franciscocpg/go-spfc`
2. Or, the best way, using some dependency manager like [glide](https://github.com/Masterminds/glide)

## Using
Supposing you have a mongodb service
 