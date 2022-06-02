.PHONY: start
start:
	docker compose up -d --build

.PHONY: stop
stop:
	docker compose rm -v --force --stop
	docker image rm golang_api:latest

.PHONY: test
test:
	sh ./scripts/e2e-testing.sh

.PHONY: dev
dev:
	go run cmd/golang_api/main.go