# Emo

Emoji based semantic scoped debuging for Go, Python, Typescript/Javascript and Dart

## Documentation

- [Go](doc/go/README.md) library
- [Python](lang/python/README.md) library
- [Typescript](doc/typescript/README.md) library
- [Dart](lang/dart/README.md) library

Complete [events list](doc/events/README.md)

## Development

### Run the tests

```bash
make test
```

Visualize the tests coverage:

```bash
make cov
```

### Generate the code

Run the codegen tools to build up the functions from the `codegen/ref.json` file. 

#### Build the Go library

```bash
make gocodegen
```

This will regenerate the `emo_gen.go` file

Build the Python library:

```bash
make pycodegen
```

This will regenerate the `lang/python/pyemo/emo_gen.py` file

#### Build the Typescript library:

```bash
make tscodegen
```

This will regenerate the `lang/typescript/src/emo_gen.ts` file

#### Build the Dart library:

```bash
make dartcodegen
```
This will regenerate the `lang/dart/lib/src/debug.dart` file

#### Build all languages at once

```bash
make codegen
```

### Generate the doc

Regenerate the complete events list 

```bash
make docgen
```