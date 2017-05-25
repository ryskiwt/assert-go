## Usage
```go
import "github.com/ryskiwt/assert-go"

v := 1
assert.Do(v == 0)
// -> panic: assrtion error

assert.Equals(v, 0)
// -> panic: assrtion error: got=1, want=0

assert.WithMsg(v == 0, "invalid v")
// -> panic: invalid v

assert.WithMsgf(v == 0, "invalid v: got=%v, want=%v", v, 0)
// -> panic: invalid v: got=1, want=0
```

## Build

### Build for Debug
```sh
$ go build
```

### Build for Release
```sh
$ go build -tags=release
```

## Benchmark

```
$ go test -run NONE -bench . -benchmem
testing: warning: no tests to run
BenchmarkNone-4             	2000000000	         0.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkDo_True-4          	2000000000	         1.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkEquals_True-4      	20000000	        58.7 ns/op	      16 B/op	       2 allocs/op
BenchmarkWithMsg_True-4     	30000000	        43.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkWithMsgf_True-4    	30000000	        55.7 ns/op	      16 B/op	       2 allocs/op
BenchmarkDo_False-4         	panic: assertion error

goroutine 33 [running]:
panic(0xe5b40, 0xc420480000)
	/usr/local/Cellar/go/1.7.3/libexec/src/runtime/panic.go:500 +0x1a1
github.com/ryskiwt/assert-go.Do(0x2b29800)
	/path/to/repos/src/github.com/ryskiwt/assert-go/debug.go:11 +0x80
github.com/ryskiwt/assert-go_test.BenchmarkDo_False(0xc42009e280)
	/path/to/repos/src/github.com/ryskiwt/assert-go/benchmark_test.go:45 +0x3e
testing.(*B).runN(0xc42009e280, 0x1)
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:139 +0xaa
testing.(*B).run1.func1(0xc42009e280)
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:208 +0x5a
created by testing.(*B).run1
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:209 +0x7f
exit status 2
FAIL	github.com/ryskiwt/assert-go	9.256s
```

```
$ go test -run NONE -bench . -benchmem -tags=release
testing: warning: no tests to run
BenchmarkNone-4             	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkDo_True-4          	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkEquals_True-4      	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsg_True-4     	100000000	        13.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsgf_True-4    	100000000	        19.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDo_False-4         	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkEquals_False-4     	100000000	        14.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsg_False-4    	100000000	        13.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsgf_False-4   	100000000	        19.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ryskiwt/assert-go	11.903s
```
