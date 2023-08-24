package main

import (
	"testing"
)

func if_max(values []int) int {
	maxV := values[0]
	for _, v := range values[1:] {
		if v > maxV {
			maxV = v
			continue // no effect
		}
	}
	return maxV
}

func if_max_lol(values []int) int {
	maxV := values[0]
	for _, v := range values[1:] {
		if v > maxV {
			maxV = v
			continue // turns this into a branch loop
		}
		print("lol")
	}
	return maxV
}

func cheat_max(values []int) int {
	maxV := values[0]
	for _, v := range values[1:] {
		maxV = v
	}
	return maxV
}

var Sink int

func BenchmarkFn(b *testing.B) {
	benchData := []struct {
		name string
		fn   func([]int) int
	}{
		//{name: "if_max", fn: if_max},
		{name: "if_max_lol", fn: if_max_lol},
		{name: "cheat_max", fn: cheat_max},
	}

	// generate a large slice of increasing values
	values := make([]int, 1e5)
	for i := range values {
		values[i] = i
	}

	b.ResetTimer()
	for _, bench := range benchData {
		b.Run("name="+bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sink = bench.fn(values)
			}
		})
	}
}
