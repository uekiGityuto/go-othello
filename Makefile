.PHONY: run
run: ## run app
	go run ./cmd/othello-go

.PHONY: test
test: ## run unit test
	go test ./...

.PHONY: build
build: ## build app
	mkdir -p ./build
	go build -o ./build ./cmd/othello-go

