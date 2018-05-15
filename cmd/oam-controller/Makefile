.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
APPNAME=oam-controller
EXECUTABLE=oam-controller

export TAG

install:
	go get .

test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o $(EXECUTABLE) .

serve: build
	./$(EXECUTABLE)

clean:
	rm ./$(EXECUTABLE)

pack:
	GOOS=linux make build
	docker build -t tschokko/$(APPNAME):$(TAG) .

upload:
	docker push tschokko/$(APPNAME):$(TAG)

deploy:
	envsubst < deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean