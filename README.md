# zerr

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/zerr/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/zerr/actions/workflows/test.yml)
[![Lint Status](https://github.com/alex-cos/zerr/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/zerr/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/zerr)](https://goreportcard.com/report/github.com/alex-cos/zerr)

`zerr` is a lightweight Go package that provides an enhanced error structure with **severity levels**, **error codes**, and **error wrapping** support — fully compatible with the Go 1.13+ `errors` package.

## Features

- **Severity levels** — `Debug`, `Info`, `Notice`, `Warning`, `Error`, `Critical`, `Fatal`
- **Error codes** — attach a numeric code to any error for machine-readable handling
- **Error wrapping** — wrap existing errors while preserving severity and code
- **Standard library compatible** — works with `errors.Is`, `errors.As`, and `fmt.Errorf`
- **Zero dependencies**

## Installation

```bash
go get github.com/alex-cos/zerr
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/alex-cos/zerr"
)

func main() {
    err := zerr.New("something went wrong")
    fmt.Println(err)
    // Output: Error - something went wrong
}
```

## Creating Errors

### Basic errors

```go
err := zerr.New("not found")
// Error - not found
```

### With severity

```go
err := zerr.NewS(zerr.Warning, "disk space low")
err := zerr.NewS(zerr.Critical, "database unreachable")
// Warning - disk space low
// Critical - database unreachable
```

### With error code

```go
err := zerr.NewC(404, "resource not found")
// Error[404] - resource not found
```

### With severity and code

```go
err := zerr.NewSC(zerr.Fatal, 500, "system shutdown")
// Fatal[500] - system shutdown
```

### Formatted errors

```go
err := zerr.Errorf("failed to open file '%s': %v", path, reason)
err := zerr.ErrorSf(zerr.Warning, "retry %d of %d", attempt, max)
err := zerr.ErrorCf(1001, "timeout after %ds", seconds)
err := zerr.ErrorSCf(zerr.Critical, 2001, "connection lost to %s", host)
```

## Wrapping Errors

Wrap an existing error while adding context, severity, and/or a code:

```go
origErr := someFunction()

err := zerr.Wrap(origErr, "failed to process request")
err := zerr.WrapS(zerr.Critical, origErr, "database connection lost")
err := zerr.WrapC(5001, origErr, "upstream timeout")
err := zerr.WrapSC(zerr.Fatal, 9001, origErr, "system unrecoverable")
```

## Inspecting Errors

### Using helper functions (recommended)

```go
zerr.IsSeverity(err, zerr.Warning) // true/false
code, ok := zerr.GetCode(err)      // 404, true
msg := zerr.GetMessage(err)        // "lookup failed"
```

### Using type assertion

```go
err := zerr.WrapC(404, zerr.New("user not found"), "lookup failed")

zerr := err.(*zerr.ZError)

zerr.Severity()  // Error
zerr.Code()      // 404
zerr.Message()   // "lookup failed"
zerr.Unwrap()    // returns the wrapped error
```

## Standard Library Compatibility

`zerr` implements `Unwrap() error`, so it works seamlessly with Go's standard error handling:

```go
// errors.Is — check if a specific error is in the chain
if errors.Is(err, io.EOF) {
    // handle EOF
}

// errors.As — extract a specific error type from the chain
var zerr *zerr.ZError
if errors.As(err, &zerr) {
    fmt.Printf("Code: %d, Severity: %s\n", zerr.Code(), zerr.Severity())
}
```

## Severity Levels

| Level | Description |
| ------- | ------------- |
| `Debug` | Debugging information |
| `Info` | Informational messages |
| `Notice` | Normal but significant condition |
| `Warning` | Warning conditions |
| `Error` | Error conditions (default) |
| `Critical` | Critical conditions |
| `Fatal` | Fatal conditions |

## API Reference

### Constructors

| Function | Description |
| ---------- | ------------- |
| `New(msg)` | Create error with default severity |
| `NewS(severity, msg)` | Create error with severity |
| `NewC(code, msg)` | Create error with code |
| `NewSC(severity, code, msg)` | Create error with severity and code |
| `Errorf(format, v...)` | Formatted error with default severity |
| `ErrorSf(severity, format, v...)` | Formatted error with severity |
| `ErrorCf(code, format, v...)` | Formatted error with code |
| `ErrorSCf(severity, code, format, v...)` | Formatted error with severity and code |

### Wrappers

| Function | Description |
| ---------- | ------------- |
| `Wrap(err, msg)` | Wrap error with message |
| `WrapS(severity, err, msg)` | Wrap error with severity and message |
| `WrapC(code, err, msg)` | Wrap error with code and message |
| `WrapSC(severity, code, err, msg)` | Wrap error with severity, code and message |

### ZError Methods

| Method | Returns | Description |
| -------- | --------- | ------------- |
| `Error()` | `string` | Full formatted error string |
| `Severity()` | `Severity` | Error severity level |
| `Code()` | `int64` | Error code |
| `Message()` | `string` | Error message |
| `Unwrap()` | `error` | Wrapped underlying error |

### Helper Functions

| Function | Returns | Description |
| -------- | --------- | ------------- |
| `IsSeverity(err, sev)` | `bool` | Check if error (or any in chain) has severity |
| `GetCode(err)` | `(int64, bool)` | Get error code, false if not a zerr |
| `GetMessage(err)` | `string` | Get message or fallback to `err.Error()` |
| `Chain(err)` | `[]error` | Get the full error chain from outer to inner |
