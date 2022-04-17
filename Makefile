help:
	# Use 'make <target>' where <target> is one of:
	#
	# codegen          Generate code for all language
	# gocodegen        Generate code for the Go library
	# tscodegen        Generate code for the Typescript library
	# pycodegen        Generate code for the Python library
	#
	# docgen           Generate the documentation
	#
	# test             Run the tests on the Go library
	# test             Run the tests on the Go library and visualize the coverage

.PHONY: codegen
codegen:
	go run codegen/main.go

.PHONY: gocodegen
gocodegen:
	go run codegen/main.go -go

.PHONY: tscodegen
tscodegen:
	go run codegen/main.go -ts

.PHONY: pycodegen
pycodegen:
	go run codegen/main.go -py

.PHONY: docgen
docgen:
	go run codegen/main.go -doc

.PHONY: test
test:
	go test -tags=emo -coverprofile=c.out

.PHONY: cov
cov: test
	go tool cover -html c.out

