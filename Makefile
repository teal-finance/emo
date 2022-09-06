help:
	# make all    Generate all
	# make go     Generate code for the Go library
	# make ts     Generate code for the Typescript library
	# make py     Generate code for the Python library
	# make dart   Generate code for the Dart library
	# make doc    Generate the documentation
	#
	# make up     Upgrade deps, Generate code, Format code, Check build
	# make test   Test the Go library
	# make cov    Test and visualize the code coverage
	# make vet    Upgrade, Generate, Format, Build, Test, Run example and Lint

.PHONY: all
all:
	go run codegen/main.go

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

.PHONY: up
up:
	go mod tidy
	go get -u -t all
	go mod tidy
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l -lang 1.19 .
	go build ./...

.PHONY: test
test:
	go test -race -vet all -tags=emo -coverprofile=code-coverage-of-tests.out ./...

.PHONY:
cov: test
	go tool cover -html code-coverage-of-tests.out

.PHONY: vet
vet: up test
	go run ./examples/go/example.go
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix
