.PHONY: build preview

build:
	@echo "Building Go source"
	go build

preview: build
	./jv-pokeapi