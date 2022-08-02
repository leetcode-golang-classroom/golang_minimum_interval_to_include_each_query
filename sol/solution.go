package sol

import (
	"container/heap"
	"sort"
)

type ByStart [][]int

func (a ByStart) Len() int {
	return len(a)
}
func (a ByStart) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Record struct {
	size, right int
}
type MinHeap []Record

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].size < (*h)[j].size ||
		((*h)[i].size == (*h)[j].size && (*h)[i].right < (*h)[j].right)
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(Record))
}
func minInterval(intervals [][]int, queries []int) []int {
	// sort intervals, copy queries
	sort.Sort(ByStart(intervals))
	nLen := len(intervals)
	qLen := len(queries)
	bufferQueries := make([]int, qLen)
	result := make([]int, qLen)
	hash := make(map[int]int)
	copy(bufferQueries, queries)
	sort.Ints(bufferQueries)
	pos := 0
	pq := MinHeap{}
	heap.Init(&pq)
	for _, q := range bufferQueries {
		for pos < nLen && intervals[pos][0] <= q {
			heap.Push(&pq, Record{size: intervals[pos][1] - intervals[pos][0] + 1,
				right: intervals[pos][1],
			})
			pos++
		}
		for pq.Len() > 0 && pq[0].right < q {
			heap.Pop(&pq)
		}
		if pq.Len() > 0 {
			hash[q] = pq[0].size
		} else {
			hash[q] = -1
		}
	}
	for pos := 0; pos < qLen; pos++ {
		result[pos] = hash[queries[pos]]
	}
	return result
}
