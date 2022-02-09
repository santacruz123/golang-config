.PHONY: init
init:
	go mod tidy -go=1.16

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: cli
cli:
	cd bin/cli; go run .

.PHONY: http
http:
	cd bin/http; go run .

