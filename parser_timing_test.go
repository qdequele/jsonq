package jsonq

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkParseRawString(b *testing.B) {
	for _, s := range []string{`""`, `"a"`, `"abcd"`, `"abcdefghijk"`, `"qwertyuiopasdfghjklzxcvb"`} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseRawString(b, s)
		})
	}
}

func benchmarkParseRawString(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, tail, err := parseRawString(s)
			if err != nil {
				panic(fmt.Errorf("cannot parse %q: %s", s, err))
			}
			if rs != s[1:len(s)-1] {
				panic(fmt.Errorf("invalid string obtained; got %q; want %q", rs, s[1:len(s)-1]))
			}
			if len(tail) > 0 {
				panic(fmt.Errorf("non-empty tail got: %q", tail))
			}
		}
	})
}

func BenchmarkParseRawNumber(b *testing.B) {
	for _, s := range []string{"1", "1234", "123456", "-1234", "1234567890.1234567", "-1.32434e+12"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseRawNumber(b, s)
		})
	}
}

func benchmarkParseRawNumber(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rn, tail, err := parseRawNumber(s)
			if err != nil {
				panic(fmt.Errorf("cannot parse %q: %s", s, err))
			}
			if rn != s {
				panic(fmt.Errorf("invalid number obtained; got %q; want %q", rn, s))
			}
			if len(tail) > 0 {
				panic(fmt.Errorf("non-empty tail got: %q", tail))
			}
		}
	})
}

func BenchmarkObjectGetBig(b *testing.B) {
	for _, itemsCount := range []int{10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("items_%d", itemsCount), func(b *testing.B) {
			for _, lookupsCount := range []int{0, 1, 2, 4, 8, 16, 32, 64} {
				b.Run(fmt.Sprintf("lookups_%d", lookupsCount), func(b *testing.B) {
					benchmarkObjectGetBig(b, itemsCount, lookupsCount)
				})
			}
		})
	}
}

func benchmarkObjectGetBig(b *testing.B, itemsCount, lookupsCount int) {
	b.StopTimer()
	var ss []string
	for i := 0; i < itemsCount; i++ {
		s := fmt.Sprintf(`"key_%d": "value_%d"`, i, i)
		ss = append(ss, s)
	}
	s := "{" + strings.Join(ss, ",") + "}"
	key := fmt.Sprintf("key_%d", len(ss)/2)
	expectedValue := fmt.Sprintf("value_%d", len(ss)/2)
	b.StartTimer()
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))

	b.RunParallel(func(pb *testing.PB) {
		var p Parser
		for pb.Next() {
			v, err := p.Parse(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			o := v.GetObject()
			for i := 0; i < lookupsCount; i++ {
				sb := o.Get(key).GetStringBytes()
				if string(sb) != expectedValue {
					panic(fmt.Errorf("unexpected value; got %q; want %q", sb, expectedValue))
				}
			}
		}
	})
}

func BenchmarkParse(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchmarkParse(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchmarkParse(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchmarkParse(b, largeFixture)
	})
}

func benchmarkParse(b *testing.B, s string) {
	b.Run("stdjson-map", func(b *testing.B) {
		benchmarkStdJSONParseMap(b, s)
	})
	b.Run("stdjson-struct", func(b *testing.B) {
		benchmarkStdJSONParseStruct(b, s)
	})
	b.Run("stdjson-empty-struct", func(b *testing.B) {
		benchmarkStdJSONParseEmptyStruct(b, s)
	})
	b.Run("fastjson", func(b *testing.B) {
		benchmarkFastJSONParse(b, s)
	})
}

func benchmarkFastJSONParse(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		var p Parser
		for pb.Next() {
			v, err := p.Parse(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			if v.Type() != TypeObject {
				panic(fmt.Errorf("unexpected value type; got %s; want %s", v.Type(), TypeObject))
			}
		}
	})
}

func benchmarkStdJSONParseMap(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	bb := s2b(s)
	b.RunParallel(func(pb *testing.PB) {
		var m map[string]interface{}
		for pb.Next() {
			if err := json.Unmarshal(bb, &m); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
		}
	})
}

func benchmarkStdJSONParseStruct(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	bb := s2b(s)
	b.RunParallel(func(pb *testing.PB) {
		var m struct {
			Sid     int
			UUID    string
			Person  map[string]interface{}
			Company map[string]interface{}
			Users   []interface{}
		}
		for pb.Next() {
			if err := json.Unmarshal(bb, &m); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
		}
	})
}

func benchmarkStdJSONParseEmptyStruct(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	bb := s2b(s)
	b.RunParallel(func(pb *testing.PB) {
		var m struct{}
		for pb.Next() {
			if err := json.Unmarshal(bb, &m); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
		}
	})
}
