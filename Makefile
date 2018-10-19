
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SHTEST=test
APP_NAME=ecr-get-token

all: test build
clean:
	$(GOCLEAN)
	rm -rf ./target
deps:
	xargs -t -a ./Godeps.txt -L 1 $(GOGET) -d
test:
	$(GOTEST) -v ./$(APP_NAME)/...
build:
	mkdir -p ./target
	$(GOBUILD) -v -o ./target/$(APP_NAME) ./$(APP_NAME)
run: build
	cd ./target; ./$(APP_NAME)
