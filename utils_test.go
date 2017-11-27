package project_euler

import (
	"fmt"
	"testing"
)

const N = 1000000

func Benchmark_PrimeGenerator0(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator0()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}

func Benchmark_PrimeGenerator1(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator1()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}

func Benchmark_PrimeGenerator2(b *testing.B) {
	PrimeGenerator := NewPrimeGenerator2()
	for i := 0; i < N; i++ {
		PrimeGenerator.Next()
	}
	fmt.Println(PrimeGenerator.Size())
}
