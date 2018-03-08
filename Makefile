deps:
	go get -d -u ./....

build: vet
	go get ./...

test:
	go test -v

vet:
	go vet ./...
