# Ra Programming Language 𓋹

Ra is a simple interpreted programming language, built in Go. It features its own lexer, parser, and evaluator, and is designed for learning and experimenting with language development (not for production use).

```js
let hello = fn(name) {
    puts("Hello, " + name + "!");
}

hello("World");
```

## Getting Started

### Prerequisites
- Go 1.20+ installed

### Installation
Clone the repository:
```bash
git clone https://github.com/MikhailWahib/Ra.git
cd Ra
```

### Usage
Start the REPL:
```bash
go run ra.go
```

### Examples
```bash
go run ra.go examples/<example-name>.ra
```

### Build
```bash
go build -o bin/ra
```

### Testing

To run all tests:
```bash
go test ./...
```

To run a single test:
```bash
go test ./<package-dir-name>
```