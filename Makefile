.PHONY: build preview


build:
	cd frontend && bun run build
	cd ..
	go build
	./cca-octoprint

preview:
	./cca-octoprint