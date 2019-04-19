package benchmarks

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

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
		{
			numSegments:    20,
			segmentLength:  500,
			clearFrequency: 2,
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

			b.Run("bytes.Buffer_with_size_mgmt", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf bytes.Buffer
					buf.Grow(bm.numSegments * bm.segmentLength)
					for _, seg := range segments {
						buf.WriteString(seg)
					}
				}
			})

			b.Run("strings.Builder_with_size_mgmt", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var buf strings.Builder
					buf.Grow(bm.numSegments * bm.segmentLength)
					for _, seg := range segments {
						buf.WriteString(seg)
					}
				}
			})

			b.Run("strings.Join", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					strings.Join(segments, "")
				}
			})

			b.Run("strings.Join_with_slice_mgmt", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					tempSegments := make([]string, 0, bm.numSegments)
					for _, seg := range segments {
						tempSegments = append(tempSegments, seg)
					}

					strings.Join(tempSegments, "")
				}
			})

			var formatBuilder strings.Builder
			genericSegments := make([]interface{}, bm.numSegments)
			for i := 0; i < bm.numSegments; i++ {
				formatBuilder.WriteString("%s")
				genericSegments[i] = segments[i]
			}
			format := formatBuilder.String()

			b.Run("fmt.Sprintf", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					fmt.Sprintf(format, genericSegments...)
				}
			})

			b.Run("fmt.Sprint", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					fmt.Sprint(genericSegments...)
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
