package max

import (
	"math/rand"
	"strconv"
	"testing"
)

func if_max_min(values []int) (int, int) {
	minV := values[0]
	maxV := values[0]
	for _, v := range values[1:] {
		if v > maxV {
			maxV = v
			continue // generate branch instruction
		}
		if v < minV {
			minV = v
			continue // generate branch instruction
		}
	}
	return minV, maxV
}

func cheat_min(values []int) (int, int) {
	minV := values[0]
	for _, v := range values[1:] {
		minV = v
	}
	return minV, 0
}

func if_min(values []int) (int, int) {
	minV := values[0]
	for _, v := range values[1:] {
		if v < minV {
			minV = v
			continue // no effect
		}
	}
	return minV, 0
}

func if_max(values []int) (int, int) {
	maxV := values[0]
	for _, v := range values[1:] {
		if v > maxV {
			maxV = v
			continue // no effect
		}
	}
	return 0, maxV
}

func cheat_max(values []int) (int, int) {
	maxV := values[0]
	for _, v := range values[1:] {
		maxV = v
	}
	return 0, maxV
}

func if_pair(values []int) (int, int) {
	minV := values[0]
	maxV := values[0]
	for _, v := range values[1:] {
		if v < minV {
			minV = v // no branching
		}
		if v > maxV {
			maxV = v // no branching
		}
	}
	return minV, maxV
}

var Sink1, Sink2 int

func BenchmarkFn(b *testing.B) {
	benchData := []struct {
		name string
		fn   func([]int) (int, int)
	}{
		{name: "if_max_min", fn: if_max_min},
		{name: "if_max", fn: if_max},
		//{name: "cheat_max", fn: cheat_max},
		//{name: "if_min", fn: if_min},
		//{name: "cheat_min", fn: cheat_min},
		//{name: "if_pair", fn: if_pair},
	}
	values := make([]int, 0, 1e5)
	//newMin := 0
	//newMax := 1
	for i := 0; i < cap(values); i++ {
		//values = append(values, newMin)
		//newMin--
		//values = append(values, newMax)
		//newMax++
		values = append(values, rand.Int())
	}

	b.ResetTimer()
	for _, l := range []int{ /*1, 5, 15, 20, 1e2,*/ 1e5} {
		b.Run("len="+strconv.Itoa(l), func(b *testing.B) {
			for _, bench := range benchData {
				b.Run("name="+bench.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						Sink1, Sink2 = bench.fn(values[:l])
					}
				})
			}
		})
	}
}
