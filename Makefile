.PHONY: build

build:
	docker-compose -f ./build/docker-compose.yml up -d
test:
	python3 API-Requests-Tests.py
