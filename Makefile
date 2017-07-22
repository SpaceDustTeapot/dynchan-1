MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)

default: build

get-deps:
	go get ./...

test:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

cover: test
	go tool cover -html=coverage-all.out

run:
	go run main.go

build: clean
	go build -a -o dynchan .

clean:
	rm -rf dynchan coverage.out coverage-all.out
