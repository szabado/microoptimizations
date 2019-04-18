package benchmarks

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
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
		b.Run(fmt.Sprintf("sourcelen_%v__targetlen_%v__strings", bm.sourceLen, bm.targetLen), func(b *testing.B) {
			b.Run("append", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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

			b.Run("customAppend", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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
		})

		b.Run(fmt.Sprintf("sourcelen_%v__targetlen_%v__string__preallocated", bm.sourceLen, bm.targetLen), func(b *testing.B) {
			b.Run("append", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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

			b.Run("customAppend", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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
		})

		b.Run(fmt.Sprintf("sourcelen_%v__targetlen_%v__pointers", bm.sourceLen, bm.targetLen), func(b *testing.B) {
			b.Run("append", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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

			b.Run("customAppend", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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
		})

		b.Run(fmt.Sprintf("sourcelen_%v__targetlen_%v__pointers__preallocated", bm.sourceLen, bm.targetLen), func(b *testing.B) {
			b.Run("append", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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

			b.Run("customAppend", func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()
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
		})
	}
}

func BenchmarkStringBuild(b *testing.B) {
	bms := []struct {
		numSegments    int
		segmentLength  int
		clearFrequency int
	}{
		{
			numSegments:    20,
			segmentLength:  20,
			clearFrequency: 2,
		},
		{
			numSegments:    20,
			segmentLength:  20,
			clearFrequency: 5,
		},
		{
			numSegments:    20,
			segmentLength:  20,
			clearFrequency: 10,
		},
		{
			numSegments:    20,
			segmentLength:  20,
			clearFrequency: 20,
		},
		{
			numSegments:    20,
			segmentLength:  20,
			clearFrequency: 5000,
		},
	}

	for _, bm := range bms {
		b.Run(fmt.Sprintf("segments_%v__segmentLength_%v", bm.numSegments, bm.segmentLength), func(b *testing.B) {
			segments := make([]string, bm.numSegments)
			for i := 0; i < bm.numSegments; i++ {
				segments[i] = RandStringRunes(bm.segmentLength)
			}

			b.Run("bytes.Buffer", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf bytes.Buffer
					for _, seg := range segments {
						buf.WriteString(seg)
					}
				}
			})

			b.Run("strings.Builder", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf strings.Builder
					for _, seg := range segments {
						buf.WriteString(seg)
					}
				}
			})
		})

		b.Run(fmt.Sprintf("segments_%v__segmentLength_%v__clearsFrequency_%v", bm.numSegments, bm.segmentLength, bm.clearFrequency), func(b *testing.B) {
			segments := make([]string, bm.numSegments)
			for i := 0; i < bm.numSegments; i++ {
				segments[i] = RandStringRunes(bm.segmentLength)
			}

			b.Run("bytes.Buffer", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf bytes.Buffer
					for i, seg := range segments {
						buf.WriteString(seg)

						rem := (i + 1) % bm.clearFrequency
						if rem == 1 {
							buf.Reset()
						}
					}
				}
			})

			b.Run("strings.Builder", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf strings.Builder
					for _, seg := range segments {
						buf.WriteString(seg)

						rem := (i + 1) % bm.clearFrequency
						if rem == 1 {
							buf.Reset()
						}
					}
				}
			})

			b.Run("bytes.Buffer_with_timer_manip", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf bytes.Buffer
					for i, seg := range segments {
						buf.WriteString(seg)

						b.StopTimer()
						rem := (i + 1) % bm.clearFrequency
						if rem == 1 {
							buf.Reset()
						}
						b.StartTimer()
					}
				}
			})

			b.Run("strings.Builder_with_timer_manip", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf strings.Builder
					for _, seg := range segments {
						buf.WriteString(seg)

						b.StopTimer()
						rem := (i + 1) % bm.clearFrequency
						if rem == 1 {
							buf.Reset()
						}
						b.StartTimer()
					}
				}
			})
		})
	}
}

func BenchmarkTimerStopStarting(b *testing.B) {
	var buf strings.Builder
	b.ReportAllocs()
	b.Run("builderWithoutTimer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf.WriteString("123123123123123")
			buf.Reset()
		}
	})
	b.Run("builderWithTimer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf.WriteString("123123123123123")
			buf.Reset()
			b.StopTimer()
			b.StartTimer()
		}
	})
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
