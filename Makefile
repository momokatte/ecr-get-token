
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
	for DEP in $(shell cat ./Godeps.txt); do $(GOGET) $$DEP; done
test:
	$(GOTEST) -v ./$(APP_NAME)/...
build:
	mkdir -p ./target
	$(GOBUILD) -v -o ./target/$(APP_NAME) ./$(APP_NAME)
run: build
	cd ./target; ./$(APP_NAME)
