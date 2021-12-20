package slice

import (
	"testing"

	"golang.org/x/exp/slices"
)

var _ = slices.Equal[int]

func assertEq[S ~[]E, E comparable](t *testing.T, expected, found S, msg string) {
	t.Helper()
	if !slices.Equal([]E(expected), []E(found)) {
		t.Errorf("assertion %s failed: expected %v but found %v", msg, expected, found)
	}
}

func assertFlatEq[S ~[]E, E comparable](t *testing.T, expected, found []S, msg string) {
	t.Helper()
	if !FlatEq[S, E](expected, found) {
		t.Errorf("assertion %s failed: expected %v but found %v", msg, expected, found)
	}
}

func TestReverse(t *testing.T) {
	Reverse([]int(nil))
	assertEq(t, []int{}, Reverse([]int(nil)), "Reverse(nil)")
	assertEq(t, []int{}, Reverse([]int{}), "Reverse([])")
	assertEq(t, []int{1}, Reverse([]int{1}), "Reverse([1])")
	assertEq(t, []int{2, 1}, Reverse([]int{1, 2}), "Reverse([1, 2])")
	assertEq(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}), "Reverse([1, 2, 3])")
}

func TestWindows(t *testing.T) {
	assertFlatEq(t, [][]int{}, Windows([]int(nil), 1), "Windows(nil, 1)")
	assertFlatEq(t, [][]int{}, Windows([]int{}, 1), "Windows([], 1)")
	assertFlatEq(t, [][]int{{1}}, Windows([]int{1}, 1), "Windows([1], 1)")
	assertFlatEq(t, [][]int{}, Windows([]int{1}, 2), "Windows([1], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, Windows([]int{1, 2}, 2), "Windows([1, 2], 2)")
	assertFlatEq(t, [][]int{{1}, {2}, {3}}, Windows([]int{1, 2, 3}, 1), "Windows([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1, 2}, {2, 3}}, Windows([]int{1, 2, 3}, 2), "Windows([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, Windows([]int{1, 2, 3, 4}, 2), "Windows([1, 2, 3, 4], 2)")
}

func TestRWindows(t *testing.T) {
	assertFlatEq(t, [][]int{}, RWindows([]int(nil), 1), "RWindows(nil, 1)")
	assertFlatEq(t, [][]int{}, RWindows([]int{}, 1), "RWindows([], 1)")
	assertFlatEq(t, [][]int{{1}}, RWindows([]int{1}, 1), "RWindows([1], 1)")
	assertFlatEq(t, [][]int{}, RWindows([]int{1}, 2), "RWindows([1], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, RWindows([]int{1, 2}, 2), "RWindows([1, 2], 2)")
	assertFlatEq(t, [][]int{{3}, {2}, {1}}, RWindows([]int{1, 2, 3}, 1), "RWindows([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{2, 3}, {1, 2}}, RWindows([]int{1, 2, 3}, 2), "RWindows([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{3, 4}, {2, 3}, {1, 2}}, RWindows([]int{1, 2, 3, 4}, 2), "RWindows([1, 2, 3, 4], 2)")
}

func TestChunks(t *testing.T) {
	assertFlatEq(t, [][]int{}, Chunks([]int(nil), 1), "Chunks(nil, 1)")
	assertFlatEq(t, [][]int{}, Chunks([]int{}, 1), "Chunks([], 1)")
	assertFlatEq(t, [][]int{{1}}, Chunks([]int{1}, 1), "Chunks([1], 1)")
	assertFlatEq(t, [][]int{{1}}, Chunks([]int{1}, 2), "Chunks([1], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, Chunks([]int{1, 2}, 2), "Chunks([1, 2], 2)")
	assertFlatEq(t, [][]int{{1}, {2}, {3}}, Chunks([]int{1, 2, 3}, 1), "Chunks([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1, 2}, {3}}, Chunks([]int{1, 2, 3}, 2), "Chunks([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2}, {3, 4}}, Chunks([]int{1, 2, 3, 4}, 2), "Chunks([1, 2, 3, 4], 2)")
}

func TestRChunks(t *testing.T) {
	assertFlatEq(t, [][]int{}, RChunks([]int(nil), 1), "RChunks(nil, 1)")
	assertFlatEq(t, [][]int{}, RChunks([]int{}, 1), "RChunks([], 1)")
	assertFlatEq(t, [][]int{{1}}, RChunks([]int{1}, 1), "RChunks([1], 1)")
	assertFlatEq(t, [][]int{{1}}, RChunks([]int{1}, 2), "RChunks([1], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, RChunks([]int{1, 2}, 2), "RChunks([1, 2], 2)")
	assertFlatEq(t, [][]int{{3}, {2}, {1}}, RChunks([]int{1, 2, 3}, 1), "RChunks([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{2, 3}, {1}}, RChunks([]int{1, 2, 3}, 2), "RChunks([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2, 3}}, RChunks([]int{1, 2, 3}, 3), "RChunks([1, 2, 3], 3)")
	assertFlatEq(t, [][]int{{3, 4}, {1, 2}}, RChunks([]int{1, 2, 3, 4}, 2), "RChunks([1, 2, 3, 4], 2)")
	assertFlatEq(t, [][]int{{2, 3, 4}, {1}}, RChunks([]int{1, 2, 3, 4}, 3), "RChunks([1, 2, 3, 4], 3)")
	assertFlatEq(t, [][]int{{3, 4, 5}, {1, 2}}, RChunks([]int{1, 2, 3, 4, 5}, 3), "RChunks([1, 2, 3, 4, 5], 3)")
}

func TestSplit(t *testing.T) {
	assertFlatEq(t, [][]int{}, Split([]int(nil), 1), "Split(nil, 1)")
	assertFlatEq(t, [][]int{}, Split([]int{}, 1), "Split([], 1)")
	assertFlatEq(t, [][]int{{}}, Split([]int{1}, 1), "Split([1], 1)")
	assertFlatEq(t, [][]int{{1}}, Split([]int{1}, 2), "Split([1], 1)")
	assertFlatEq(t, [][]int{{1}}, Split([]int{1, 2}, 2), "Split([1, 2], 2)")
	assertFlatEq(t, [][]int{{}, {2}}, Split([]int{1, 2}, 1), "Split([1, 2], 1)")
	assertFlatEq(t, [][]int{{}, {2, 3}}, Split([]int{1, 2, 3}, 1), "Split([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1}, {3}}, Split([]int{1, 2, 3}, 2), "Split([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, Split([]int{1, 2, 3}, 3), "Split([1, 2, 3], 3)")
	assertFlatEq(t, [][]int{{}, {}, {}}, Split([]int{1, 1, 1}, 1), "Split([1, 1, 1], 1)")
	assertFlatEq(t, [][]int{{}, {2}, {3}}, Split([]int{1, 2, 1, 3}, 1), "Split([1, 2, 1, 3], 1)")
	assertFlatEq(t, [][]int{{}, {2}, {3}}, Split([]int{1, 2, 1, 3, 1}, 1), "Split([1, 2, 1, 3], 1)")
}

func TestSplitFunc(t *testing.T) {
	eq := func(a int) func(b int) bool {
		return func(b int) bool {
			return a == b
		}
	}
	assertFlatEq(t, [][]int{}, SplitFunc([]int(nil), eq(1)), "SplitFunc(nil, 1)")
	assertFlatEq(t, [][]int{}, SplitFunc([]int{}, eq(1)), "SplitFunc([], 1)")
	assertFlatEq(t, [][]int{{}}, SplitFunc([]int{1}, eq(1)), "SplitFunc([1], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitFunc([]int{1}, eq(2)), "SplitFunc([1], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitFunc([]int{1, 2}, eq(2)), "SplitFunc([1, 2], 2)")
	assertFlatEq(t, [][]int{{}, {2}}, SplitFunc([]int{1, 2}, eq(1)), "SplitFunc([1, 2], 1)")
	assertFlatEq(t, [][]int{{}, {2, 3}}, SplitFunc([]int{1, 2, 3}, eq(1)), "SplitFunc([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1}, {3}}, SplitFunc([]int{1, 2, 3}, eq(2)), "SplitFunc([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2}}, SplitFunc([]int{1, 2, 3}, eq(3)), "SplitFunc([1, 2, 3], 3)")
	assertFlatEq(t, [][]int{{}, {}, {}}, SplitFunc([]int{1, 1, 1}, eq(1)), "SplitFunc([1, 1, 1], 1)")
	assertFlatEq(t, [][]int{{}, {2}, {3}}, SplitFunc([]int{1, 2, 1, 3}, eq(1)), "SplitFunc([1, 2, 1, 3], 1)")
	assertFlatEq(t, [][]int{{}, {2}, {3}}, SplitFunc([]int{1, 2, 1, 3, 1}, eq(1)), "SplitFunc([1, 2, 1, 3], 1)")
}

func TestSplitInclusive(t *testing.T) {
	assertFlatEq(t, [][]int{}, SplitInclusive([]int(nil), 1), "SplitInclusive(nil, 1)")
	assertFlatEq(t, [][]int{}, SplitInclusive([]int{}, 1), "SplitInclusive([], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitInclusive([]int{1}, 1), "SplitInclusive([1], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitInclusive([]int{1}, 2), "SplitInclusive([1], 1)")
	assertFlatEq(t, [][]int{{1, 2}}, SplitInclusive([]int{1, 2}, 2), "SplitInclusive([1, 2], 2)")
	assertFlatEq(t, [][]int{{1}, {2}}, SplitInclusive([]int{1, 2}, 1), "SplitInclusive([1, 2], 1)")
	assertFlatEq(t, [][]int{{1}, {2, 3}}, SplitInclusive([]int{1, 2, 3}, 1), "SplitInclusive([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1, 2}, {3}}, SplitInclusive([]int{1, 2, 3}, 2), "SplitInclusive([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2, 3}}, SplitInclusive([]int{1, 2, 3}, 3), "SplitInclusive([1, 2, 3], 3)")
	assertFlatEq(t, [][]int{{1}, {1}, {1}}, SplitInclusive([]int{1, 1, 1}, 1), "SplitInclusive([1, 1, 1], 1)")
	assertFlatEq(t, [][]int{{1}, {2, 1}, {3}}, SplitInclusive([]int{1, 2, 1, 3}, 1), "SplitInclusive([1, 2, 1, 3], 1)")
	assertFlatEq(
		t, [][]int{{1}, {2, 1}, {3, 1}}, SplitInclusive([]int{1, 2, 1, 3, 1}, 1), "SplitInclusive([1, 2, 1, 3], 1)",
	)
}

func TestSplitInclusiveFunc(t *testing.T) {
	eq := func(a int) func(b int) bool {
		return func(b int) bool {
			return a == b
		}
	}
	assertFlatEq(t, [][]int{}, SplitInclusiveFunc([]int(nil), eq(1)), "SplitInclusiveFunc(nil, 1)")
	assertFlatEq(t, [][]int{}, SplitInclusiveFunc([]int{}, eq(1)), "SplitInclusiveFunc([], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitInclusiveFunc([]int{1}, eq(1)), "SplitInclusiveFunc([1], 1)")
	assertFlatEq(t, [][]int{{1}}, SplitInclusiveFunc([]int{1}, eq(2)), "SplitInclusiveFunc([1], 1)")
	assertFlatEq(t, [][]int{{1, 2}}, SplitInclusiveFunc([]int{1, 2}, eq(2)), "SplitInclusiveFunc([1, 2], 2)")
	assertFlatEq(t, [][]int{{1}, {2}}, SplitInclusiveFunc([]int{1, 2}, eq(1)), "SplitInclusiveFunc([1, 2], 1)")
	assertFlatEq(t, [][]int{{1}, {2, 3}}, SplitInclusiveFunc([]int{1, 2, 3}, eq(1)), "SplitInclusiveFunc([1, 2, 3], 1)")
	assertFlatEq(t, [][]int{{1, 2}, {3}}, SplitInclusiveFunc([]int{1, 2, 3}, eq(2)), "SplitInclusiveFunc([1, 2, 3], 2)")
	assertFlatEq(t, [][]int{{1, 2, 3}}, SplitInclusiveFunc([]int{1, 2, 3}, eq(3)), "SplitInclusiveFunc([1, 2, 3], 3)")
	assertFlatEq(
		t, [][]int{{1}, {1}, {1}}, SplitInclusiveFunc([]int{1, 1, 1}, eq(1)), "SplitInclusiveFunc([1, 1, 1], 1)",
	)
	assertFlatEq(
		t, [][]int{{1}, {2, 1}, {3}}, SplitInclusiveFunc([]int{1, 2, 1, 3}, eq(1)),
		"SplitInclusiveFunc([1, 2, 1, 3], 1)",
	)
	assertFlatEq(
		t, [][]int{{1}, {2, 1}, {3, 1}}, SplitInclusiveFunc([]int{1, 2, 1, 3, 1}, eq(1)),
		"SplitInclusiveFunc([1, 2, 1, 3], 1)",
	)
}

func benchmark[T, R any](b *testing.B, n int, f func([]T) R) {
	b.StopTimer()
	var s = make([]T, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkReverseInt100(b *testing.B) {
	benchmark(b, 100, Reverse[[]int])
}

func BenchmarkReverseBlob100(b *testing.B) {
	benchmark(b, 100, Reverse[[][256]byte])
}

func windows[T any](s []T) [][]T {
	return Windows(s, 2)
}

func BenchmarkWindows10(b *testing.B) {
	benchmark(b, 100, windows[int])
}

func BenchmarkWindows1000(b *testing.B) {
	benchmark(b, 1000, windows[int])
}
