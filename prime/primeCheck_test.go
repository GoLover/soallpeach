package main

import (
	"testing"
)

func BenchmarkPrimeGenerator_IsPrime(b *testing.B) {
	b.ReportAllocs()
	primeGenerator := NewPrimeGenerator()
	for i := 0; i < b.N; i++ {
		primeGenerator.IsPrime(10000000)
	}
}