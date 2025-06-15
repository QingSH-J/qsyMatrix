# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go project implementing a generic matrix library using Go 1.18+ generics. The project is in early development stages with a focus on type-safe matrix operations.

**Module:** `github.com/QingSH-J/qsyMatrix`
**Go Version:** 1.24.2

## Development Commands

### Basic Operations
```bash
# Run the application
go run main.go

# Build the project
go build

# Format code
go fmt ./...

# Vet code for common issues
go vet ./...

# Manage dependencies
go mod tidy
```

### Testing
```bash
# Run all tests (once tests are added)
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./matrix
```

## Code Architecture

### Package Structure
- **main package** (`main.go`): Application entry point
- **matrix package** (`matrix/matrix.go`): Core matrix implementation using generics

### Key Design Patterns
- **Generic Programming**: Uses Go 1.18+ generics with type constraints
- **Number Interface**: Supports multiple numeric types (int, float64, complex128, etc.) plus strings and bytes
- **Type Safety**: Leverages Go's type system for compile-time safety

### Matrix Package
The matrix package defines:
- `Number` interface with union type constraints
- `Matrix[T Number]` generic struct
- Constructor function `New[T Number]` for matrix creation

## Development Notes

- The project uses modern Go features including generics and type constraints
- No external dependencies currently
- Code follows standard Go conventions and formatting
- Project is in early development stage with basic structure in place