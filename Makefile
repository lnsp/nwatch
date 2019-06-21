all: build

clean:
	rm nwatch

build:
	go build

install: build
	cp nwatch /usr/local/bin/nwatch