
---

# Testing in Go

* Built into the standard library
* Tests live in there own files close to production code
* Test code is not included in the binary

---

## One command for

* unit tests
* benchmarks
* examples
* coverage analysis
* race detection

<!-- pause -->

```bash
    go test {-cover}
```

## Go test

<!--
Weet niet of deze slide nodig is.
-->

Running `go test` will

1. Compile your normal package code
2. Compiles your _test.go files
3. Create a temporary test executable
4. Runs that executable
5. Discard the executable

---

## Differences with Python

Go intentionally has no assert keyword. 

The Go philosophy is:

> Tests are programs. Use normal language constructs.

---

## How to write tests

Test files end in _test.go and live in the same directory as the code they test.

```text
 myapp/
  │ ├── integers/
  │ │   ├── addition.go
  │ │   └── addition_test.go
  │ ├── main.go
  │ └── go.mod
```

---

## How to write tests

---

