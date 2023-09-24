build:
	go build ./cmd/app/main.go
	./main.exe

cli:
	go build -o cli.exe ./cmd/cli/main.go
	./cli.exe

integration-tests:
	go test ./tests/integrational -c -o tests.exe
	./tests.exe

unit-tests:
	go test ./internal/domain/services -cover

run:
	docker-compose up -d

down:
	docker-compose down