package bingo

import (
	"fmt"
	"strconv"

	"testing"
)

const (
	intNumber   int   = 100
	int8Number  int8  = 100
	int16Number int16 = 100
	int32Number int32 = 100
	int64Number int64 = 100
)

// test integer
func BenchmarkStrconvFmtInt(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.FormatInt(int64(intNumber), 10)
	}
}
func BenchmarkFmtSprintInt(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprint(intNumber)
	}
}

// test integer 8bit
func BenchmarkStrconvFmtInt8(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.FormatInt(int64(int8Number), 10)
	}
}
func BenchmarkFmtSprintInt8(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprint(int8Number)
	}
}

// test integer 16bit
func BenchmarkStrconvFmtInt16(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.FormatInt(int64(int16Number), 10)
	}
}
func BenchmarkFmtSprintInt16(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprint(int16Number)
	}
}

// test integer 32bit
func BenchmarkStrconvFmtInt32(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.FormatInt(int64(int32Number), 10)
	}
}
func BenchmarkFmtSprintInt32(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprint(int32Number)
	}
}

// test integer 64bit
func BenchmarkStrconvFmtInt64(t *testing.B) {
	for i := 0; i < t.N; i++ {
		strconv.FormatInt(int64Number, 10)
	}
}
func BenchmarkFmtSprintInt64(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fmt.Sprint(int64Number)
	}
}
