build:
	@go build -o bin/task-manager ./delivery/main.go

test:
	@test_packages=$$(github.com/abe16s/Go-Backend-Learning-path/task_manager/tests); \
	go test -v -cover $$test_packages

coverage:
	@test_packages=$$(github.com/abe16s/Go-Backend-Learning-path/task_manager/tests); \
	go test -v -coverprofile=coverage.out $$test_packages; \
	go tool cover -html=coverage.out

run: build
	@./bin/task-manager