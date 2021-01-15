all: build

build:
	@echo "Запущена сборка проекта"
	@go build -o bin/api main.go
	@echo "Проект собран и находится в ./bin/server"

run: build
	@echo "Запуск сервера"
	@cd bin && ./api

test: build
	@echo "Запуск сервера в тестовом режиме без логирования в файл"
	@cd bin && ./api -nolog