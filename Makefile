
.PHONY: help
help: ### Prints this help
	@awk 'BEGIN {FS = ":.*?### "} /^[a-zA-Z0-9_-]+:.*?### / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "\033[32;1m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

VERSION = $(shell cat VERSION 2>/dev/null || echo 'v0.0.0')
.PHONY: version
version: ### Shows the version
	@echo $(VERSION)

.PHONY: hello
hello: ### Says hello
	@echo Hello

.PHONY: test
test: ### Runs the tests
	go test -cover -coverprofile coverage.out
	go tool cover -func coverage.out
	go tool cover -html coverage.out -o coverage.html
