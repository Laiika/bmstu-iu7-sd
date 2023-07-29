build:
	go build ./cmd/main.go
	./main.exe

integration-tests:
	go test ./tests/integrational -c -o tests.exe
	./tests.exe

unit-tests:
	go test ./internal/domain/services -cover

run:
	docker-compose up -d

down:
	docker-compose down