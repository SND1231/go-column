.PHONY: build
build:
	go build .

.PHONY: sqlboiler
sqlboiler:
	sqlboiler mysql

.PHONY: test
test:
	# 毎回テスト実行時にキャッシュを削除
	go clean -testcache
	# shuffle=onでランダムに実行され、coverでテストのカバレッジを測定している。
	go test -shuffle=on -cover ./usecase

.PHONY: checkvuln
checkvuln:
	govulncheck ./...