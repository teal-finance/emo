# Emo

Emoji based semantic scoped debugging for Go, Python, Typescript/Javascript and Dart

## Documentation

- [Go](doc/go/) library
- [Python](lang/python/) library
- [Typescript](doc/typescript/README.md) library
- [Dart](lang/dart/) library

Complete [events list](doc/events/README.md)

## How to contribute

To add more emojis and methods please edit the [`codegen/ref.json`](codegen/ref.json) file.
The code in all languages is generated in from this file.

### Generate the code

Run the codegen tools to build up the functions from the [`codegen/ref.json`](codegen/ref.json) file.

#### Build the Go library

```bash
make go
```

This will regenerate the `generated.go` file

#### Build the Python library

```bash
make py
```

This will regenerate the `lang/python/pyemo/emo_gen.py` file

#### Build the Typescript library

```bash
make ts
```

This will regenerate the `lang/typescript/src/emo_gen.ts` file

#### Build the Dart library

```bash
make dart
```

This will regenerate the `lang/dart/lib/src/debug.dart` file

#### Build all languages at once

```bash
make all
```

### Generate the doc

Regenerate the complete events list:

```bash
make doc
```

## Development

### Run the tests

```bash
make test
```

Visualize the tests coverage:

```bash
make cov
```
