##
#-

.PHONY: all
all: ensure-deps

.PHONY: ensure-deps
ensure-deps:
	@echo "Syncing dependencies with go mod tidy"
	@go mod tidy