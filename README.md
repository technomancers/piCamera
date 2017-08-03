# piCamera [![GoDoc](https://godoc.org/github.com/technomancers/piCamera?status.svg)](https://godoc.org/github.com/technomancers/piCamera)

This package is a wrapper for the `raspivid` command on the Raspberry Pi. To make development easier on a PC, there is are different `Start()` methods depending on what is compiled. One is for the Raspberry Pi and the other is for everything else.

## Installation

Since this package depends on a Raspberry Pi only command there is an extra flag needed to build this package for the Raspberry Pi.

```sh
env GOOS=linux GOARCH=arm GOARM=7 go build -tags pi -a .
```

Make note of the `-tags pi` on build. Any `main` package that has any dependency to this package should have that flag so that this package is built correctly for the Raspberry Pi.
