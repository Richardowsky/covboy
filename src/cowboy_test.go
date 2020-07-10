package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"math"
)

var n = 5
var array = []int{2, 1, 1, 3, 2}
var n2 = 3
var array2 = []int{30, 10, 20}

func Test1(t *testing.T) {
	excpect := []int{1, 2, 3, 2, 1}
	assert.Equal(t, solution(n, array), excpect)
	assert.Equal(t, discomfort(solution(n, array)), 1)
}

func Test2(t *testing.T) {
	excpect := []int{10, 30, 20}
	assert.Equal(t, solution(n2, array2), excpect)
	assert.Equal(t, discomfort(solution(n2, array2)), 20)
}

func Benchmark1(b *testing.B) {
	benchmarkZero(n, array, b)
}

func Benchmark2(b *testing.B) {
	benchmarkZero(n2, array2, b)
}

func benchmarkZero(n int, array []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		solution(n, array)
	}
}
func discomfort (arr []int ) int {
	disc := 0
	for i := 0; i < len(arr)-1; i++ {
		if int(math.Abs(float(arr[i] - arr[i+1]))) > disc {
			disc = int(math.Abs(float(arr[i] - arr[i+1])))
		}
	}
	return disc
}

