# nwatch

Simple replacement for `watch` command without all the bad unicode support bullshit.

## Installation

```bash
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/lnsp/nwatch
```

## Usage

```bash
$ nwatch -n 1s curl -s example.com/path
```