# Benchmarking Go Micro-Optimizations

A collection of benchmarks designed to answer all those questions you have about the "fastest" way to write code.
These are all benchmarks parts of code that are unlikely to _ever_ be the bottle neck of your application, it's mainly
for fun.

Take all of these with a grain of salt.

Have your own benchmark that you want to add? Hit me up!

## How to run them

`go test -bench=. .` runs all of them, if you're only interested in a subset then throw a `-bench=Benchmark<name>` in
there and Bob's your uncle.

## Comparisons

### String Building

`strings.Builder` vs `bytes.Buffer` vs `strings.Join` vs `fmt.Sprintf` vs `fmt.Sprintf`

`bytes.Buffer` was the go to (haha) when I was learning Go, but since 1.10 `strings.Builder` has existed. I figured
it'd be worth benchmarking them. [Prior work](https://medium.com/@felipedutratine/string-concatenation-in-golang-since-1-10-bytes-buffer-vs-strings-builder-2b3081848c45)
only considered the writing, not the `Reset` method.

I've also wondered which if those two can outperform `strings.Join` and `fmt.Sprint(f)`.

#### Results

```
go test -bench=BenchmarkStringBuild  .
goos: darwin
goarch: amd64
pkg: github.com/szabado/microoptimizations
BenchmarkStringBuild/segments_20__segmentLength_20/bytes.Buffer-12         	 3000000	       527 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/strings.Builder-12      	 5000000	       343 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/bytes.Buffer_with_size_mgmt-12         	 5000000	       268 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/strings.Builder_with_size_mgmt-12      	10000000	       162 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/strings.Join-12                        	10000000	       236 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/strings.Join_with_slice_mgmt-12        	 5000000	       349 ns/op	     736 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/fmt.Sprintf-12                         	 3000000	       458 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20/fmt.Sprint-12                          	 3000000	       490 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/bytes.Buffer-12     	 5000000	       375 ns/op	      64 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_2/strings.Builder-12  	 2000000	       628 ns/op	     816 B/op	      12 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/bytes.Buffer-12                     	 3000000	       609 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/strings.Builder-12                  	 5000000	       357 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/bytes.Buffer_with_size_mgmt-12      	 5000000	       282 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/strings.Builder_with_size_mgmt-12   	10000000	       171 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/strings.Join-12                     	10000000	       238 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/strings.Join_with_slice_mgmt-12     	 5000000	       356 ns/op	     736 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/fmt.Sprintf-12                      	 3000000	       483 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#01/fmt.Sprint-12                       	 3000000	       503 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/bytes.Buffer-12     	 3000000	       461 ns/op	     224 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5/strings.Builder-12  	 3000000	       570 ns/op	     921 B/op	       8 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/bytes.Buffer-12                     	 3000000	       542 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/strings.Builder-12                  	 5000000	       331 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/bytes.Buffer_with_size_mgmt-12      	 5000000	       261 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/strings.Builder_with_size_mgmt-12   	10000000	       164 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/strings.Join-12                     	10000000	       237 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/strings.Join_with_slice_mgmt-12     	 5000000	       356 ns/op	     736 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/fmt.Sprintf-12                      	 3000000	       474 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#02/fmt.Sprint-12                       	 3000000	       490 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/bytes.Buffer-12    	 3000000	       566 ns/op	     544 B/op	       3 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_10/strings.Builder-12 	 3000000	       539 ns/op	     956 B/op	       6 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/bytes.Buffer-12                     	 3000000	       530 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/strings.Builder-12                  	 5000000	       326 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/bytes.Buffer_with_size_mgmt-12      	 5000000	       257 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/strings.Builder_with_size_mgmt-12   	10000000	       158 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/strings.Join-12                     	10000000	       231 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/strings.Join_with_slice_mgmt-12     	 5000000	       340 ns/op	     736 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/fmt.Sprintf-12                      	 3000000	       464 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#03/fmt.Sprint-12                       	 3000000	       492 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/bytes.Buffer-12    	 2000000	       709 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_20/strings.Builder-12 	 3000000	       516 ns/op	     974 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/bytes.Buffer-12                     	 3000000	       529 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/strings.Builder-12                  	 5000000	       323 ns/op	     992 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/bytes.Buffer_with_size_mgmt-12      	 5000000	       260 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/strings.Builder_with_size_mgmt-12   	10000000	       163 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/strings.Join-12                     	10000000	       233 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/strings.Join_with_slice_mgmt-12     	 5000000	       356 ns/op	     736 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/fmt.Sprintf-12                      	 3000000	       467 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20#04/fmt.Sprint-12                       	 3000000	       485 ns/op	     416 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/bytes.Buffer-12  	 2000000	       749 ns/op	    1248 B/op	       4 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_20__clearsFrequency_5000/strings.Builder-12         	 3000000	       520 ns/op	     991 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/bytes.Buffer-12                                 	  300000	      3632 ns/op	   30720 B/op	       5 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/strings.Builder-12                              	  300000	      4102 ns/op	   39168 B/op	       9 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/bytes.Buffer_with_size_mgmt-12                  	 1000000	      1178 ns/op	   10240 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/strings.Builder_with_size_mgmt-12               	 1000000	      1090 ns/op	   10240 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/strings.Join-12                                 	 1000000	      1185 ns/op	   10240 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/strings.Join_with_slice_mgmt-12                 	 1000000	      1335 ns/op	   10560 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/fmt.Sprintf-12                                  	 1000000	      1597 ns/op	   10248 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500/fmt.Sprint-12                                   	 1000000	      1541 ns/op	   10248 B/op	       1 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500__clearsFrequency_2/bytes.Buffer-12              	 2000000	       856 ns/op	    2048 B/op	       2 allocs/op
BenchmarkStringBuild/segments_20__segmentLength_500__clearsFrequency_2/strings.Builder-12           	  500000	      3259 ns/op	   24704 B/op	      14 allocs/op
PASS
ok  	github.com/szabado/microoptimizations	117.051s
```

I'm thoroughly shocked by these results. Let's break down the things the benchmark shows:
- `strings.Builder` outperforms `bytes.Buffer` when building strings.
    - This isn't true if you're passing large strings into `WriteString()`, or if you're calling `Reset()` a lot. In
    both of these cases, `bytes.Buffer` outperforms `strings.Builder` by a healthy margin.
- `fmt.Sprintf` is generally faster than `fmt.Sprint`
- `strings.Join` is _blazing_ fast, and takes at least 100ns less than most other methods.
    - This is a naive benchmark. Adding in the overhead of building the slice first makes it slower than 
    `strings.Builder` in most cases (see the `strings.Join_with_slice_mgmt` tests).
    - If you pre-allocate the size of the buffer, `strings.Builder` outperforms `strings.Join` and `bytes.Buffer`
    is slightly slower than it.
