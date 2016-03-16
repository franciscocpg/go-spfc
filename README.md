# go-spfc
[![GoDoc](https://godoc.org/github.com/franciscocpg/go-spfc?status.svg)](https://godoc.org/github.com/franciscocpg/go-spfc) [![Build Status](https://travis-ci.org/franciscocpg/go-spfc.svg?branch=master)](https://travis-ci.org/franciscocpg/go-spfc)

A lib for wrapping services in Linux ([Upstart](http://upstart.ubuntu.com/) or [systemd](https://www.freedesktop.org/wiki/Software/systemd/)) and MacOS ([launchd](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man8/launchd.8.html)).

## Installing

1. With go get: `go get github.com/franciscocpg/go-spfc`
2. Or, the best way, using some dependency manager like [glide](https://github.com/Masterminds/glide)

## Using
Supposing you have a mongodb service installed in a CentOS (running systemd) or Ubuntu (running systemd or upstart) or MacOS (running launchd) machine you can instantiate a new handler this way:
```golang
 package main

import (
	"fmt"
	"github.com/franciscocpg/go-spfc"
)

func main() {
	h := service.NewHandler("mongod")
	st, err := h.Start()
	errorHandler(err)
	printStatus(st)
	st, err = h.GetStatus()
	errorHandler(err)
	printStatus(st)
	st, err = h.Stop()
	errorHandler(err)
	printStatus(st)
}

func printStatus(st service.Status) {
	fmt.Printf("Running: %t, st.PID: %d\n", st.Running, st.PID)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
```
**result**
```
Running: true, st.PID: 3765
Running: true, st.PID: 3765
Running: false, st.PID: 0
```

### Wait service to start or stop
Specially when you are using a Mac the previous code can result something like this.
**result**
```
Running: false, st.PID: 0
Running: false, st.PID: 0
Running: false, st.PID: 0
```
What happens? The problem is that [launchd](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man8/launchd.8.html) runs commands in a asynchronous way, so when we get the status, the service was not started (or stopped) yet.
To workaround this problem just wait the service start or stop this way
```
package main

import (
	"fmt"
	"github.com/franciscocpg/go-spfc"
	"time"
)

func main() {
	timeout := 10 * time.Second
	h := service.NewHandler("mongod")
	st, err := h.StartAndWait(timeout)
	errorHandler(err)
	printStatus(st)
	st, err = h.GetStatus()
	errorHandler(err)
	printStatus(st)
	st, err = h.StopAndWait(timeout)
	errorHandler(err)
	printStatus(st)
}

func printStatus(st service.Status) {
	fmt.Printf("Running: %t, st.PID: %d\n", st.Running, st.PID)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

```
If someting goes wrong, the *err* is populate with a timeout message.

## TODO
- Implement the [systemd](https://www.freedesktop.org/wiki/Software/systemd/) communication using DBUS ([go-systemd](https://github.com/coreos/go-systemd)) instead of a process call to **systemctl**.
