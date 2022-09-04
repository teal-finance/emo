help:
	# Use 'make <target>' where <target> is one of:
	#
	# all   Generate all
	# go    Generate code for the Go library
	# ts    Generate code for the Typescript library
	# py    Generate code for the Python library
	# dart  Generate code for the Dart library
	# doc   Generate the documentation
	#
	# test  Test the Go library
	# cov   Test and visualize the coverage
	# vet   Upgrade deps, Format code, Test and Run example

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

.PHONY: test
test:
	go test -race -vet all -tags=emo -coverprofile=c.out ./...

.PHONY: cov
cov: test
	go tool cover -html c.out

.PHONY: vet
vet:
	go mod tidy
	go get -u -t all
	go mod tidy
	go generate ./...
	go run mvdan.cc/gofumpt@latest -w -extra -l -lang 1.19 .
	go build ./...
	go test -race -vet all -tags=emo -coverprofile=c.out ./...
	go run ./examples/go/example.go
