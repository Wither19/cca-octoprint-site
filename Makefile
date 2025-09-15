.PHONY: go vite build preview

go: 
	go build

vite:
	cd frontend && bun run build

build: go vite
	./cca-octoprint

preview:
	./cca-octoprint