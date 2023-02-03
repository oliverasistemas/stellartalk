# vim: set ft=make ts=8 noet

.PHONY: help
help:	### Get some help about make targets.
### this screen. Keep it first target to be default
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif



.PHONY: init
init:  ### Runs initialization script
	./init.sh

.PHONY: docker-build
docker-build: ### Build docker image
	docker build -t stellartalk:latest .

.PHONY: install
install:  ### Install stellartalk binaries in your $GOPATH/bin folder.
	go install ./cmd/stellartalkd

start:  ### start the chain
	stellartalkd start

.PHONY: docker-run
docker-run: docker-build ### Run docker container with tag stellartalk:latest
	docker run stellartalk:latest

.PHONY: lint
lint: ### Run linter.
	golangci-lint run

.PHONY: test
test: ### Run tests.
	go test -v --cover "./..."
