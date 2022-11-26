.PHONY: run
run: ## run app
	go run cmd/othello-go/main.go

.PHONY: test
run: ## run unit test
	go test ./...


