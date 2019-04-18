package benchmarks

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testStruct struct {
}

func customAppendString(base []string, tail []string) []string {
	var ret []string
	if cap(base) >= len(base)+len(tail) {
		ret = base[:len(base)+len(tail)]
	} else {
		ret = make([]string, len(base)+len(tail))
		copy(ret[:len(base)], base)
	}

	copy(ret[len(base):], tail)
	return ret
}

func customAppendStruct(base []*testStruct, tail []*testStruct) []*testStruct {
	var ret []*testStruct
	if cap(base) >= len(base)+len(tail) {
		ret = base[:len(base)+len(tail)]
	} else {
		ret = make([]*testStruct, len(base)+len(tail))
		copy(ret[:len(base)], base)
	}

	copy(ret[len(base):], tail)
	return ret
}

func BenchmarkAppend(b *testing.B) {
	b.ReportAllocs()

	benchmarks := []struct {
		sourceLen int
		targetLen int
		stringLen int
	}{
		{
			sourceLen: 20,
			targetLen: 10,
			stringLen: 20,
		},
		{
			sourceLen: 500,
			targetLen: 500,
			stringLen: 20,
		},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v string slice", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]string, bm.sourceLen)
					target := make([]string, bm.targetLen)
					for i := range source {
						source[i] = RandStringRunes(bm.stringLen)
					}
					for i := range target {
						target[i] = RandStringRunes(bm.stringLen)
					}
					b.StartTimer()

					target = append(target, source...)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v string slice with preallocated capacity", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]string, bm.sourceLen)
					target := make([]string, bm.targetLen, bm.targetLen+bm.sourceLen)
					for i := range source {
						source[i] = RandStringRunes(bm.stringLen)
					}
					for i := range target {
						target[i] = RandStringRunes(bm.stringLen)
					}
					b.StartTimer()

					target = append(target, source...)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v pointer slice", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]*testStruct, bm.sourceLen)
					target := make([]*testStruct, bm.targetLen)
					for i := range source {
						source[i] = &testStruct{}
					}
					for i := range target {
						target[i] = &testStruct{}
					}
					b.StartTimer()

					target = append(target, source...)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v pointer slice with preallocated capacity", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]*testStruct, bm.sourceLen)
					target := make([]*testStruct, bm.targetLen, bm.targetLen+bm.sourceLen)
					for i := range source {
						source[i] = &testStruct{}
					}
					for i := range target {
						target[i] = &testStruct{}
					}
					b.StartTimer()

					target = append(target, source...)
				}
			})
	}
}

func BenchmarkCustomAppend(b *testing.B) {
	b.ReportAllocs()

	benchmarks := []struct {
		sourceLen int
		targetLen int
		stringLen int
	}{
		{
			sourceLen: 20,
			targetLen: 10,
			stringLen: 20,
		},
		{
			sourceLen: 500,
			targetLen: 500,
			stringLen: 20,
		},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v string slice", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]string, bm.sourceLen)
					target := make([]string, bm.targetLen)
					for i := range source {
						source[i] = RandStringRunes(bm.stringLen)
					}
					for i := range target {
						target[i] = RandStringRunes(bm.stringLen)
					}
					b.StartTimer()

					target = customAppendString(target, source)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v string slice with preallocated capacity", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]string, bm.sourceLen)
					target := make([]string, bm.targetLen, bm.targetLen+bm.sourceLen)
					for i := range source {
						source[i] = RandStringRunes(bm.stringLen)
					}
					for i := range target {
						target[i] = RandStringRunes(bm.stringLen)
					}
					b.StartTimer()

					target = customAppendString(target, source)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v pointer slice", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]*testStruct, bm.sourceLen)
					target := make([]*testStruct, bm.targetLen)
					for i := range source {
						source[i] = &testStruct{}
					}
					for i := range target {
						target[i] = &testStruct{}
					}
					b.StartTimer()

					target = customAppendStruct(target, source)
				}
			})

		b.Run(fmt.Sprintf("sourcelen: %v targetlen: %v pointer slice with preallocated capacity", bm.sourceLen, bm.targetLen),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					source := make([]*testStruct, bm.sourceLen)
					target := make([]*testStruct, bm.targetLen, bm.targetLen+bm.sourceLen)
					for i := range source {
						source[i] = &testStruct{}
					}
					for i := range target {
						target[i] = &testStruct{}
					}
					b.StartTimer()

					target = customAppendStruct(target, source)
				}
			})
	}
}

// random string generator from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
