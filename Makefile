.PHONY: containers-up
containers-up:
	docker-compose up -d

.PHONY: container-build
container-build:
	docker-compose build --no-cache

.PHONY: containers-down
containers-down:
	docker-compose down