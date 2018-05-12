.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o hellok8s .

serve: build
	./hellok8s

clean:
	rm ./hellok8s

pack:
	GOOS=linux make build
	docker build -t tschokko/hellok8s:$(TAG) .

upload:
	docker push tschokko/hellok8s:$(TAG)

deploy:
	envsubst < deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean