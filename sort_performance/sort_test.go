package main

import (
	"testing"
)

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort_some_people()
	}
}
