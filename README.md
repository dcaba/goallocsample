# goallocsample
Just an example in regards to redeclaration of variables or reusing vars...

## Just declaring inside a loop or reuse a var from an external scope and clean it on each iter?

Looks its actually better to declare inside loops rather than inheriting a var from the external scope and caring about cleaning it up on each iteration:

```go
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dcaba/goallocsample
Benchmark_uniqueDeclaration-8             200000              9442 ns/op
Benchmark_declarationInsideLoop-8       20000000                71.3 ns/op
PASS
ok      github.com/dcaba/goallocsample  3.493s
```

Probably the reason is just how the compiler works (var declaration is not a high cost operation, and can actually
reusing registers) vs garbage collection...