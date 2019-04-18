# Benchmarking Go Micro-Optimizations

A collection of benchmarks designed to answer all those questions you have about the "fastest" way to write code.
These are all benchmarks parts of code that are unlikely to _ever_ be the bottle neck of your application, it's mainly
for fun.

Take all of these with a grain of salt, they're just for fun.

Have your own benchmark that you want to add? Hit me up!

## How to run them

Normally `go test -bench=. .` would work, but the tests can time out because the number of iterations gets too high.
Try running it with `go  test -bench=. -benchtime=100000x .`, and tweak the number of iterations as you see fit!

## Comparisons

### Slice Concatenation
```
sl1 := append(sl1, sl2)
```

**vs**

```
func customAppend(sl1, sl2 []string) []string {
	var ret []string
	if cap(base) >= len(base) + len(tail) {
		ret = base[:len(base) + len(tail)]
	} else {
		ret = make([]string, len(base) + len(tail))
		copy(ret[:len(base)], base)
	}

	copy(ret[len(base):], tail)
	return ret
}

sl1 = customAppend(sl1, sl2)
```

#### Results

```
$ go test -bench=. -benchtime=10000x .
goos: darwin
goarch: amd64
pkg: github.com/szabado/microoptimizations
BenchmarkAppend/sourcelen:_20_targetlen:_10_string_slice-12         	   10000	       560 ns/op
BenchmarkAppend/sourcelen:_20_targetlen:_10_string_slice_with_preallocated_capacity-12         	   10000	       135 ns/op
BenchmarkAppend/sourcelen:_20_targetlen:_10_pointer_slice-12                                   	   10000	       495 ns/op
BenchmarkAppend/sourcelen:_20_targetlen:_10_pointer_slice_with_preallocated_capacity-12        	   10000	       122 ns/op
BenchmarkAppend/sourcelen:_500_targetlen:_500_string_slice-12                                  	   10000	      3313 ns/op
BenchmarkAppend/sourcelen:_500_targetlen:_500_string_slice_with_preallocated_capacity-12       	   10000	       444 ns/op
BenchmarkAppend/sourcelen:_500_targetlen:_500_pointer_slice-12                                 	   10000	      1799 ns/op
BenchmarkAppend/sourcelen:_500_targetlen:_500_pointer_slice_with_preallocated_capacity-12      	   10000	       228 ns/op
BenchmarkCustomAppend/sourcelen:_20_targetlen:_10_string_slice-12                              	   10000	       481 ns/op
BenchmarkCustomAppend/sourcelen:_20_targetlen:_10_string_slice_with_preallocated_capacity-12   	   10000	       146 ns/op
BenchmarkCustomAppend/sourcelen:_20_targetlen:_10_pointer_slice-12                             	   10000	       449 ns/op
BenchmarkCustomAppend/sourcelen:_20_targetlen:_10_pointer_slice_with_preallocated_capacity-12  	   10000	       113 ns/op
BenchmarkCustomAppend/sourcelen:_500_targetlen:_500_string_slice-12                            	   10000	      3207 ns/op
BenchmarkCustomAppend/sourcelen:_500_targetlen:_500_string_slice_with_preallocated_capacity-12 	   10000	       453 ns/op
BenchmarkCustomAppend/sourcelen:_500_targetlen:_500_pointer_slice-12                           	   10000	      1805 ns/op
BenchmarkCustomAppend/sourcelen:_500_targetlen:_500_pointer_slice_with_preallocated_capacity-12         	   10000	       232 ns/op
PASS
ok  	github.com/szabado/microoptimizations	32.915s```
```

Oh baby, look at those _savings_. Between 14% performance increase and an 8% performance decrease! Gotta love the
variability.

**Outcome:** Use append, and try to pre-allocate enough space when you can. It makes a _huge_ difference. If you happen
to need to squeeze that extra 50ns of performance out of your code... there's probably somewhere easier you can find it.
