.PHONY: build preview dev

build:
	@echo "Building Go source"
	go build
	cd frontend && bun run build

preview: build
	./cca-octoprint

dev:
	./cca-octoprint