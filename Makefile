.PHONY: build

build:
	docker-compose -f ./build/docker-compose.yml up -d
