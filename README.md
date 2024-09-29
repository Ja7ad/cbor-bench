# CBOR Benchmark

```shell
goos: linux
goarch: amd64
pkg: bench
cpu: AMD Ryzen 7 5700U with Radeon Graphics         
BenchmarkFxamacker_CBOREncode
BenchmarkFxamacker_CBOREncode-16        	 3737881	       322.1 ns/op	     160 B/op	       3 allocs/op
BenchmarkFxamacker_CBORDecode
BenchmarkFxamacker_CBORDecode-16        	 1679008	       707.0 ns/op	     589 B/op	       4 allocs/op
BenchmarkUgorji_CBOREncode
BenchmarkUgorji_CBOREncode-16           	 3404401	       353.9 ns/op	     272 B/op	       2 allocs/op
BenchmarkUgorji_CBORDecode
BenchmarkUgorji_CBORDecode-16           	 1943492	       616.0 ns/op	     605 B/op	       3 allocs/op
BenchmarkWhyrusleeping_CBOREncode
BenchmarkWhyrusleeping_CBOREncode-16    	 1686949	       712.2 ns/op	     272 B/op	      11 allocs/op
BenchmarkWhyrusleeping_CBORDecode
BenchmarkWhyrusleeping_CBORDecode-16    	  193090	      5810 ns/op	    8584 B/op	      58 allocs/op
PASS
```