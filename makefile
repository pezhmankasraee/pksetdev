VERSION := $(shell git describe --tags --dirty --always)
build:
	go build -ldflags "-X 'github.com/pezhmankasraee/pksetdev/help.Version=$(VERSION)'"

