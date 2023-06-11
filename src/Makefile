.PHONY: build
build:
	go build .

.PHONY: sqlboiler
sqlboiler:
	sqlboiler mysql

.PHONY: test
test:
	go clean -testcache
	go test -shuffle=on -cover ./usecase