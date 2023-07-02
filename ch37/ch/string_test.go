package ch

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

// +
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}

	}
	b.StopTimer()
}

// fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

// strings.Builder
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))

		}
		_ = builder.String()
	}
	b.StopTimer()
}

// bytes.Buffer
func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

// strings.join
func BenchmarkStringsJoin(b *testing.B) {
	b.ResetTimer()
	var strs []string
	for i := 0; i < b.N; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	_ = strings.Join(strs, "")
	b.StopTimer()
}

// 切片
func BenchmarkByteSliceString(b *testing.B) {
	b.ResetTimer()
	buf := make([]byte, 0)
	for i := 0; i < b.N; i++ {
		buf = append(buf, byte(i))
	}
	_ = string(buf)
	b.StopTimer()
}
