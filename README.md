## Usage
```go
import "github.com/ryskiwt/assert-go"

v := 1
assert.Do(v == 0)
// -> panic: assrtion error

assert.WithMsg(v == 0, "invalid v")
// -> panic: assrtion error: invalid v

assert.WithMsgf(v == 0, "invalid v: want=%v, got=%v", 0, v)
// -> panic: assrtion error: invalid v: want=0, got=1
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
$ go test -run NONE -bench . -benchmem -tags=release
testing: warning: no tests to run
BenchmarkWithMsgf_True-4    	20000000	        64.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithMsgf_False-4   	20000000	        61.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ryskiwt/assert-go	2.662s
```

```
$ go test -run NONE -bench . -benchmem
testing: warning: no tests to run
BenchmarkWithMsgf_True-4    	10000000	       222 ns/op	      80 B/op	       5 allocs/op
panic: assertion error: 0000000000, 0000000000, 0000000000, 0000000000, 0000000000

goroutine 4 [running]:
panic(0xe55a0, 0xc4204840a0)
...
exit status 2
FAIL	github.com/ryskiwt/assert-go	2.462s
```
