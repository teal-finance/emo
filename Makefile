help:
	# make all    Upgrade, Generate, Format, Go test/lint
	#
	# make go     Generate code for the Go library
	# make ts     Generate code for the Typescript library
	# make py     Generate code for the Python library
	# make dart   Generate code for the Dart library
	# make doc    Generate the documentation
	#
	# make up     Go: Upgrade the patch version of the dependencies
	# make up+    Go: Upgrade the minor version of the dependencies
	# make fmt    Go: Generate code and Format code
	# make test   Go: Check build and Test
	# make cov    Go: Browse test coverage
	# make vet    Go: Run example and Lint

.PHONY: all
all: up fmt test vet

.PHONY: go
go:
	go run codegen/main.go -go

.PHONY: ts
ts:
	go run codegen/main.go -ts

.PHONY: py
py:
	go run codegen/main.go -py

.PHONY: dart
dart:
	go run codegen/main.go -dart

.PHONY: doc
doc:
	go run codegen/main.go -doc

go.sum: go.mod
	go mod tidy

.PHONY: up
up: go.sum
	GOPROXY=direct go get -t -u=patch all
	go mod tidy

.PHONY: up+
up+: go.sum
	go get -u -t all
	go mod tidy

.PHONY: fmt
fmt:
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l -lang 1.19 .

.PHONY: test
test:
	go build ./...
	go test -race -vet all -tags=emo -coverprofile=code-coverage.out ./...

code-coverage.out: go.sum */*.go
	go test -race -vet all -tags=emo -coverprofile=code-coverage.out ./...

.PHONY: cov
cov: code-coverage.out
	go tool cover -html code-coverage.out

.PHONY: vet
vet:
	go run -race ./examples/go
	go run -race github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix
