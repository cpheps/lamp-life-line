BINARY=lamp-life-line

.PHONY: build clean fmt run vet imports test
default: build

build: | clean
	go build -o ./bin/${BINARY}

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
	