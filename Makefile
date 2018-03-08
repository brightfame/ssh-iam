build: vet
	go get ./src

deps:
	go get -d -u ./src

test:
	go test -v src

vet:
	go vet ./src

image:
	docker build -t brightfame/ssh-iam .
