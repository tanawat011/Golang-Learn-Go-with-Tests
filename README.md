# Learn Go with Tests — my exercises

My work through the first chapters of
[Learn Go with Tests](https://github.com/quii/learn-go-with-tests) by Chris James:
write a failing test, make it pass, then refactor.

| Folder | Chapter | What the tests drive |
|---|---|---|
| `Hello` | Hello, World | constants, `if`, `switch`, named return values; greetings in English, Thai and French |
| `Integers` | Integers | a first function, and an `Example` that `go test` checks |
| `Iteration` | Iteration | a `for` loop and a benchmark |
| `ArrayAndSlices` | Arrays and slices | `range`, variadic functions, `reflect.DeepEqual` |
| `StructMethodsAndInterface` | Structs, methods and interfaces | `Shape` with `Area()`, table-driven tests |
| `PointersAndErrors` | Pointers and errors | a `Wallet` with pointer receivers and an error value |
| `FizzBuzz` | (extra) | FizzBuzz, tested from a table |

Originally written in 2019 while learning Go with the book above; refreshed
in Nov 2026 (Go modules, current Go, gofmt, a few tests the book's later
editions would add, CI).

## Run the tests

Needs Go 1.26 or newer.

```sh
git clone https://github.com/tanawat011/Golang-Learn-Go-with-Tests.git
cd Golang-Learn-Go-with-Tests
go test ./...                 # every chapter
go test -v ./Hello            # one chapter, each case by name
go test -cover ./...          # coverage (100% in every package)
go test -run '^$' -bench . ./Iteration
```

CI runs gofmt, `go vet`, the tests with the race detector and the
benchmarks once on the two supported Go releases.

## Credit

The exercises follow the book, whose code is MIT licensed
(Copyright (c) Christopher James).
