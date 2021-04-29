package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{4,3,2,1}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 4, len(els))

	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{1,2,3,4}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 4, len(els))
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0

	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}