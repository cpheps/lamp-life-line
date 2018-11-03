BINARY=lamp-life-line

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/cpheps/lamp-life-line/main.Version=${VERSION} -X github.com/cpheps/lamp-life-line/main.BuildTime=${BUILD_TIME}"

.PHONY: build clean fmt run vet imports test
default: build

build: | clean
	go build ${LDFLAGS} -o ./bin/${BINARY}

clean:
	if [ -f /bin/${BINARY} ] ; then rm bin/${BINARY} ; fi
	if [ -f /bin/application ] ; then rm bin/application ; fi

fmt: 
	@echo Formatting
	@goimports -w .
	@gofmt -s -w .

run:
	go run application.go

test:
#Add other package tests here
	go test ./...
	