# slow-addcleanup-repro

For golang/go issue [#75188](https://github.com/golang/go/issues/75188)

Benchmark code inspired by [Daniel Lemire's](https://lemire.me/blog/2023/05/19/the-absurd-cost-of-finalizers-in-go/) article

Tests executed on Go 1.25.0

### darwin/arm64 - Macbook Pro M1 14"
```
go test -bench=. ./...
goos: darwin
goarch: arm64
pkg: example.com/slow-addcleanup-repro
cpu: Apple M1 Pro
BenchmarkAllocateFree-8                 15169078                73.42 ns/op
BenchmarkAllocateDefer-8                15890224                75.04 ns/op
BenchmarkAllocateAddCleanup-8            3516811               359.7 ns/op
BenchmarkAllocateWithFinalizer-8         1739244               676.9 ns/op
PASS
ok      example.com/slow-addcleanup-repro       6.186s
```

### linux/amd64 - powerful VM
```
BenchmarkAllocateFree-166               12526917                94.33 ns/op
BenchmarkAllocateDefer-166              12716973                98.02 ns/op
BenchmarkAllocateAddCleanup-166          2054509               567.0 ns/op
BenchmarkAllocateWithFinalizer-166        996853              1058 ns/op
PASS
```
