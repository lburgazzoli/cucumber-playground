
test:
	go test -test.v ./...

godog:
	go test -test.v --godog.format=pretty ./...

PHONY: .test
PHONY: .godog