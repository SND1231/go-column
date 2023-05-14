.PHONY: containers-up
containers-up:
	docker-compose up -d

.PHONY: containers-down
containers-down:
	docker-compose down -d