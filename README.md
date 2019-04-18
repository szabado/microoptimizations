# Benchmarking Go Micro-Optimizations

A collection of benchmarks designed to answer all those questions you have about the "fastest" way to write code.
These are all benchmarks parts of code that are unlikely to _ever_ be the bottle neck of your application, it's mainly
for fun.

Take all of these with a grain of salt.

Have your own benchmark that you want to add? Hit me up!

## How to run them

Normally `go test -bench=. .` would work, but the tests can time out because the number of iterations gets too high.
Try running it with `go  test -bench=. -benchtime=10000x .`, and tweak the number of iterations as you see fit!
If you only want to run some of the benchmarks, throw a `-bench=Benchmark<name>` in there and Bob's your uncle.

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
$ go test -bench=BenchmarkAppend -benchtime=10000x .
goos: darwin
goarch: amd64
pkg: github.com/szabado/microoptimizations
BenchmarkAppend/sourcelen_20__targetlen_10__strings/append-12  	   10000	       564 ns/op	     480 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__strings/customAppend-12         	   10000	       554 ns/op	     480 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__string__preallocated/append-12  	   10000	       184 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__string__preallocated/customAppend-12         	   10000	       168 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__pointers/append-12                           	   10000	       525 ns/op	     240 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__pointers/customAppend-12                     	   10000	       514 ns/op	     240 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__pointers__preallocated/append-12             	   10000	       116 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_20__targetlen_10__pointers__preallocated/customAppend-12       	   10000	       135 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__strings/append-12                          	   10000	      3244 ns/op	   16384 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__strings/customAppend-12                    	   10000	      3180 ns/op	   16384 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__string__preallocated/append-12             	   10000	       495 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__string__preallocated/customAppend-12       	   10000	       526 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__pointers/append-12                         	   10000	      2227 ns/op	    8192 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__pointers/customAppend-12                   	   10000	      2015 ns/op	    8192 B/op	       1 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__pointers__preallocated/append-12           	   10000	       263 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppend/sourcelen_500__targetlen_500__pointers__preallocated/customAppend-12     	   10000	       247 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/szabado/microoptimizations	31.571s
```

Oh baby, look at those _savings_. Between 14% performance increase and an 8% performance decrease! Gotta love the
variability.

**Outcome:** Use append, because the standard library is always better where possible.
Try to pre-allocate enough space when you can. It makes a _huge_ difference. If you happen
to need to squeeze that extra 50ns of performance out of your code... there's probably somewhere easier you can find it.

### String Building

```
strings.Builder
```

**vs**

```
bytes.Buffer
```

#### Results

This one is actually interesting! `bytes.Buffer` was the go to (haha) when I was learning Go, but since 1.10
`strings.Builder` has existed. I figured it'd be worth benchmarking them. [Prior work](https://medium.com/@felipedutratine/string-concatenation-in-golang-since-1-10-bytes-buffer-vs-strings-builder-2b3081848c45)
only considered the writing, not the `Reset` method.

```
go test -bench=BenchmarkStringBuild -benchtime=10000x .
goos: darwin
goarch: amd64
pkg: github.com/szabado/microoptimizations
BenchmarkStringBuild/segments_20__segmentLength_20/bytes.Buffer-12         	   10000	       835 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/strings.Builder-12      	   10000	       381 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/bytes.Buffer-12         	   10000	       421 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/strings.Builder-12      	   10000	       691 ns/op	     816 B/op	      12 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/bytes.Buffer_with_timer_manip-12         	   10000	      2799 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/strings.Builder_with_timer_manip-12      	   10000	      5731 ns/op	     816 B/op	      12 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/bytes.Buffer-12                                          	   10000	       583 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/strings.Builder-12                                       	   10000	       310 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/bytes.Buffer-12                          	   10000	       437 ns/op	     224 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/strings.Builder-12                       	   10000	       586 ns/op	     921 B/op	       8 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/bytes.Buffer_with_timer_manip-12         	   10000	      3486 ns/op	     224 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/strings.Builder_with_timer_manip-12      	   10000	      4641 ns/op	     921 B/op	       8 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/bytes.Buffer-12                                          	   10000	       631 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/strings.Builder-12                                       	   10000	       578 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/bytes.Buffer-12                         	   10000	       616 ns/op	     544 B/op	       3 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/strings.Builder-12                      	   10000	       733 ns/op	     956 B/op	       6 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/bytes.Buffer_with_timer_manip-12        	   10000	      4037 ns/op	     544 B/op	       3 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/strings.Builder_with_timer_manip-12     	   10000	      4616 ns/op	     956 B/op	       6 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/bytes.Buffer-12                                          	   10000	       591 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/strings.Builder-12                                       	   10000	       334 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/bytes.Buffer-12                         	   10000	       764 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/strings.Builder-12                      	   10000	       634 ns/op	     974 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/bytes.Buffer_with_timer_manip-12        	   10000	      4445 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/strings.Builder_with_timer_manip-12     	   10000	      4301 ns/op	     974 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/bytes.Buffer-12                                          	   10000	       579 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/strings.Builder-12                                       	   10000	       314 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/bytes.Buffer-12                       	   10000	       696 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/strings.Builder-12                    	   10000	       479 ns/op	     991 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/bytes.Buffer_with_timer_manip-12      	   10000	      4538 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/strings.Builder_with_timer_manip-12   	   10000	      4030 ns/op	     991 B/op	       5 allocs/op
PASS
ok  	github.com/szabado/microoptimizations	53.100s
```

If you're clearing the buffer you should be using `bytes.Buffer`, since it actually re-uses memory.
There's an interesting corollary I discovered in this benchmarking: `b.StartTimer()` and `b.StopTimer()` have
an incredibly non-zero effect on benchmarks (the `with_timer_manip` tests call those methods).

**tl;dr:** if you're building a single string, use `string.Builder`. If you're building multiple and calling `Reset`,
`bytes.Buffer` is the way to go.
