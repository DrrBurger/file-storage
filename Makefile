.PHONY: build run clean test fmt vet lint help coverage-html coverage

## build: Билдит бинарный файл
build:
	go build -o bin/app -v cmd/main.go

## run_serv: Запускает сервер grpc
run:
	go run cmd/main.go

## clean: Очищяет и удаляет бинарный файл
clean:
	go clean
	rm -f bin/app

## test: Запускает все тесты
test:
	go test -v -race ./...

## cover-html: запускает тесты с получением отчёта в html формате
cover-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

 ## cover: Запускает тесты с покрытием
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

## fmt: Форматирование кода для соответствия стандартному стилю Go
fmt:
	go fmt ./...

## vet: Статический анализ кода на поиск подозрительных конструкций
vet:
	go vet ./...

## lint: Запускает линтер
lint:
	golangci-lint run

help: Makefile
	@echo " Choose a command run in "file-storage":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
