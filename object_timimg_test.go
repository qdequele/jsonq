package jsonq

import (
	"testing"
)

func BenchmarkKeepSmallSimple(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepSmallSimple()
	}
}

func BenchmarkKeepSmallMedium(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepSmallMedium()
	}
}
func BenchmarkKeepSmallHard(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepSmallHard()
	}
}

func BenchmarkKeepMediumSimple(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepMediumSimple()
	}
}

func BenchmarkKeepMediumMedium(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepMediumMedium()
	}
}
func BenchmarkKeepMediumHard(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepMediumHard()
	}
}

func BenchmarkKeepLargeSimple(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepLargeSimple()
	}
}

func BenchmarkKeepLargeMedium(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepLargeMedium()
	}
}
func BenchmarkKeepLargeHard(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		KeepLargeHard()
	}
}
