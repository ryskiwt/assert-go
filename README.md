## Usage
```go
import "github.com/ryskiwt/assert-go"

v := 1

// CASE 1 (about 0.33 ns/op)
if assert.ENABLED && v == 0 {
	panic(assert.DefaultMsg)
}
// -> panic: assrtion error


// CASE 2 (about 0.33 ns/op)
assert.Do(v == 0)
// -> panic: assrtion error


// CASE 3 (about 1.00 ns/op)
assert.Equals(func() (assert.G, assert.W) { return v, 0 })
// -> panic: assrtion error: got=1, want=0


// CASE 4 (about 1.00 ns/op)
assert.WithMsg(func() (bool, string) {
	return v == 0, assert.NewMsg(v, 0)
})
// -> panic: assertion error: got=1, want=0
```

## Build

### assertion enabled
```sh
$ go build -tags=assert
```

### assertion disabled
```sh
$ go build
```

## Benchmark

```
$ go test -run NONE -bench . -benchmem -tags=assert                                                                                                       +[develop]
testing: warning: no tests to run
BenchmarkNone-4            	2000000000	         0.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlag_True-4       	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkDo_True-4         	1000000000	         1.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkEquals_True-4     	30000000	        58.8 ns/op	      16 B/op	       2 allocs/op
BenchmarkWithMsg_True-4    	 5000000	       335 ns/op	      64 B/op	       4 allocs/op
BenchmarkDo_False-4        	panic: assertion error

goroutine 27 [running]:
panic(0xe5c40, 0xc420490030)
	/usr/local/Cellar/go/1.7.3/libexec/src/runtime/panic.go:500 +0x1a1
pgithub.com/ryskiwt/assert-go.Do(0x16fedc00)
	/path/to/repos/src/github.com/ryskiwt/assert-go/debug.go:9 +0x80
pgithub.com/ryskiwt/assert-go_test.BenchmarkDo_False(0xc4200ac280)
	/path/to/repos/src/github.com/ryskiwt/assert-go/benchmark_test.go:47 +0x3e
testing.(*B).runN(0xc4200ac280, 0x1)
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:139 +0xaa
testing.(*B).run1.func1(0xc4200ac280)
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:208 +0x5a
created by testing.(*B).run1
	/usr/local/Cellar/go/1.7.3/libexec/src/testing/benchmark.go:209 +0x7f
exit status 2
FAIL	github.com/ryskiwt/assert-go	7.332s
```

```
$ go test -run NONE -bench . -benchmem                                                                                                                    +[develop]
testing: warning: no tests to run
BenchmarkNone-4            	2000000000	         0.31 ns/op	       0 B/op	       0 allocs/op // no assertion
BenchmarkFlag_True-4       	2000000000	         0.33 ns/op	       0 B/op	       0 allocs/op // fastest
BenchmarkDo_True-4         	2000000000	         0.31 ns/op	       0 B/op	       0 allocs/op // fastest
BenchmarkEquals_True-4     	2000000000	         1.01 ns/op	       0 B/op	       0 allocs/op // 3 times slower
BenchmarkWithMsg_True-4    	2000000000	         0.96 ns/op	       0 B/op	       0 allocs/op // 3 times slower
BenchmarkDo_False-4        	2000000000	         0.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlag_False-4      	2000000000	         0.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkEquals_False-4    	2000000000	         1.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsg_False-4   	2000000000	         1.03 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ryskiwt/assert-go	12.039s
```
