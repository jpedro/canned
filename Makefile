
.PHONY: help
help: ### Prints this help
	@awk 'BEGIN {FS = ":.*?### "} /^[a-zA-Z0-9_-]+:.*?### / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "\033[33;1m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

VERSION = $(shell cat VERSION 2>/dev/null || echo 'v0.1.0')
.PHONY: version
version: ### Shows the version
	@echo $(VERSION)

.PHONY: hello
hello: ### Says hello
	@echo Hello

