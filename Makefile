.PHONY: node build preview

node: 
	cd frontend && npm i

build:
	cd frontend && bun run build
	go build
	./cca-octoprint

preview:
	./cca-octoprint