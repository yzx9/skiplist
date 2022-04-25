# Skip List

This repo implements [skiplist](https://homepage.divms.uiowa.edu/~ghosh/skip.pdf) alg over 100~lines.

## Benchmark

Skiplist:

```
goos: linux
goarch: amd64
pkg: github.com/yzx9/skiplist
cpu: AMD Ryzen 7 5700G with Radeon Graphics         
BenchmarkSkipListGet100-16               1000000              1063 ns/op
BenchmarkSkipListGet1000-16                62349             34864 ns/op
BenchmarkSkipListGet10000-16                1848            674151 ns/op
BenchmarkSkipListGet100000-16                 63          16840410 ns/op
```

Dullist:

```
goos: linux
goarch: amd64
pkg: github.com/yzx9/skiplist/dullist
cpu: AMD Ryzen 7 5700G with Radeon Graphics         
BenchmarkDullistGet100-16                 415412              3205 ns/op
BenchmarkDullistGet1000-16                  2754            441719 ns/op
BenchmarkDullistGet10000-16                   10         105793238 ns/op
BenchmarkDullistGet100000-16                   1       39059786993 ns/op
```
