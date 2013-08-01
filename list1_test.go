package main

import (
	"math/rand"
	"testing"
)

func allocateArrays(n, length int) (ret [][]int) {
	for i := 0; i < n; i++ {
		// Generate a list of 10,000 elements
		ret = append(ret, rand.Perm(length))
	}
	return
}

func benchListTakingFunc(f func([]int) int, b *testing.B, length int) {
	tdata := allocateArrays(b.N, length)
	b.ResetTimer()
	for _, l := range tdata {
		f(l)
	}
}

func BenchmarkHashImplementation500(b *testing.B) {
	benchListTakingFunc(FindSmallest_imp, b, 500)
}

func BenchmarkSortImplementation500(b *testing.B) {
	benchListTakingFunc(FindSmallest, b, 500)
}

func BenchmarkHashImplementation500000(b *testing.B) {
	benchListTakingFunc(FindSmallest_imp, b, 500000)
}

func BenchmarkSortImplementation500000(b *testing.B) {
	benchListTakingFunc(FindSmallest, b, 500000)
}
