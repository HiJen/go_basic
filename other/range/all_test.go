package main

import (
	"fmt"
	"testing"
)

const N = 1000

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "www.flysnow.org"
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}

//-------------------------------------------------

func RangeForMap1(m map[int]string) {
	for k, v := range m {
		_, _ = k, v
	}
}

//const N = 1000

func initMap() map[int]string {
	m := make(map[int]string, N)
	for i := 0; i < N; i++ {
		m[i] = fmt.Sprint("www.flysnow.org", i)
	}
	return m
}

func BenchmarkForMap1(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap1(m)
	}

}

//------value copy, too bad time ------------------------------
func RangeForMap2(m map[int]string) {
	for k, _ := range m {
		_, _ = k, m[k]
	}
}

func BenchmarkRangeForMap2(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap2(m)
	}
}

//------ ------------------------------
