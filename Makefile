help:
	# make all    Upgrade deps, Generate all, Format code, Test Go code, Lint Go code
	#
	# make go     Generate code for the Go library
	# make ts     Generate code for the Typescript library
	# make py     Generate code for the Python library
	# make dart   Generate code for the Dart library
	# make doc    Generate the documentation
	#
	# make up     Go: Upgrade deps
	# make fmt    Go: Generate code and Format code
	# make test   Go: Check build and Test
	# make cov    Go: Test and Visualize the code coverage
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

.PHONY: up
up:
	go mod tidy
	go get -u -t all

.PHONY: fmt
fmt:
	go mod tidy
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l -lang 1.19 .

.PHONY: test
test:
	go build ./...
	go test -race -vet all -tags=emo -coverprofile=code-coverage.out ./...

.PHONY: cov
cov: test
	go tool cover -html code-coverage.out

.PHONY: vet
vet:
	go run ./examples/go
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix
