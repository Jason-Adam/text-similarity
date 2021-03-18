.PHONY: build
build:
	go build \
		-o bin/similarityapi \
		cmd/api/main.go

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: test
test:
	sh bin/test.sh

.PHONY: build-docker
build-docker:
	sh bin/build-docker.sh

.PHONY: run-docker
run-docker: build-docker
	sh bin/run-docker.sh
