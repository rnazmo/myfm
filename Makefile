.PHONY: mod
mod:
	go get -u ./... && go mod tidy

.PHONY: test
test:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: install-devel-tools
install-devel-tools:
# Install golangci-lint
# Ref: https://golangci-lint.run/usage/install/#linux-and-windows
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./devel-tools/bin/ v1.40.1
	./devel-tools/bin/golangci-lint --version

.PHONY: lint
lint:
# Ref: https://golangci-lint.run/usage/quick-start/
	./devel-tools/bin/golangci-lint -v run
