GOPATH=${HOME}/Go

create-base:
	mkdir -p ${GOPATH}

create-structure: create-base
	mkdir -p ${GOPATH}/src/github.com/andriolisp/word-count

copy-structure: create-structure
	cp -R ./* ${GOPATH}/src/github.com/andriolisp/word-count/
	cd ${GOPATH}/src/github.com/andriolisp/word-count/

get-dependencies: create-structure
	go get

test-code: get-dependencies
	go test -v

install: test-code
	go install

build: test-code
	go build -o word-count
