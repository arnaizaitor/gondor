# Makefile for Gondor - Go n-dimensional array reimplementation

APP_NAME = gondor

.PHONY: all build run test coverage fmt lint vet clean

## Compila el proyecto
build:
	go build -o bin/$(APP_NAME) ./...

## Ejecuta el proyecto principal (desde ./main.go o /cmd)
run:
	go run ./cmd/main.go

## Ejecuta todos los tests
test:
	go test ./...

## Test con cobertura
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📄 Ver coverage en coverage.html"

clean-coverage:
	rm -f coverage.out coverage.html coverage

## Formatea el código
fmt:
	go fmt ./...

## Linter (requiere instalar golangci-lint)
lint:
	golangci-lint run

## Análisis estático
vet:
	go vet ./...

## Limpia binarios y archivos temporales
clean:
	go clean
	rm -f coverage.out coverage.html
	rm -rf bin/

## Instala dependencias necesarias para análisis y linting
setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
